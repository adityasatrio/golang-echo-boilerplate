package config

type IRabbitMQConfig interface {
	GetExchangeDirect() string
	GetExchangeKind() string
	GetQueueDirect() string
	GetRoutingKeyDirect() string
	GetExchangeDlx() string
	GetQueueDlq() string
	GetRoutingKeyDlx() string
	GetExchangeJunk() string
	GetQueueJunk() string
	GetRoutingKeyJunk() string
	GetTtl() int64
	GetDelay() int64
	GetLimit() int64
}