.PHONY: clean gen-schema gen-mocks test build migration run

# Directory migration:
MIGRATE_DIR := migrations/migration

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
schema-gen:
	go generate ./ent

# Generate mockery mocks
mocks-gen:
	mockery --all --dir internal/applications --output mocks --packageprefix mock_ --keeptree

confirm:
	@read -p "$(shell echo -e '\033[0;31m')Warning: This action will clean up coverage reports, ent schema, and mockery generated codes. Do you want to continue? [Y/n]: $(shell tput sgr0)" choice; \
	if [ "$$choice" != "Y" ]; then \
		echo "$(shell echo -e '\033[0;31m')Terminating the clean-up process.$(shell output sgr0)"; \
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

migration-create:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(name)

migration-up:
	 go run migrations/cmd/main.go -type up

migration-down:
	go run migrations/cmd/main.go -type down -version $(version)

migration-force:
	go run migrations/cmd/main.go -type force -version $(version)

all: schema-gen mocks-gen test build run



