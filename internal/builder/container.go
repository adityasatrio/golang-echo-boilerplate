package builder

import (
	"github.com/go-redis/redis/v8"
	"myapp/configs/rabbitmq/connection"
	"myapp/ent"
	"myapp/internal/component/cache"
	"myapp/internal/component/rabbitmq/channel"
	"myapp/internal/component/rabbitmq/producer"
	"myapp/internal/component/transaction"
)

// Container holds all external dependencies for the application.
// It provides a fluent API for building services.
type Container struct {
	db     *ent.Client
	redis  *redis.Client
	rabbit *connection.RabbitMQConnection
}

// NewBuilder creates a new empty container with no dependencies set.
// Dependencies must be added using the With* methods.
func NewBuilder() *Container {
	return &Container{}
}

// WithDatabase sets the database client dependency.
func (c *Container) WithDatabase(db *ent.Client) *Container {
	c.db = db
	return c
}

// WithCache sets the Redis client dependency.
func (c *Container) WithCache(redis *redis.Client) *Container {
	c.redis = redis
	return c
}

// WithRabbit sets the RabbitMQ connection dependency.
func (c *Container) WithRabbit(rabbit *connection.RabbitMQConnection) *Container {
	c.rabbit = rabbit
	return c
}

// BuildTrx builds the transaction component.
func (c *Container) BuildTrx() transaction.Trx {
	return transaction.NewTrx(c.db)
}

// BuildCache builds the cache component.
func (c *Container) BuildCache() cache.Cache {
	return cache.NewCache(c.redis)
}

// BuildProducer builds the RabbitMQ producer component.
func (c *Container) BuildProducer() producer.Producer {
	wrappedChannel := channel.NewWrappedChannel(c.rabbit)
	return producer.NewProducerService(wrappedChannel)
}
