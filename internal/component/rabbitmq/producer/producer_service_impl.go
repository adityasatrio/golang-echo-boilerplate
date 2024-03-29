package producer

import (
	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"myapp/internal/component/rabbitmq/channel"
	"myapp/internal/component/rabbitmq/config"
	"myapp/internal/component/rabbitmq/utils"
)

type ProducerServiceImpl struct {
	ch channel.WrappedChannelService
}

func NewProducerService(ch channel.WrappedChannelService) *ProducerServiceImpl {
	return &ProducerServiceImpl{ch: ch}
}

func (p *ProducerServiceImpl) SendToDirect(producer config.IRabbitMQConfig, message []byte) (bool, error) {
	// Send messages to Exchange:
	err := p.ch.PublishMessage(
		producer.GetExchangeDirect(),
		producer.GetRoutingKeyDirect(),
		amqp.Publishing{
			ContentType:  utils.GetContentType(),
			DeliveryMode: amqp.Persistent,
			Body:         message,
		},
	)
	if err != nil {
		log.Errorf("Failed to publish a message direct: %v", err)
		return false, err
	}

	log.Infof("[%s] Sent a message direct: %s\n", producer.GetExchangeDirect(), message)
	return true, err
}

func (p *ProducerServiceImpl) SendToJunk(producer config.IRabbitMQConfig, message []byte) (bool, error) {
	// Send messages to Exchange:
	err := p.ch.PublishMessage(
		producer.GetExchangeJunk(),
		producer.GetRoutingKeyJunk(),
		amqp.Publishing{
			ContentType:  utils.GetContentType(),
			DeliveryMode: amqp.Persistent,
			Body:         message,
		},
	)
	if err != nil {
		log.Errorf("Failed to publish a message junk: %v", err)
		return false, err
	}

	log.Infof("[%s] Sent a message junk: %s\n", producer.GetExchangeJunk(), message)
	return true, err
}
