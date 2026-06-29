package recovery

import (
	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"myapp/configs/rabbitmq/connection"
	"myapp/internal/component/rabbitmq/registry"
	"time"
)

type ConsumerRegisterer interface {
	Register()
}

func RabbitMQRecovery(rabbitConf *connection.RabbitMQConnection, consumerFactory func() ConsumerRegisterer) {
	go func() {
		for {
			reason, ok := <-rabbitConf.GetConnection().NotifyClose(make(chan *amqp.Error))
			if !ok {
				log.Errorf("connection closed")
				break
			}
			log.Errorf("connection closed, reason: %v", reason)

			for {
				//time sleep for waiting connection up:
				timeRecovery := viper.GetInt("rabbitmq.configs.recovery")
				time.Sleep(time.Duration(timeRecovery) * time.Second)

				err := rabbitConf.Reconnect()
				if err == nil {
					log.Infof("reconnect rabbitmq success")

					//rabbitmq registry exchange, queue, dlq and other:
					registerMq := registry.NewProducerRegistry(rabbitConf)
					registerMq.Register()

					//rabbitmq registry consumer:
					registerConsumer := consumerFactory()
					registerConsumer.Register()

					break
				}

				log.Errorf("reconnect rabbitmq failed, err: %v", err)
			}

		}
	}()
}
