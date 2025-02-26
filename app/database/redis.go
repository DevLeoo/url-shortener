package redis

import (
	"log"

	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func Connect() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := RedisDB.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

}
