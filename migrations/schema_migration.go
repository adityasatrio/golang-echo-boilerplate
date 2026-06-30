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

// MigrationConnection builds the connection string for the given goose
// driver ("mysql", "postgres" or "sqlite3"), reading credentials from .env.
// It returns the connection string to open the DB with, plus a redacted
// copy (password masked) safe to print in logs/errors.
func MigrationConnection(driver string) (dsn string, redacted string) {
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

	username := viper.GetString("db.configs.username")
	password := viper.GetString("db.configs.password")
	host := viper.GetString("db.configs.host")
	port := viper.GetString("db.configs.port")
	database := viper.GetString("db.configs.database")

	const masked = "****"

	if driver == "postgres" {
		sslMode := viper.GetString("db.configs.sslmode")
		if sslMode == "" {
			sslMode = "disable"
		}
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, username, password, database, sslMode)
		redacted = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, username, masked, database, sslMode)
		return dsn, redacted
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password, host, port, database)
	redacted = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, masked, host, port, database)
	return dsn, redacted
}

func MigrationPath() string {
	return "./migrations"
}
