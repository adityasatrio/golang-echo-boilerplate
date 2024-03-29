package cache

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

func TestCachingServiceImpl(t *testing.T) {
	// Create a mini redis instance to mock the Redis server:
	mockRedis := miniredis.NewMiniRedis()

	// Start Redis server mock:
	if err := mockRedis.Start(); err != nil {
		t.Fatalf("Failed to start mock Redis server: %v", err)
	}
	defer mockRedis.Close()

	// Create a Redis client using a mock Redis address:
	redisClient := redis.NewClient(&redis.Options{
		Addr: mockRedis.Addr(),
	})

	// Create a CachingServiceImpl instance with Redis client:
	cachingService := NewCachingService(redisClient)

	// Data to be stored and retrieved from cache:
	data := "Hello, world!"
	key := "myKey"
	expiration := 1 * time.Hour

	// Test method Create
	t.Run("Ping", func(t *testing.T) {
		err := cachingService.Ping(context.Background())
		if err != nil {
			t.Errorf("Failed to create cache: %v", err)
		}
	})

	// Test method Create
	t.Run("Create", func(t *testing.T) {
		_, err := cachingService.Create(context.Background(), key, data, expiration)
		if err != nil {
			t.Errorf("Failed to create cache: %v", err)
		}
	})

	// Test method Get
	t.Run("Get", func(t *testing.T) {
		var result string
		_, err := cachingService.Get(context.Background(), key, &result)
		if err != nil {
			t.Errorf("Failed to get cache: %v", err)
		}

		if result != data {
			t.Errorf("Unexpected result. Expected: %s, Got: %s", data, result)
		}
	})

	// Test method Delete
	t.Run("Delete", func(t *testing.T) {
		_, err := cachingService.Delete(context.Background(), key)
		if err != nil {
			t.Errorf("Failed to delete cache: %v", err)
		}
	})

}
