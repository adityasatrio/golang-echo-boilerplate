package inbound

import (
	"context"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"myapp/internal/applications/example/rabbitmq/service"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/component/rabbitmq/channel"
	"myapp/internal/component/rabbitmq/config"
	"myapp/internal/component/rabbitmq/producer"
	"myapp/internal/component/rabbitmq/utils"
)

type ExampleRabbitMQInboundImpl struct {
	ch             channel.WrappedChannelService
	exampleService service.ExampleRabbitMQService
	producer       producer.Producer
}

func NewExampleRabbitMQInbound(ch channel.WrappedChannelService, exampleService service.ExampleRabbitMQService, producer producer.Producer) *ExampleRabbitMQInboundImpl {
	return &ExampleRabbitMQInboundImpl{ch: ch, exampleService: exampleService, producer: producer}
}

func (t *ExampleRabbitMQInboundImpl) GetMessage(cfg config.IRabbitMQConfig) (bool, error) {

	// Consuming messages from the Queue:
	message, err := t.ch.ConsumeMessage(cfg.GetQueueDirect())

	if err != nil {
		log.Errorf("Failed to consume messages: %v", err)
		return false, err
	}

	go func() {
		for msg := range message {
			log.Infof("Received a message: %s", msg.Body)

			count := utils.CheckLimitRetry(msg)

			// Process message here:
			data := parsingData(msg.Body)
			_, err := t.exampleService.GetMessage(context.Background(), &data)
			if err != nil {
				log.Warnf("Failed to process service %v", err)
			} else {
				err = msg.Ack(false)
			}

			if err != nil {
				isHasExceeded := utils.IsHasExceeded(cfg.GetLimit(), count, msg)
				if isHasExceeded {
					// Process message to junk here:
					_, err := t.producer.SendToJunk(cfg, msg.Body)
					if err != nil {
						log.Warnf("Failed to publish a message junk: %v", err)
					}
				}
			}
		}
	}()
	return true, err
}

func parsingData(msg []byte) dto.SystemParameterCreateRequest {
	data := dto.SystemParameterCreateRequest{}
	err := json.Unmarshal(msg, &data)

	if err != nil {
		log.Infof("Failed to parsing the message: %v", err)
		return data
	}

	return data
}
