package config

type RabbitMQConfig struct {
	ExchangeKind     string
	ExchangeDirect   string
	QueueDirect      string
	RoutingKeyDirect string

	ExchangeDlx   string
	QueueDlq      string
	RoutingKeyDlx string

	ExchangeJunk   string
	QueueJunk      string
	RoutingKeyJunk string

	Ttl   int64
	Delay int64
	Limit int64
}

func (p *RabbitMQConfig) GetExchangeKind() string {
	return p.ExchangeKind
}

func (p *RabbitMQConfig) GetExchangeDirect() string {
	return p.ExchangeDirect
}

func (p *RabbitMQConfig) GetQueueDirect() string {
	return p.QueueDirect
}

func (p *RabbitMQConfig) GetRoutingKeyDirect() string {
	return p.RoutingKeyDirect
}

func (p *RabbitMQConfig) GetExchangeDlx() string {
	return p.ExchangeDlx
}

func (p *RabbitMQConfig) GetQueueDlq() string {
	return p.QueueDlq
}

func (p *RabbitMQConfig) GetRoutingKeyDlx() string {
	return p.RoutingKeyDlx
}

func (p *RabbitMQConfig) GetExchangeJunk() string {
	return p.ExchangeJunk
}

func (p *RabbitMQConfig) GetQueueJunk() string {
	return p.QueueJunk
}

func (p *RabbitMQConfig) GetRoutingKeyJunk() string {
	return p.RoutingKeyJunk
}

func (p *RabbitMQConfig) GetTtl() int64 {
	return p.Ttl
}

func (p *RabbitMQConfig) GetDelay() int64 {
	return p.Delay
}

func (p *RabbitMQConfig) GetLimit() int64 {
	return p.Limit
}