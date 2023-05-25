.PHONY: all build test generate-ent generate-mocks clean


# Build the project
build:
	rm -f main
	go build -o main cmd/main.go

# Run tests and generate coverage report
run:
	rm -f main
	go build -o main cmd/main.go
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
#
## Clean generated files
#clean:
#	echo "CAUTION THIS WILL REMOVE coverage.out coverage.html !"
#	delay 10
#	rm -f coverage.out coverage.html
#	echo "CAUTION THIS WILL REMOVE /ENT FOLDER !"
#	delay 10
#	find ./ent/* -mindepth 1 -type d ! -name "schema" -exec rm -rf {}/\* \;
#	find ./ent -type f ! -name "generate.go" -exec rm -f {} \;
#
#	#rm -rf ./ent/*
#	#rm -rf ./mocks/*

all: gen-model gen-mocks test build
