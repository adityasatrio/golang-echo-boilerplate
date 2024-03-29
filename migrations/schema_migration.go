package migrations

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

const (
	fileConfigType = ".env"
)

func MigrationConnection() string {
	viper.SetConfigFile(fileConfigType)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString("db.configs.username"),
		viper.GetString("db.configs.password"),
		viper.GetString("db.configs.host"),
		viper.GetString("db.configs.port"),
		viper.GetString("db.configs.database"))
	return dsn
}

func MigrationPath() string {
	return "./migrations"
}
