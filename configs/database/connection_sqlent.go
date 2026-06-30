package database

import (
	"database/sql"
	entSql "entgo.io/ent/dialect/sql"
	"github.com/labstack/gommon/log"
	"myapp/configs/credential"
	"myapp/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func NewSqlEntClient() *ent.Client {

	driverName, dsn := BuildDSN()

	log.Debug("DB driver=", driverName) // for debug only

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	db.SetMaxIdleConns(credential.GetInt("db.configs.maxIdleConn"))
	db.SetMaxOpenConns(credential.GetInt("db.configs.maxOpenConn"))
	db.SetConnMaxLifetime(time.Hour)

	drv := entSql.OpenDB(driverName, db)

	var client *ent.Client
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
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	if err != nil {
		log.Printf("err : %s\n", err)
	}

	log.Info("initialized database x sqlDb x orm ent : success")

	// setup hooks
	SetupHooks(client)

	return client
}
