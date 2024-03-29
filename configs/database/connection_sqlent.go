package database

import (
	"database/sql"
	entSql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	"myapp/configs/credential"
	"myapp/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewSqlEntClient() *ent.Client {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		credential.GetString("db.configs.username"),
		credential.GetString("db.configs.password"),
		credential.GetString("db.configs.host"),
		credential.GetString("db.configs.port"),
		credential.GetString("db.configs.database"))

	log.Debug("DSN=", dsn) //for debug only

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	db.SetMaxIdleConns(credential.GetInt("db.configs.maxIdleConn"))
	db.SetMaxOpenConns(credential.GetInt("db.configs.maxOpenConn"))
	db.SetConnMaxLifetime(time.Hour)

	drv := entSql.OpenDB("mysql", db)

	var client = &ent.Client{}
	appMode := credential.GetString("application.mode")
	if appMode == "prod" {
		log.Info("initialized database x sqlDb x orm ent : DEV")
		client = ent.NewClient(ent.Driver(drv))
	} else {
		log.Info("initialized database x sqlDb x orm ent : PROD")
		client = ent.NewClient(ent.Driver(drv), ent.Debug())
	}

	if drv == nil || client == nil {
		log.Fatalf("failed opening connection to DB : driver or DB new client is null")
	}

	// Run the auto migration tool.
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	if err != nil {
		log.Printf("err : %s\n", err)
	}

	log.Info("initialized database x sqlDb x orm ent : success")

	//setup hooks
	SetupHooks(client)

	return client
}
