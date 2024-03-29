.PHONY: clean schema mocks wire test build migration run

MIGRATE_DIR := migrations/migration
WIRE_DIR := internal/applications

OPENAPI_ENTRY_POINT := cmd/main.go
OPENAPI_OUTPUT_DIR := cmd/docs

# Build the project
build:
	go build -o main cmd/main.go

# Run tests and generate coverage report
run:
	./main

# Run tests and generate coverage report
test:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Generate ent models
schema:
	go generate ./ent

# Generate mockery mocks
mocks:
	mockery --all --dir internal/applications --output mocks --packageprefix mock_ --keeptree

wire:
	@echo "Enter directory: "; \
	read dir; \
	echo "Accessing directory and wire all DI $(WIRE_DIR)/$$dir"; \
	cd $(WIRE_DIR)/$$dir && wire

# Generate OpenAPI Docs
swagger:
	swag fmt && swag init -g $(OPENAPI_ENTRY_POINT) -o $(OPENAPI_OUTPUT_DIR)

confirm:
	@read -p "$(shell echo -e '\033[0;31m') Warning: This action will clean up coverage reports, ent schema, and mockery generated codes. Do you want to continue? [Y/n]: $(shell tput sgr0)" choice; \
	if [ "$$choice" != "Y" ]; then \
		echo "$(shell echo -e '\033[0;31m') Terminating the clean-up process.$(shell output sgr0)"; \
    	exit 1; \
    fi

clean: confirm
	@echo "Warning this action will clean-up coverage report, ent schema and mockery generated codes "
	sleep 10

	@echo "Deleting coverage.out coverage.html on 5s"
	sleep 5
	rm -f coverage.out coverage.html

	@echo "Deleting all directories and files in ./ent except ./ent/schema and ./ent/generate.go on 5s"
	sleep 5
	@find ./ent/* ! -path "./ent/schema*" ! -path "./ent/generate.go" ! -path "./ent/hook*" -delete

	@echo "Deleting all directories and files in /.mocks on 5s"
	sleep 5
	rm -rf ./mocks/*

all: schema mocks test build run

migration-create:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(name)

migration-up:
	 go run database/cmd/main.go -type up

migration-down:
	go run database/cmd/main.go -type down -version $(version)

migration-force:
	go run database/cmd/main.go -type force -version $(version)