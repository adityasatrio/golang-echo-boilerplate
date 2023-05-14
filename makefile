.PHONY: all build test generate-ent generate-mocks clean


# Build the project
build:
	go build -o cmd/main cmd/main.go

# Run tests and generate coverage report
run:
	rm -f ./cmd/main
	#go test -coverprofile=coverage.out ./...
	go build -o cmd/main cmd/main.go
	./cmd/main

# Run tests and generate coverage report
test:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Generate ent models
gen-ent:
	go generate ./ent

# Generate mockery mocks
gen-mocks:
	mockery --all --dir internal/applications --output mocks --keeptree --packageprefix mock_

# Clean generated files
clean:
	rm -f coverage.out coverage.html
	rm -rf ./ent/*
	rm -rf ./mocks/*

all: build test generate-ent generate-mocks
