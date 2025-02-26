package services

import (
	"errors"
	"fmt"
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

	var shortenURLs []string
	for _, url := range urls {
		shortKey := generateShortKey()
		err := redis.RedisDB.Set(shortKey, url, 0).Err()
		if err != nil {
			return nil, err
		}

		shortURL := fmt.Sprintf("http://localhost:%d/%s", port, shortKey)
		shortenURLs = append(shortenURLs, shortURL)
	}

	elapse := time.Since(start)
	fmt.Printf("Shorten took %s\n", elapse)
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
