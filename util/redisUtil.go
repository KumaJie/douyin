package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var RedisClient *redis.Client

func RedisInit() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := RedisClient.Ping(context.Background()).Err()
	return err
}

func GetToken(token string) (string, error) {
	return RedisClient.Get(context.Background(), token).Result()
}

func SetToken(token string, expiration time.Duration) error {
	return RedisClient.Set(context.Background(), token, token, expiration).Err()
}
