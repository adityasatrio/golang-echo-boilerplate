package database

import (
	"github.com/labstack/gommon/log"
	"myapp/ent"
	"myapp/ent/hook"
)

func SetupHooks(dbConnection *ent.Client) {
	dbConnection.Use(hook.VersionHook())

	log.Info("initialized SetupHooks configuration=")
}
