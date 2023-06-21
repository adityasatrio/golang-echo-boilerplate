# Nama file Makefile

# Direktori migrasi
MIGRATE_DIR := database/migration

.PHONY: migration
create-migration:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(MIGRATION_NAME)

run-migration:
	 go run database/executor/migration_up/main.go

rollback-migration:
	go run database/executor/migration_down/main.go $(MIGRATION_VERSION)

force-migration:
	go run database/executor/migration_force/main.go $(MIGRATION_VERSION)