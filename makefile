.PHONY: all build test generate-ent generate-mocks clean


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
gen-model:
	go generate ./ent

# Generate mockery mocks
gen-mocks:
	mockery --all --dir internal/applications --output mocks --keeptree --packageprefix mock_

clean:
	@echo "Deleting coverage.out coverage.html on 5s"
	sleep 5
	rm -f coverage.out coverage.html

	@echo "Deleting all directories and files in ./ent except ./ent/schema and ./ent/generate.go on 10s"
	sleep 10
	@find ./ent/* ! -path "./ent/schema*" ! -path "./ent/generate.go" -delete

	@echo "Deleting all directories and files in /.mocks on 10s"
	sleep 10
	rm -rf ./mocks/*

all: gen-model gen-mocks test build
