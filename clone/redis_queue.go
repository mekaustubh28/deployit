package main

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

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

func redis_push(id string, container string) {
	client := GetRedisClient()

	err := client.LPush(ctx, container, id).Err()

	if err != nil {
		log.Fatal(err)
	}
}

func redis_db_set(key string, value string) {
	client := GetRedisClient()

	err := client.Set(ctx, key, value, 0)

	if err == nil {
		log.Fatal("Error while setting value in redis DB", err)
	}
}

func redis_db_get(key string) string {
	client := GetRedisClient()

	value, err := client.Get(ctx, key).Result()

	if err != redis.Nil && err != nil {
		log.Fatal("Error while getting value in redis DB", err)
	}

	return value
}
