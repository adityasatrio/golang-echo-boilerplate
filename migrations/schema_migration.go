package migrations

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
		viper.GetString("db.configs.username"),
		viper.GetString("db.configs.password"),
		viper.GetString("db.configs.host"),
		viper.GetString("db.configs.port"),
		viper.GetString("db.configs.database"))

	log.Println("DSN=", dsn)

	return "mysql://" + dsn
}

func MigrationPath() string {
	//return "file://database/migration"
	return "file://migrations/migration"
}
