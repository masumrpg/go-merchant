package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConfig returns the default redis client configuration
func RedisConfig() *redis.Options {
	return &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", RedisHost, RedisPort),
		Password:     RedisPassword,
		DB:           RedisDB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		PoolTimeout:  5 * time.Second,
		MaxRetries:   3,
		MinIdleConns: 5,
	}
}

// RedisContext creates a context with timeout for redis operations
func RedisContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
