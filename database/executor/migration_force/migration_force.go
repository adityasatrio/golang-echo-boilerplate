package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"myapp/database"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	var version = 0
	if len(args) > 0 {
		version, _ = strconv.Atoi(args[0])
	} else {
		log.Println("No arguments are given.")
		return
	}
	dbConnString := database.MigrationConnection()
	migrationsPath := database.MigrationPath()

	//this for run executor type force:
	runMigrationDown(migrationsPath, dbConnString, version)
}

func runMigrationDown(migrationsPath, dbConnString string, version int) {

	if version == 0 {
		log.Println("the arguments given are incorrect")
		return
	}

	m, err := migrate.New(migrationsPath, dbConnString)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return
	}

	err = m.Force(version)
	if err != nil && err != migrate.ErrNoChange {
		log.Println("Error running migrations:", err)
		return
	}

	log.Println("Migrations completed successfully.")
}
