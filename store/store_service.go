package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const cacheDuration = time.Hour * 3

var (
	storageService = &StorageService{}
	ctx            = context.Background()
)

type StorageService struct {
	redisClient *redis.Client
}

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis ping: ", pong)
	storageService.redisClient = redisClient
	return storageService
}

func SaveUrl(shortURL, originalURL string) {
	err := storageService.redisClient.Set(ctx, shortURL, originalURL, cacheDuration).Err()
	if err != nil {
		panic(err)
	}
}

func RetriveUrl(shortURL string) string {
	result, err := storageService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		panic(err)
	}
	return result
}
