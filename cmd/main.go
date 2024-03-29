package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/configs"
	"myapp/configs/cache"
	"myapp/configs/credential"
	"myapp/configs/database"
	"myapp/configs/rabbitmq/connection"
	"myapp/configs/rabbitmq/initialize"
	"myapp/configs/swagger"
	"myapp/configs/validator"
	"myapp/ent"
	restApi "myapp/internal/adapter/rest"
	"myapp/internal/component/rabbitmq/registry"
	"myapp/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "myapp/cmd/docs"
)

//	@title			Micro Go Template Service
//	@version		0.0.1
//	@description	Please welcome a holy high-speed and high-performance Echo service!

//	@contact.url	https://example.com

// @host		localhost:8888
// @basePath	/micro-go-template
func main() {
	e := echo.New()
	fmt.Println("initialized echo framework")

	configs.InitGeneralEnv(e)
	credential.InitCredentialEnv(e)
	configs.SetupLogger(e)

	middleware.SetupMiddlewares(e)

	validator.SetupValidator(e)
	validator.SetupGlobalHttpUnhandleErrors(e)

	dbConnection := database.NewSqlEntClient() //using sqlDb wrapped by ent
	//dbConnection := database.NewEntClient() //using ent only
	log.Info("initialized database configuration=", dbConnection)

	//from docs define close on this function, but will impact cant create DB session on repository:
	defer func(dbConnection *ent.Client) {
		err := dbConnection.Close()
		if err != nil {
			log.Fatalf("error initialized database configuration=", err)
		}
	}(dbConnection)

	//configuration for redis client:
	redisConnection := cache.NewRedisClient()

	//configuration for redis client, for close connection:
	defer func() {
		err := redisConnection.Close()
		if err != nil {
			log.Fatalf("Error closing Redis connection:", err)
		}
	}()

	isQueueEnable := viper.GetBool("rabbitmq.configs.enable")
	var rabbitInit *connection.RabbitMQConnection

	if isQueueEnable {
		//configuration for rabbit client:
		rabbitInit := initialize.RabbitMQInitialize(dbConnection)
		defer func() {
			err := rabbitInit.GetConnection().Close()
			if err != nil {
				log.Errorf("Error closing RabbitMQConnection connection: %v", err)
			}
		}()

		// rabbitmq registry exchange, queue, dlq and other:
		registerMq := registry.NewProducerRegistry(rabbitInit)
		registerMq.Register()

		//rabbitmq registry consumer:
		registerConsumer := registry.NewConsumerRegistry(dbConnection, rabbitInit)
		registerConsumer.Register()
	}

	//setup swagger:
	swagger.InitSwagger()

	//setup router
	restApi.SetupRouteHandler(e, dbConnection, redisConnection, rabbitInit)

	port := viper.GetString("application.port")
	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
			e.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) //not tested
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
