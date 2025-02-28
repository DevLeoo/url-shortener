package redis

import (
	"log"

	"github.com/go-redis/redis"
)

func Connect() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		redisClient.Close()
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	return redisClient, nil
}
