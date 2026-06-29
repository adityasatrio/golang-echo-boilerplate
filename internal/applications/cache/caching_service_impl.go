package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"github.com/vmihailenco/msgpack/v4"
	"time"
)

type CachingServiceImpl struct {
	redisClient *redis.Client
}

func NewCachingService(redisClient *redis.Client) *CachingServiceImpl {
	return &CachingServiceImpl{redisClient: redisClient}
}

func (c *CachingServiceImpl) Ping(ctx context.Context) error {
	_, err := c.redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *CachingServiceImpl) Create(ctx context.Context, key string, data interface{}, expiration time.Duration) (bool, error) {

	serializedData, err := msgpack.Marshal(&data)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}

		log.Errorf("Failed for marshaling data:", err)
		return false, err
	}

	compressedData, err := CompressData(serializedData)
	if err != nil {
		log.Errorf("Failed for compressing data:", err)
		return false, err
	}

	err = c.redisClient.Set(ctx, key, compressedData, expiration).Err()
	if err != nil {
		log.Errorf("Failed save data on Redis:", err)
	}

	return true, nil
}

func (c *CachingServiceImpl) Get(ctx context.Context, key string, data interface{}) (interface{}, error) {

	redisData, err := c.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		log.Errorf("Failed get data from Redis:", err)
		return nil, err
	}

	decompressedData, err := DecompressData(redisData, len(redisData))
	if err != nil {
		log.Errorf("Failed for decompressing data:", err)
		return nil, err
	}

	err = msgpack.Unmarshal(decompressedData, data)
	if err != nil {
		log.Errorf("Failed for unMarshaling data:", err)
		return nil, err
	}

	return data, nil
}

func (c *CachingServiceImpl) Delete(ctx context.Context, key string) (bool, error) {
	_, err := c.redisClient.Del(ctx, key).Result()
	if err != nil {
		log.Errorf("Failed for delete data on redis:", err)
		return false, err
	}

	return true, nil
}
