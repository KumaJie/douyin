package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client

func RedisInit() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Ping(context.Background()).Err()
	return err
}

func GetToken(token string) (string, error) {
	return rdb.Get(context.Background(), token).Result()
}

func SetToken(token string, expiration time.Duration) error {
	return rdb.Set(context.Background(), token, token, expiration).Err()
}
