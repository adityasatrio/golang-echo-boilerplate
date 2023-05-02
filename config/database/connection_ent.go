package database

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"myapp/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// NewEntClient use if not wrap sql db
//
//goland:noinspection GoUnusedExportedFunction
func NewEntClient() *ent.Client {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("db.config.username"),
		viper.GetString("db.config.password"),
		viper.GetString("db.config.host"),
		viper.GetString("db.config.port"),
		viper.GetString("db.config.database"))
	fmt.Println("DSN ", dsn)

	client, err := ent.Open("mysql", dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))

	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	//from docs define close on this function, but will impact cant create DB session on repository
	//defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Default().Println("initialized database x orm : success")
	return client
}
