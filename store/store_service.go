package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
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