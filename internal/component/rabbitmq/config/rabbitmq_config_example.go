package config

import (
	"github.com/spf13/viper"
)

type RabbitMQConfigExample struct {
	RabbitMQConfig
}

func NewRabbitMQConfigExample() IRabbitMQConfig {
	return &RabbitMQConfigExample{
		RabbitMQConfig: RabbitMQConfig{
			ExchangeKind: viper.GetString("rabbitmq.example.exchangeKind"),

			ExchangeDirect:   viper.GetString("rabbitmq.example.exchangeDirect"),
			QueueDirect:      viper.GetString("rabbitmq.example.queueDirect"),
			RoutingKeyDirect: viper.GetString("rabbitmq.example.routingKeyDirect"),

			ExchangeDlx:   viper.GetString("rabbitmq.example.exchangeDlx"),
			QueueDlq:      viper.GetString("rabbitmq.example.queueDlq"),
			RoutingKeyDlx: viper.GetString("rabbitmq.example.routingKeyDlx"),

			ExchangeJunk:   viper.GetString("rabbitmq.example.exchangeJunk"),
			QueueJunk:      viper.GetString("rabbitmq.example.queueJunk"),
			RoutingKeyJunk: viper.GetString("rabbitmq.example.routingKeyJunk"),

			Ttl:   viper.GetInt64("rabbitmq.example.ttl"),
			Delay: viper.GetInt64("rabbitmq.example.delay"),
			Limit: viper.GetInt64("rabbitmq.example.limit"),
		}}
}
