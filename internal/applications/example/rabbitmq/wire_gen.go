// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rabbitmq

import (
	"github.com/google/wire"
	"myapp/configs/rabbitmq/connection"
	"myapp/ent"
	"myapp/internal/applications/example/rabbitmq/repository/inbound"
	"myapp/internal/applications/example/rabbitmq/service"
	"myapp/internal/applications/system_parameter/repository/db"
	"myapp/internal/component/rabbitmq/channel"
	"myapp/internal/component/rabbitmq/producer"
)

// Injectors from example_rabbitmq_injector.go:

func InitializedExampleInbound(dbClient *ent.Client, conn *connection.RabbitMQConnection) *inbound.ExampleRabbitMQInboundImpl {
	wrappedChannelServiceImpl := channel.NewWrappedChannel(conn)
	systemParameterRepositoryImpl := db.NewSystemParameterRepository(dbClient)
	exampleRabbitMQServiceImpl := service.NewExampleRabbitMQService(systemParameterRepositoryImpl)
	producerServiceImpl := producer.NewProducerService(wrappedChannelServiceImpl)
	exampleRabbitMQInboundImpl := inbound.NewExampleRabbitMQInbound(wrappedChannelServiceImpl, exampleRabbitMQServiceImpl, producerServiceImpl)
	return exampleRabbitMQInboundImpl
}

func InitializedExampleService(dbClient *ent.Client) *service.ExampleRabbitMQServiceImpl {
	systemParameterRepositoryImpl := db.NewSystemParameterRepository(dbClient)
	exampleRabbitMQServiceImpl := service.NewExampleRabbitMQService(systemParameterRepositoryImpl)
	return exampleRabbitMQServiceImpl
}

// example_rabbitmq_injector.go:

var providerExampleInbound = wire.NewSet(inbound.NewExampleRabbitMQInbound, db.NewSystemParameterRepository, channel.NewWrappedChannel, service.NewExampleRabbitMQService, producer.NewProducerService, wire.Bind(new(inbound.ExampleRabbitMQInbound), new(*inbound.ExampleRabbitMQInboundImpl)), wire.Bind(new(db.SystemParameterRepository), new(*db.SystemParameterRepositoryImpl)), wire.Bind(new(channel.WrappedChannelService), new(*channel.WrappedChannelServiceImpl)), wire.Bind(new(service.ExampleRabbitMQService), new(*service.ExampleRabbitMQServiceImpl)), wire.Bind(new(producer.Producer), new(*producer.ProducerServiceImpl)))

var providerExampleService = wire.NewSet(service.NewExampleRabbitMQService, db.NewSystemParameterRepository, wire.Bind(new(db.SystemParameterRepository), new(*db.SystemParameterRepositoryImpl)), wire.Bind(new(service.ExampleRabbitMQService), new(*service.ExampleRabbitMQServiceImpl)))
