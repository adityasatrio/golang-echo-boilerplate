package cache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

func TestCachingImpl(t *testing.T) {
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

	// Create a CacheImpl instance with Redis client:
	cachingService := NewCache(redisClient)

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

func TestCachingServiceImpl_failureConnection(t *testing.T) {
	// Create a mini redis instance to mock the Redis server:
	mockRedis := miniredis.NewMiniRedis()

	// Start Redis server mock:
	if err := mockRedis.Start(); err != nil {
		t.Fatalf("Failed to start mock Redis server: %v", err)
	}
	//defer mockRedis.Close()

	// Create a Redis client using a mock Redis address:
	redisClient := redis.NewClient(&redis.Options{
		Addr: mockRedis.Addr(),
	})

	// Create a CachingServiceImpl instance with Redis client:
	cachingService := NewCache(redisClient)
	mockRedis.Close()

	// Test method Create
	t.Run("Failure Ping", func(t *testing.T) {
		err := cachingService.Ping(context.Background())
		assert.Error(t, err)
	})

	// Data to be stored and retrieved from cache:
	data := "Hello, world!"
	key := "myKey"
	expiration := 1 * time.Hour

	// Test method Create
	t.Run("Failure Create", func(t *testing.T) {
		result, err := cachingService.Create(context.Background(), key, data, expiration)
		assert.Nil(t, err)
		assert.False(t, result)
	})

	// Test method Get
	t.Run("Failure Get", func(t *testing.T) {
		var result string
		output, err := cachingService.Get(context.Background(), key, &result)
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	// Test method Delete
	t.Run("Failure Delete", func(t *testing.T) {
		result, err := cachingService.Delete(context.Background(), key)
		assert.Error(t, err)
		assert.False(t, result)

	})

}

func TestCachingServiceImpl_failureMsgPack(t *testing.T) {
	// Create a mini redis instance to mock the Redis server:
	mockRedis := miniredis.NewMiniRedis()

	// Start Redis server mock:
	if err := mockRedis.Start(); err != nil {
		t.Fatalf("Failed to start mock Redis server: %v", err)
	}
	//defer mockRedis.Close()

	// Create a Redis client using a mock Redis address:
	redisClient := redis.NewClient(&redis.Options{
		Addr: mockRedis.Addr(),
	})

	// Create a CachingServiceImpl instance with Redis client:
	cachingService := NewCache(redisClient)
	//mockRedis.Close()

	// Data to be stored and retrieved from cache:
	//data := "Hello, world!"
	key := "myKey"
	expiration := 1 * time.Hour

	// Test method Create
	t.Run("Failure Create", func(t *testing.T) {
		result, err := cachingService.Create(context.Background(), key, func() {}, expiration)
		assert.NotNil(t, err)
		assert.False(t, result)
	})
}
