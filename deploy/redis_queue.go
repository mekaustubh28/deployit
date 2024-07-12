package main

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	once   sync.Once
)

func InitializeRedisClient() {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
	})
}

func init() {
	InitializeRedisClient()
}

// GetRedisClient returns the Redis client
func GetRedisClient() *redis.Client {
	return client
}

func redis_pop(queue string) string {
	client := GetRedisClient()
	item, _ := client.RPop(queue).Result()
	return item
}

func redis_push(id string, container string) {
	client := GetRedisClient()

	err := client.LPush(container, id).Err()

	if err != nil {
		log.Fatal(err)
	}
}
