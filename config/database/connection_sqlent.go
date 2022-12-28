package database

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"

	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"myapp/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewSqlEntClient() *ent.Client {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("db.config.username"),
		viper.GetString("db.config.password"),
		viper.GetString("db.config.host"),
		viper.GetString("db.config.port"),
		viper.GetString("db.config.database"))

	log.Print("DSN ", dsn)

	db, err := sql.Open("mysql", "<mysql-dsn>")
	if err != nil {
		return nil
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	drv := entsql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv))

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		fmt.Print("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		fmt.Print("failed creating schema resources: %v", err)
	}

	if err != nil {
		log.Printf("err : %s", err)
	}

	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Printf("err : %s", err)
		}
	}(client)

	if err != nil {
		log.Println("Fail to initialize client")
	}

	return client
}
