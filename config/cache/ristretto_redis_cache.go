package cache

import (
	"context"
	"fmt"
	//"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	redisStore "github.com/eko/gocache/store/redis/v4"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"time"
	//ristrettoStore "github.com/eko/gocache/store/ristretto/v4"
)

const RistrettoDefaultBufferItems = 64

func NewCacheManager() *cache.ChainCache[any] {

	/*ristrettoNumCounters := viper.GetInt64("cache.config.ristretto.numCounters")
	ristrettoMaxCost := viper.GetInt64("cache.config.ristretto.maxCost")
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: ristrettoNumCounters,
		MaxCost:     ristrettoMaxCost,
		BufferItems: RistrettoDefaultBufferItems,
	})*/

	/*if err != nil {
		panic(err)
	}*/
	log.Default().Println("initialized ristrettoCache connection: success")

	redisUsername := viper.GetString("cache.config.redis.username")
	redisPassword := viper.GetString("cache.config.redis.password")
	redisDb := viper.GetInt("cache.config.redis.DB")
	redisPoolSize := viper.GetInt("cache.config.redis.poolSize")
	addr := fmt.Sprintf("%s:%s",
		viper.GetString("cache.config.redis.host"),
		viper.GetString("cache.config.redis.port"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: redisUsername,
		Password: redisPassword,
		DB:       redisDb,
		PoolSize: redisPoolSize,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			pong := cn.Ping(ctx)
			log.Default().Println("OnConnect: initialized cache redis connection: success ", pong)
			return nil
		},
	})

	//pong := redisClient.Ping(e.AcquireContext().Request().Context())
	//log.Default().Println("initialized cache redis connection: success ", pong)

	//ristrettoStoreConn := ristrettoStore.NewRistretto(ristrettoCache)
	redisStoreConn := redisStore.NewRedis(redisClient, store.WithExpiration(5*time.Second))

	cacheManager := cache.NewChain[any](
		//cache.New[any](ristrettoStoreConn),
		cache.New[any](redisStoreConn),
	)

	log.Default().Println("initialized cache redis  connection: success")
	return cacheManager
}
