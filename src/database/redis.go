package database

import (
	"app/src/config"
	"app/src/utils"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

// ConnectRedis establishes a connection to Redis
func ConnectRedis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	client := redis.NewClient(config.RedisConfig())

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		utils.Log.Errorf("Failed to connect to Redis: %v", err)
		return nil
	}

	utils.Log.Info("Connected to Redis successfully")
	redisClient = client
	return redisClient
}

// CloseRedis closes the Redis connection
func CloseRedis(client *redis.Client) {
	if client == nil {
		return
	}

	if err := client.Close(); err != nil {
		utils.Log.Errorf("Error closing Redis connection: %v", err)
	} else {
		utils.Log.Info("Redis connection closed successfully")
	}
}

// PingRedis checks if Redis connection is working
func PingRedis() error {
	if redisClient == nil {
		return fmt.Errorf("Redis client not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		utils.Log.Errorf("Redis ping failed: %v", err)
		return err
	}

	return nil
}

// GetKey generates a Redis key with the configured prefix
func GetKey(key string) string {
	return fmt.Sprintf("%s:%s", config.RedisPrefix, key)
}

// Set stores a value in Redis with expiration
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if redisClient == nil {
		return fmt.Errorf("Redis client not initialized")
	}

	err := redisClient.Set(ctx, GetKey(key), value, expiration).Err()
	if err != nil {
		utils.Log.Errorf("Failed to set Redis key %s: %v", key, err)
		return err
	}
	return nil
}

// Get retrieves a value from Redis
func Get(ctx context.Context, key string) (string, error) {
	if redisClient == nil {
		return "", fmt.Errorf("Redis client not initialized")
	}

	val, err := redisClient.Get(ctx, GetKey(key)).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		utils.Log.Errorf("Failed to get Redis key %s: %v", key, err)
		return "", err
	}
	return val, nil
}

// Delete removes a key from Redis
func Delete(ctx context.Context, key string) error {
	if redisClient == nil {
		return fmt.Errorf("Redis client not initialized")
	}

	err := redisClient.Del(ctx, GetKey(key)).Err()
	if err != nil {
		utils.Log.Errorf("Failed to delete Redis key %s: %v", key, err)
		return err
	}
	return nil
}

// HashSet sets a hash field in Redis
func HashSet(ctx context.Context, key, field string, value interface{}) error {
	if redisClient == nil {
		return fmt.Errorf("Redis client not initialized")
	}

	err := redisClient.HSet(ctx, GetKey(key), field, value).Err()
	if err != nil {
		utils.Log.Errorf("Failed to set Redis hash field %s:%s: %v", key, field, err)
		return err
	}
	return nil
}

// HashGet gets a hash field from Redis
func HashGet(ctx context.Context, key, field string) (string, error) {
	if redisClient == nil {
		return "", fmt.Errorf("Redis client not initialized")
	}

	val, err := redisClient.HGet(ctx, GetKey(key), field).Result()
	if err == redis.Nil {
		return "", nil // Field does not exist
	} else if err != nil {
		utils.Log.Errorf("Failed to get Redis hash field %s:%s: %v", key, field, err)
		return "", err
	}
	return val, nil
}
