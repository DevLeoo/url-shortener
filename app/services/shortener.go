package services

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"url-shortener/app/config"
	redis "url-shortener/app/database"

	"golang.org/x/exp/rand"
)

type Params interface {
	GetURL() []string
	GetShortKeys() []string
}

func Shorten(params Params) ([]string, error) {
	start := time.Now()
	port := config.Port
	urls := params.GetURL()
	if len(urls) == 0 {
		return nil, errors.New("missing URL")
	}

	var wg sync.WaitGroup
	shortenURLs := make([]string, len(urls))
	errCh := make(chan error, len(urls))

	for i, url := range urls {
		wg.Add(1)

		go func(index int, originalURL string) {
			defer wg.Done()

			shortKey := generateShortKey()
			err := redis.RedisDB.Set(shortKey, originalURL, 0).Err()
			if err != nil {
				errCh <- err
				return
			}

			shortenURLs[index] = fmt.Sprintf("http://localhost:%d/%s", port, shortKey)
		}(i, url)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		return nil, <-errCh
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)

	return shortenURLs, nil
}

func generateShortKey() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(uint64(time.Now().UnixNano()))
	key := make([]byte, 6)
	for i := range key {
		key[i] = chars[rand.Intn(len(chars))]
	}
	return string(key)
}
