# Makefile

# Directory for migrasi file
MIGRATE_DIR := database/migration

.PHONY: migration
migration-create:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(name)

migration-up:
	 go run database/cmd/main.go -type up

migration-down:
	go run database/cmd/main.go -type down -version $(version)

migration-force:
	go run database/cmd/main.go -type force -version $(version)