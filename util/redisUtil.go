package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var RedisClient *redis.Client
var ctx = context.Background()

func RedisInit() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := RedisClient.Ping(context.Background()).Err()
	return err
}

// GetToken 查看一个token串是否存在于redis中
func GetToken(token string) (string, error) {
	return RedisClient.Get(ctx, token).Result()
}

func SetToken(token string, expiration time.Duration) error {
	return RedisClient.Set(ctx, token, token, expiration).Err()
}

func RenewToken(token string) error {
	return RedisClient.Set(ctx, token, token, expiration).Err()
}
