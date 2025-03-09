package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     getEndpoint(),
		Password: getPassword(),
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("Redisに接続できませんでした")
	}
	return client
}

func getEndpoint() string {
	address, ok := os.LookupEnv("REDIS_ADDRESS")
	if !ok {
		panic("環境変数に\"REDIS_ADDRESS\"が設定されていません")
	}
	return address
}

func getPassword() string {
	password, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		panic("環境変数に\"REDIS_PASSWORD\"が設定されていません")
	}
	return password
}
