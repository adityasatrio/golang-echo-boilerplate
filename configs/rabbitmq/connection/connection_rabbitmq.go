package connection

import (
	"fmt"
	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"myapp/configs/credential"
)

type RabbitMQConnection struct {
	name    string
	conn    *amqp.Connection
	channel *amqp.Channel
	err     chan error
}

func NewRabbitMQ() *RabbitMQConnection {
	c := &RabbitMQConnection{
		err: make(chan error),
	}
	return c
}

func (c *RabbitMQConnection) Connection() (*amqp.Connection, error) {

	config := fmt.Sprintf("amqp://%s:%s@%s:%s",
		credential.GetString("rabbitmq.configs.username"),
		credential.GetString("rabbitmq.configs.password"),
		credential.GetString("rabbitmq.configs.host"),
		credential.GetString("rabbitmq.configs.port"))

	var err error
	c.conn, err = amqp.Dial(config)
	if err != nil {
		log.Errorf("Failed to connect to RabbitMQConnection: %v", err)
		return nil, err
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		log.Errorf("Failed to open a channel: %v", err)
	}

	return c.conn, nil
}

func (c *RabbitMQConnection) ChannelRabbitMQ(conn *amqp.Connection) *amqp.Channel {
	var err error
	c.channel, err = conn.Channel()
	if err != nil {
		log.Errorf("Failed to open a channel: %v", err)
	}

	return c.channel
}

func (c *RabbitMQConnection) Reconnect() error {
	if _, err := c.Connection(); err != nil {
		return err
	}

	return nil
}

func (c *RabbitMQConnection) GetConfig() *RabbitMQConnection {
	return c
}

func (c *RabbitMQConnection) GetConnection() *amqp.Connection {
	return c.conn
}

func (c *RabbitMQConnection) GetChannel() *amqp.Channel {
	return c.channel
}
