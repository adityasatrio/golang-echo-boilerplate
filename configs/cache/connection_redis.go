package cache

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"myapp/configs/credential"
)

func NewRedisClient() *redis.Client {

	addr := fmt.Sprintf("%s:%s",
		credential.GetString("cache.configs.redis.host"),
		credential.GetString("cache.configs.redis.port"))

	redisOptions := redis.Options{
		Addr:     addr,
		Username: credential.GetString("cache.configs.redis.username"),
		Password: credential.GetString("cache.configs.redis.password"),
		DB:       credential.GetInt("cache.configs.redis.db"),
		PoolSize: credential.GetInt("cache.configs.redis.poolSize"),
	}

	isUseTls := credential.GetBool("cache.configs.redis.isTls")

	if isUseTls {

		tlsConfig := &tls.Config{
			InsecureSkipVerify: credential.GetBool("cache.configs.redis.insecureSkipVerify"),
		}

		redisOptions.TLSConfig = tlsConfig

		client := redis.NewClient(&redisOptions)
		return client
	}

	client := redis.NewClient(&redisOptions)
	return client
}
