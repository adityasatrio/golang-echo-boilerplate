package database

import (
	"fmt"

	"entgo.io/ent/dialect"
	"myapp/configs/credential"
)

// SupportedDriver is a database driver supported by the ent client setup.
type SupportedDriver string

const (
	DriverMySQL    SupportedDriver = "mysql"
	DriverPostgres SupportedDriver = "postgres"
)

// ResolveDriver reads db.configs.driver and normalizes it to a SupportedDriver,
// defaulting to mysql when unset to preserve existing behavior.
func ResolveDriver() SupportedDriver {
	driver := credential.GetString("db.configs.driver")

	switch SupportedDriver(driver) {
	case DriverPostgres:
		return DriverPostgres
	case DriverMySQL, "":
		return DriverMySQL
	default:
		return DriverMySQL
	}
}

// BuildDSN returns the sql.Open/ent driver name (e.g. "mysql", "postgres")
// alongside the connection string for the currently configured db.configs.driver.
func BuildDSN() (driverName string, dsn string) {
	return buildDSN(ResolveDriver())
}

func buildDSN(driver SupportedDriver) (string, string) {
	username := credential.GetString("db.configs.username")
	password := credential.GetString("db.configs.password")
	host := credential.GetString("db.configs.host")
	port := credential.GetString("db.configs.port")
	database := credential.GetString("db.configs.database")

	switch driver {
	case DriverPostgres:
		sslMode := credential.GetString("db.configs.sslmode")
		if sslMode == "" {
			sslMode = "disable"
		}
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, username, password, database, sslMode)
		return dialect.Postgres, dsn
	default:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			username, password, host, port, database)
		return dialect.MySQL, dsn
	}
}
