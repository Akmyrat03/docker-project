package test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRedisSetGet(t *testing.T) {
	ctx := context.Background()

	rdb := redis.NewClient(
		&redis.Options{
			Addr: "redis-server:6379",
		},
	)

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		t.Fatalf("Redis is not reachable: %v", err)
	}

	err := rdb.Set(ctx, "test-key", "test-value", 10*time.Minute).Err()
	if err != nil {
		t.Fatalf("Failed to set value: %v", err)
	}

	val, err := rdb.Get(ctx, "test-key").Result()
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if val != "test-value" {
		t.Fatalf("Expected 'test-value', got '%s'", val)
	}
}
