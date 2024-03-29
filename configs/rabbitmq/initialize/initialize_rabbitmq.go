package initialize

import (
	"github.com/labstack/gommon/log"
	"myapp/configs/rabbitmq/connection"
	"myapp/configs/rabbitmq/recovery"
	"myapp/ent"
)

func RabbitMQInitialize(client *ent.Client) *connection.RabbitMQConnection {

	newRabbitMQ := connection.NewRabbitMQ()
	_, err := newRabbitMQ.Connection()
	if err != nil {
		log.Errorf("Error closing RabbitMQConnection connection:", err)
	}

	rabbitConf := newRabbitMQ.GetConfig()
	if err != nil {
		log.Errorf("Error closing RabbitMQConnection connection: %v", err)
	}

	//for recovery reconnection RabbitMQ:
	go recovery.RabbitMQRecovery(client, rabbitConf)

	return rabbitConf
}
