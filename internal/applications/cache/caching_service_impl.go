package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pierrec/lz4/v4"
	"github.com/vmihailenco/msgpack/v4"
	"time"
)

type CachingServiceImpl struct {
	redisClient *redis.Client
}

func NewCachingServiceImpl(redisClient *redis.Client) *CachingServiceImpl {
	return &CachingServiceImpl{redisClient: redisClient}
}

func (c *CachingServiceImpl) Create(ctx context.Context, key string, data interface{}, expiration time.Duration) (bool, error) {

	serializedData, err := msgpack.Marshal(&data)
	if err != nil {
		fmt.Println("Failed for marshaling data:", err)
		return false, err
	}

	//compress data:
	compressedData, err := compressData(serializedData)

	err = c.redisClient.Set(ctx, key, compressedData, expiration).Err()
	if err != nil {
		fmt.Println("Failed save data on Redis:", err)
	}

	return true, nil
}

func (c *CachingServiceImpl) Get(ctx context.Context, key string, data interface{}) (interface{}, error) {

	redisData, err := c.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		fmt.Println("Failed get data from Redis:", err)
		return nil, err
	}

	//decompress data:
	decompressedData, err := decompressData(redisData, len(redisData))

	err = msgpack.Unmarshal(decompressedData, data)
	if err != nil {
		fmt.Println("Failed for unMarshaling data:", err)
		return nil, err
	}

	return data, nil
}

func (c *CachingServiceImpl) Delete(ctx context.Context, key string) (bool, error) {
	_, err := c.redisClient.Del(ctx, key).Result()
	if err != nil {
		fmt.Println("Failed for delete data on redis:", err)
		return false, err
	}

	return true, nil
}

// this method for compress data with LZ4:
func compressData(data []byte) ([]byte, error) {
	compressedSize := lz4.CompressBlockBound(len(data))
	compressedData := make([]byte, compressedSize)
	compressedSize, err := lz4.CompressBlock(data, compressedData, nil)
	if err != nil {
		return nil, err
	}

	return compressedData[:compressedSize], nil
}

// this method for decompress data from LZ4:
func decompressData(compressedData []byte, originalSize int) ([]byte, error) {
	decompressedData := make([]byte, originalSize*10)
	_, err := lz4.UncompressBlock(compressedData, decompressedData)
	if err != nil {
		return nil, err
	}

	return decompressedData, nil
}
