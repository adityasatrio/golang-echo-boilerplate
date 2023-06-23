package database

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

const (
	fileConfigPath = "."
	fileConfigType = "yml"
	fileConfigName = "app.config"
)

func MigrationConnection() string {

	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigType(fileConfigType)
	viper.SetConfigName(fileConfigName)

	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("db.config.username"),
		viper.GetString("db.config.password"),
		viper.GetString("db.config.host"),
		viper.GetString("db.config.port"),
		viper.GetString("db.config.database"))

	log.Println("DSN=", dsn)

	return "mysql://" + dsn
}

func MigrationPath() string {
	return "file://database/migration"
}
