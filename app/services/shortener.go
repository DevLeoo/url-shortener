package services

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"url-shortener/app/config"
	redisDB "url-shortener/app/database"

	"github.com/go-redis/redis"
	"golang.org/x/exp/rand"
)

const maxShorteningWorkers = 100
const maxRedisWorkers = 10

type Params interface {
	GetURL() []string
	GetShortKeys() []string
}

type shortJob struct {
	index    int
	url      string
	shortKey string
}

type redisResult struct {
	index int
	short string
	err   error
}

func Shorten(params Params) ([]string, error) {
	start := time.Now()
	port := config.Port
	urls := params.GetURL()
	if len(urls) == 0 {
		return nil, errors.New("missing URL")
	}

	redisClient, err := redisDB.Connect()
	if err != nil {
		return nil, err
	}
	defer redisClient.Close()

	var wg sync.WaitGroup
	shortenURLs := make([]string, len(urls))
	jobs := make(chan shortJob, len(urls))
	results := make(chan redisResult, len(urls))

	shorteningWg := sync.WaitGroup{}
	shorteningWg.Add(len(urls))
	for i, url := range urls {
		go func(index int, originalURL string) {
			defer shorteningWg.Done()
			shortKey := generateShortKey()
			jobs <- shortJob{index: index, url: originalURL, shortKey: shortKey}
		}(i, url)
	}

	go func() {
		shorteningWg.Wait()
		close(jobs)
	}()

	for w := 0; w < maxRedisWorkers; w++ {
		wg.Add(1)
		go redisWorker(&wg, redisClient, jobs, results, port)
	}

	for i := 0; i < len(urls); i++ {
		res := <-results
		if res.err != nil {
			return nil, res.err
		}
		shortenURLs[res.index] = res.short
	}

	wg.Wait()
	close(results)
	close(jobs)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)

	return shortenURLs, nil
}

// Worker responsável por inserir no Redis
func redisWorker(wg *sync.WaitGroup, redisClient *redis.Client, jobs <-chan shortJob, results chan<- redisResult, port int) {
	defer wg.Done()
	for j := range jobs {
		err := redisClient.Set(j.shortKey, j.url, 0).Err()
		if err != nil {
			results <- redisResult{index: j.index, err: err}
			continue
		}
		results <- redisResult{index: j.index, short: fmt.Sprintf("http://localhost:%d/%s", port, j.shortKey)}
	}
}

// Gera uma chave curta aleatória
func generateShortKey() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(uint64(time.Now().UnixNano()))
	key := make([]byte, 6)
	for i := range key {
		key[i] = chars[rand.Intn(len(chars))]
	}
	return string(key)
}
