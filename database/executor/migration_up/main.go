package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"myapp/database"
)

func main() {

	dbConnString := database.MigrationConnection()
	migrationsPath := database.MigrationPath()

	//this for run executor type up:
	runMigrationUp(migrationsPath, dbConnString)
}

func runMigrationUp(migrationsPath, dbConnString string) {
	m, err := migrate.New(migrationsPath, dbConnString)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println("Error running migrations:", err)
		return
	}

	fmt.Println("Migrations completed successfully.")
}
