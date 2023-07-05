package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func NewRedisClient() *redis.Client {
	addr := fmt.Sprintf("%s:%s",
		viper.GetString("cache.configs.redis.host"),
		viper.GetString("cache.configs.redis.port"))

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: viper.GetString("cache.configs.redis.username"),
		Password: viper.GetString("cache.configs.redis.password"),
		DB:       viper.GetInt("cache.configs.redis.db"),
		PoolSize: viper.GetInt("cache.configs.redis.poolSize"),
	})
	return client
}
