package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var RedisClient *redis.Client

func ConnectRedis() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file, error=%s", err.Error())
	}

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	RedisClient := rdb

	RedisClient.Set(ctx, "foo", "bar", 0)
	result, err := RedisClient.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the Redis")
	fmt.Println(result)
	return nil
}
