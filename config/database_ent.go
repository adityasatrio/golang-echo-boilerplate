package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"myapp/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}

func NewEntClient() (*ent.Client, error) {

	/*dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
	os.Getenv("DB_USERNAME"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_DATABASE"))*/

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("db.config.username"),
		viper.GetString("db.config.password"),
		viper.GetString("db.config.host"),
		viper.GetString("db.config.port"),
		viper.GetString("db.config.database"))

	fmt.Print("DSN ", dsn)
	//client, err := ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	client, err := ent.Open("mysql", dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))
	if err != nil {
		//log.Fatalf("failed opening connection to sqlite: %v", err)
		fmt.Print("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		//log.Fatalf("failed creating schema resources: %v", err)
		fmt.Print("failed creating schema resources: %v", err)
	}

	return client, err
}
