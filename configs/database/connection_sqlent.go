package database

import (
	"database/sql"
	entSql "entgo.io/ent/dialect/sql"

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
		viper.GetString("db.configs.username"),
		viper.GetString("db.configs.password"),
		viper.GetString("db.configs.host"),
		viper.GetString("db.configs.port"),
		viper.GetString("db.configs.database"))
	log.Println("DSN=", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	drv := entSql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv))

	if drv == nil || client == nil {
		log.Fatalf("failed opening connection to DB : driver or DB new client is null")
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err != nil {
		log.Printf("err : %s\n", err)
	}

	//from docs define close on this function, but will impact cant create DB session on repository
	//defer client.Close()

	//result := Client{Connection: *client}
	log.Default().Println("initialized database x sqlDb x orm ent : success")
	return client
}
