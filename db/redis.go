package db

import (
	"github.com/redis/go-redis/v9"
)

func RedisConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return client

}
