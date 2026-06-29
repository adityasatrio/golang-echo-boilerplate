package initialize

import (
	"github.com/labstack/gommon/log"
	"myapp/configs/rabbitmq/connection"
	"myapp/configs/rabbitmq/recovery"
	"myapp/ent"
)

func RabbitMQInitialize(client *ent.Client) *connection.RabbitMQConnection {
	conf := RabbitMQInitializeWithoutRecovery(client)

	//for recovery reconnection RabbitMQ (backward compatibility):
	consumerFactory := func() recovery.ConsumerRegisterer {
		//fallback to nil for backward compatibility
		return nil
	}
	SetupRabbitMQRecovery(conf, consumerFactory)

	return conf
}

func RabbitMQInitializeWithoutRecovery(client *ent.Client) *connection.RabbitMQConnection {

	newRabbitMQ := connection.NewRabbitMQ()
	_, err := newRabbitMQ.Connection()
	if err != nil {
		log.Errorf("Error closing RabbitMQConnection connection:", err)
	}

	rabbitConf := newRabbitMQ.GetConfig()
	if err != nil {
		log.Errorf("Error closing RabbitMQConnection connection: %v", err)
	}

	return rabbitConf
}

func SetupRabbitMQRecovery(rabbitConf *connection.RabbitMQConnection, consumerFactory func() recovery.ConsumerRegisterer) {
	//for recovery reconnection RabbitMQ:
	go recovery.RabbitMQRecovery(rabbitConf, consumerFactory)
}
