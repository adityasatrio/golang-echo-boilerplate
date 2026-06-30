package database

import (
	"context"
	"github.com/labstack/gommon/log"
	"myapp/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// NewEntClient use if not wrap sql db
//
//goland:noinspection GoUnusedExportedFunction
func NewEntClient() *ent.Client {

	driverName, dsn := BuildDSN()

	log.Debugf("DB driver=", driverName)

	client, err := ent.Open(driverName, dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			log.Debugf(time.Now().Format("2006-01-02 15:04:05"), v)
		}
	}))

	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Infof("initialized database x orm : success")
	return client
}
