package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

// top level declarations for StorageService and Redis Context 
var storeService = &StorageService{}
var redisCtx = context.Background() //context allows us to control connection context and manage timeouts

const CacheDuration = 6 * time.Hour // Cache duration for the shortened URLs

func InitializeStore() *StorageService {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})
	pong, err := redisClient.Ping(redisCtx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	} else {
		fmt.Printf("Redis started successfully: pong message = {%s}", pong)
	}
	storeService.redisClient = redisClient
	return storeService
}

// function to save the mapping between short and orignal URL
func StoreMapping (shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(redisCtx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl))
	}
} 

// function to retrieve the original URL from the short URL when the user calls the short URL
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(redisCtx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}