package services

import (
	"errors"
	redis "url-shortener/app/database"
)

func RedirectURL(params Params) ([]string, error) {
	shortKeyParams := params.GetShortKeys()

	var shortKeys []string
	redisClient, err := redis.Connect()
	if err != nil {
		return nil, err
	}
	defer redisClient.Close()

	for _, shortKey := range shortKeyParams {

		longURL, err := redisClient.Get(shortKey).Result()
		if err != nil {
			return nil, errors.New("url not found")
		}
		shortKeys = append(shortKeys, longURL)
	}

	return shortKeys, nil
}
