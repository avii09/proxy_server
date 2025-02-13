package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	client *redis.Client
)

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,
	})

	// Test the connection
	_, err := client.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Connected to Redis")
}

func GetClient() *redis.Client {
	return client
}
