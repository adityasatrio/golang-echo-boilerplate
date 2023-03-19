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
	log.Println("DSN=", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	drv := entsql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv))

	if drv == nil || client == nil {
		log.Fatalf("failed opening connection to DB : driver or DB new client is null")
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err != nil {
		log.Println("err : %s", err)
	}

	//from docs define close on this function, but will impact cant create DB session on repository
	//defer client.Close()

	log.Default().Println("initialized database x sqldb x orm ent : success")
	return client
}
