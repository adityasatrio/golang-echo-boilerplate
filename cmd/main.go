package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/configs"
	"myapp/configs/database"
	"myapp/configs/validator"
	"myapp/ent"
	restApi "myapp/internal/adapter/rest"
	"myapp/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()

	configs.SetupConfigEnv(e)
	configs.SetupLogger(e)
	middleware.SetupMiddlewares(e)
	validator.SetupValidator(e)
	validator.SetupGlobalHttpUnhandleErrors(e)

	dbConnection := database.NewSqlEntClient() //using sqlDb wrapped by ent
	//dbConnection := database.NewEntClient() //using ent only
	database.SetupHooks(dbConnection)

	log.Info("initialized database configuration=", dbConnection)
	//from docs define close on this function, but will impact cant create DB session on repository
	defer func(dbConnection *ent.Client) {
		err := dbConnection.Close()
		if err != nil {
			log.Fatalf("error initialized database configuration=", err)
		}
	}(dbConnection)

	//setup router
	restApi.SetupRouteHandler(e, dbConnection)

	port := viper.GetString("application.port")
	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
			e.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
