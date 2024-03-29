.PHONY: clean schema-advance mocks wire test build migration run

WIRE_DIR := internal/applications
OPENAPI_ENTRY_POINT := cmd/main.go
OPENAPI_OUTPUT_DIR := cmd/docs

all: schema-advance mocks swagger test build run

OPENAPI_ENTRY_POINT := cmd/main.go
OPENAPI_OUTPUT_DIR := cmd/docs

# Build the project
build: test
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

schema-advance:
	@echo "*****\n Advance mode will granted you a super power, use it wisely\n [Generate with entgo feature sql/modifier,sql/execquery]\n*****"
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier,sql/execquery ./ent/schema

# Generate mockery mocks
mocks:
	mockery --all --dir internal --output mocks --packageprefix mock_ --keeptree

wire:
	@echo "This command will add wire_gen.go in PATH={root}/internal/applications/{your-directory} make sure you already create {domain}_injector.go \nEnter directory: {your-directory} "; \
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

migration-build:
	@echo "Warning this action will build unix executable file "
	go build -v -o migration migrations/cmd/main.go

migration-create:
	./migration mysql create $(name) $(type)

migration-up:
	go build -v -o migration migrations/cmd/main.go
	 ./migration mysql up

migration-down:
	go build -v -o migration migrations/cmd/main.go
	./migration mysql down

migration-status:
	./migration mysql status