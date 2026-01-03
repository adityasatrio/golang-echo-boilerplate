.PHONY: clean schema-advance mocks wire test build migration run template

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

# Scaffold a new module from A_templates_directory
# Usage:
#   make template name=<module_name>
# Prompts for name if not provided. Creates internal/applications/<module_name>
# and performs placeholder replacements (imports, identifiers, routes).
template:
	@set -e; \
	if [ -z "$(name)" ]; then read -p "Enter module name (e.g., user): " name; else name="$(name)"; fi; \
	newdir="internal/applications/$$name"; \
	if [ -d "$$newdir" ]; then echo "Module directory already exists: $$newdir"; exit 1; fi; \
	echo "Scaffolding module at $$newdir from A_templates_directory"; \
	cp -R internal/applications/A_templates_directory "$$newdir"; \
	rm -f "$$newdir/wire_gen.go" || true; \
	Pascal=$$(printf '%s' "$$name" | awk '{print toupper(substr($$0,1,1)) substr($$0,2)}'); \
	find "$$newdir" -type f \( -name '*.go' -o -name '*.MD' \) -print0 | xargs -0 sed -i -e "s|A_templates_directory|$$name|g" -e "s|/A_templates_directory|/$$name|g"; \
	find "$$newdir" -type f -name '*.go' -print0 | xargs -0 sed -i -e "s|Template|$$Pascal|g" -e "s|template|$$name|g"; \
	if [ -f "$$newdir/template_injector.go" ]; then sed -i "s/^package .*/package $$name/" "$$newdir/template_injector.go"; mv "$$newdir/template_injector.go" "$$newdir/$${name}_injector.go"; fi; \
	for f in $$(find "$$newdir" -type f -name 'template_*'); do newf=$$(echo "$$f" | sed "s/\/template_/\/$${name}_/g"); mv "$$f" "$$newf"; done; \
	bash scripts/adjust_routes.sh "$$name"; \
	echo "Done. Next steps:"; \
	echo "  - Update routes in $$newdir/controller to use your desired path (currently '/$$name')."; \
	echo "  - Run: make wire (select '$$name') and make mocks"; \
	echo "  - Integrate routes in internal/adapter/rest/routes_setup.go (match your AddRoutes signature)."

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
