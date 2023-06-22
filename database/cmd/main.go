package main

import (
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"myapp/database"
)

type MigrationType string

const (
	Up    = "up"
	Down  = "down"
	Force = "force"
)

func main() {

	//schema migration & file path:
	migrationsPath := database.MigrationPath()
	dbConnString := database.MigrationConnection()

	//this for capture argument:
	migrationType := flag.String("type", "no-type", "type your migration")
	migrationVersion := flag.Int64("version", 0, "version your migration")
	flag.Parse()

	if *migrationType == Up {
		executeMigrationUp(migrationsPath, dbConnString, Up)
	} else if *migrationType == Down {
		executeMigration(migrationsPath, dbConnString, Down, int(*migrationVersion))
	} else if *migrationType == Force {
		executeMigration(migrationsPath, dbConnString, Force, int(*migrationVersion))
	} else {
		log.Println("please use arguments to run the migration you need")
	}
}

func executeMigrationUp(migrationsPath, dbConnString string, migrationType MigrationType) {

	m, err := migrate.New(migrationsPath, dbConnString)
	if err != nil {
		log.Fatalf("failed for migration %s creating schema resources: %v", migrationType, err)
		return
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Println("Error running migrations UP:", err)
		return
	}

	log.Printf("Migrations %s completed successfully.", migrationType)
}

func executeMigration(migrationsPath, dbConnString string, migrationType MigrationType, version int) {

	if version == 0 {
		log.Printf("the arguments version given are incorrect for migration %s", migrationType)
		return
	}

	m, err := migrate.New(migrationsPath, dbConnString)
	if err != nil {
		log.Fatalf("failed for migration %s creating schema resources: %v", migrationType, err)
		return
	}

	//this for migration type DOWN:
	if migrationType == Down {
		err = m.Migrate(uint(version))
		if err != nil && err != migrate.ErrNoChange {
			log.Println("Error running migrations Down:", err)
			return
		}
	}

	//this for migration type FORCE:
	if migrationType == Force {
		err = m.Force(version)
		if err != nil && err != migrate.ErrNoChange {
			log.Println("Error running migrations Force:", err)
			return
		}
	}

	log.Printf("Migrations %s completed successfully.", migrationType)
}
