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

func SaveUrl(shortURL, originalURL string) error {
	err := storageService.redisClient.Set(ctx, shortURL, originalURL, cacheDuration).Err()
	if err != nil {
		return err
	}
	return nil
}
func RetrieveUrl(shortURL string) (string, error) {
	result, err := storageService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
func CheckExistence(shortURL string) (bool, error) {
	result, err := storageService.redisClient.Exists(ctx, shortURL).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}
