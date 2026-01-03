# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go microservice boilerplate using Echo framework (v4), Ent ORM, Redis caching, and RabbitMQ messaging. The architecture follows MVC pattern with repository pattern, dependency injection via Google Wire, and comprehensive testing with Mockery.

**Key Technologies:**
- Go 1.22
- Echo v4 (HTTP framework)
- Ent (ORM with schema-first approach)
- Google Wire (DI code generation)
- Redis (caching layer)
- RabbitMQ (optional message broker)
- Viper (configuration)
- Swagger/OpenAPI (API documentation)
- Goose (database migrations)

## Common Development Commands

### Build and Run
```bash
make build          # Build binary to ./main
./main              # Run the application
make run            # Run the built binary
make all            # Full cycle: schema, mocks, swagger, test, build, run
```

### Testing
```bash
make test                    # Run all tests with coverage
go test ./...                # Run tests without coverage
go test ./internal/applications/health/...  # Run tests for specific module
```

After running `make test`, open `coverage.html` in your browser to view coverage report.

### Ent Schema Management
```bash
go run entgo.io/ent/cmd/ent init ModelName    # Create new schema
make schema                                    # Generate Ent code (basic)
make schema-advance                            # Generate with sql/modifier and sql/execquery features
go generate ./ent                              # Alternative to make schema
```

**Important:** Only edit files in `ent/schema/`. All other files in `ent/` are generated.

### Database Migrations
```bash
make migration-create name=add_users_table type=sql   # Create SQL migration
make migration-create name=seed_data type=go          # Create Go migration
make migration-up                                      # Apply migrations
make migration-down                                    # Rollback last migration
make migration-status                                  # Check migration status
```

### Dependency Injection (Wire)
```bash
make wire           # Interactive: prompts for module name
# When prompted, enter the module directory name (e.g., "system_parameter")
```

This generates `wire_gen.go` in the specified module. Requires `{module}_injector.go` to exist first.

### Mocking
```bash
make mocks          # Generate mocks for all interfaces in internal/
```

Regenerate mocks whenever you modify interface signatures.

### Swagger Documentation
```bash
make swagger        # Generate OpenAPI docs from controller annotations
```

Access Swagger UI at: `http://localhost:8888/micro-go-template/swagger/index.html` (adjust port/path based on `.env`)

### Scaffolding New Modules
```bash
make template name=products    # Create new module from template
```

This copies `A_templates_directory`, renames files/identifiers, and updates routing. After scaffolding:
1. Run `make wire` (select your new module)
2. Run `make mocks`
3. Manually integrate routes in `internal/adapter/rest/routes_setup.go`

## Architecture and Code Structure

### Request Flow
```
HTTP Request → Echo Router → Controller → Service → Repository → Ent (Database)
                                       ↓
                                    Cache (Redis)
```

### Module Structure
Each domain module lives in `internal/applications/{domain}/`:
```
{domain}/
├── controller/          # HTTP handlers and route definitions
│   ├── {domain}_controller.go
│   ├── {domain}_routes.go
│   └── {domain}_controller_test.go
├── service/             # Business logic layer
│   ├── {domain}_service.go (interface)
│   ├── {domain}_service_impl.go
│   └── {domain}_service_impl_test.go
├── repository/          # Data access layer
│   ├── db/              # Database operations (Ent)
│   │   ├── {domain}_repository.go
│   │   ├── {domain}_repository_impl.go
│   │   └── {domain}_repository_impl_test.go
│   ├── outbound/        # External API calls
│   └── inbound/         # RabbitMQ consumers
├── dto/                 # Data transfer objects
│   └── {domain}_dto.go
├── {domain}_injector.go # Wire DI configuration
└── wire_gen.go          # Generated Wire code (DO NOT EDIT)
```

### Key Directories
- `cmd/main.go`: Application entry point and server setup
- `cmd/docs/`: Generated Swagger documentation (DO NOT EDIT)
- `configs/`: Configuration loaders (Viper), database/cache/RabbitMQ connections, validators, logging
- `ent/schema/`: Ent ORM schemas (ONLY PLACE to manually edit in ent/)
- `ent/hook/`: Custom Ent hooks (e.g., version_hook.go for optimistic locking)
- `exceptions/`: Custom error types (BusinessLogicError with error codes)
- `middleware/`: Echo and Resty middleware configuration
- `migrations/`: Goose migration files and CLI tool
- `mocks/`: Generated mock implementations (DO NOT EDIT)
- `internal/adapter/rest/`: Centralized route registration in `routes_setup.go`
- `internal/component/`: Shared components (cache, rabbitmq, transaction helpers)
- `internal/vars/`: Application info constants

### Configuration Management
Configuration is split across two files:
- `.env`: Application config (port, feature flags, cache TTLs, RabbitMQ topology)
- `secret.env`: Credentials (database, Redis, RabbitMQ passwords)

Override credential file location: `./main -credentials-path=/custom/path -credentials-name=prod-secrets`

**Important config keys:**
- `application.name`: Must start with `/` (e.g., `/micro-go-template`)
- `application.port`: HTTP port (default: 8888)
- `rabbitmq.configs.enable`: Set to `false` to disable RabbitMQ entirely
- `swagger.host`: Must match `localhost:{port}` for correct Swagger behavior

### Dependency Injection Pattern
This project uses Google Wire for compile-time DI. Each module has:
1. `{module}_injector.go`: Wire provider sets and initialization function
   ```go
   //go:build wireinject
   var providerSet = wire.NewSet(
       repository.New...,
       service.New...,
       wire.Bind(new(Interface), new(*Implementation)),
   )
   func InitializedService(...) *ServiceImpl {
       wire.Build(providerSet)
       return nil
   }
   ```
2. `wire_gen.go`: Generated by Wire (run `make wire`)
3. Registration in `internal/adapter/rest/routes_setup.go`:
   ```go
   svc := module.InitializedService(dbClient, redisClient)
   controller.NewController(svc).AddRoutes(e, appName)
   ```

### Testing Strategy
- Tests live alongside source files as `*_test.go`
- Use `testify/assert` for assertions
- Use Mockery-generated mocks for interface dependencies
- Shared test helpers in `test/` directory (Echo server setup, Resty client, Ent test DB)
- Example patterns in `internal/applications/health/` and `internal/applications/system_parameter/`

### Ent ORM Key Features
**Base Mixins** (ent/schema/base_mixin.go):
- `BaseFieldMixin`: Adds audit fields (created_by, created_at, updated_by, updated_at, deleted_by, deleted_at)
- `VersionMixin`: Optimistic locking via `versions` field (Unix timestamp)

**Optimistic Locking:** Enabled by VersionMixin + hook in `ent/hook/version_hook.go`. Updates fail if version doesn't match, preventing concurrent modification issues.

**Schema Definition:** Create schemas in `ent/schema/`, then run `make schema-advance` to generate. Use advanced mode for raw SQL modifier support.

**Transactions:** Available via `ent.Client.Tx()` - see `internal/applications/user/service/` for examples.

### Error Handling
Custom errors in `exceptions/`:
- `BusinessLogicError`: Wraps application errors with error codes
- `err_enum.go`: Error code definitions
- Global error handler in `configs/validator/response_error_handler.go`
- Validation errors mapped via `response_error_mapper.go`

### Caching Layer
Redis caching via `internal/component/cache/`:
- `caching_service.go`: Cache interface
- LZ4 compression for large values
- TTL presets in `.env`: `cache.ttl.short`, `cache.ttl.medium`, `cache.ttl.long`

### RabbitMQ Integration (Optional)
Controlled by `rabbitmq.configs.enable` in `.env`:
- Producers: `internal/component/rabbitmq/producer/`
- Consumers: `internal/component/rabbitmq/consumer/`
- Registry: `internal/component/rabbitmq/registry/` (exchange/queue/DLQ setup)
- Recovery: Auto-reconnection configured in `configs/rabbitmq/recovery/`
- Example implementation: `internal/applications/example/rabbitmq/`

## Important Patterns and Conventions

### Naming Conventions
- Packages: lowercase (e.g., `system_parameter`)
- Files: snake_case (e.g., `user_service.go`)
- Exported identifiers: PascalCase
- Unexported identifiers: camelCase
- Test files: `*_test.go` with `TestXxx` functions

### Generated Files - DO NOT EDIT
- `wire_gen.go` in each module
- Everything in `ent/` except `ent/schema/` and `ent/hook/`
- Everything in `mocks/`
- Everything in `cmd/docs/`

### Adding New Endpoints
1. Annotate controller method with Swagger comments
2. Run `make swagger` to update docs
3. Restart application to load new routes

### Controller Swagger Annotation Example
```go
// CreateUser godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "User data"
// @Success 201 {object} ent.User
// @Router /users [post]
func (c *UserController) Create(ctx echo.Context) error { ... }
```

### Module Lifecycle
When creating a new feature module:
1. Scaffold: `make template name=mymodule`
2. Define Ent schema: `go run entgo.io/ent/cmd/ent init MyModel`
3. Edit `ent/schema/mymodel.go`, run `make schema-advance`
4. Implement DTOs in `mymodule/dto/`
5. Implement repository (interface + impl) in `mymodule/repository/db/`
6. Implement service (interface + impl) in `mymodule/service/`
7. Implement controller + routes in `mymodule/controller/`
8. Create `mymodule_injector.go` with Wire provider set
9. Run `make wire`, select `mymodule`
10. Run `make mocks`
11. Register routes in `internal/adapter/rest/routes_setup.go`
12. Run `make swagger` if adding HTTP endpoints
13. Write tests for repository, service, and controller

## CI/CD

GitHub Actions workflow (`.github/workflows/go.yml`):
- Triggers on push/PR to `master`, `release`, `develop`
- Steps: checkout, setup Go, build (`go build -v ./...`), test (`go test -v ./...`)
- No coverage upload configured by default

## Local Development Setup

1. Copy and configure environment files:
   ```bash
   cp .env.example .env         # If exists, otherwise edit .env directly
   cp secret.env.example secret.env
   ```

2. Ensure database is accessible (MySQL/PostgreSQL):
   ```
   db.configs.username=root
   db.configs.password=password
   db.configs.host=127.0.0.1
   db.configs.port=3306
   db.configs.database=echo_sample
   ```

3. Run migrations:
   ```bash
   make migration-up
   ```

4. Generate Ent schemas and mocks:
   ```bash
   make schema-advance
   make mocks
   ```

5. Build and run:
   ```bash
   make build
   ./main
   ```

6. Access application:
   - Health check: `http://localhost:8888/micro-go-template/health`
   - Swagger UI: `http://localhost:8888/micro-go-template/swagger/index.html`

## Troubleshooting Common Issues

**Swagger not loading:** Check that `swagger.host` in `.env` matches your actual host and port.

**RabbitMQ connection errors:** Set `rabbitmq.configs.enable=false` in `.env` if not using RabbitMQ.

**Wire generation fails:** Ensure `{module}_injector.go` exists and has correct `//go:build wireinject` tag.

**Ent query fails:** Verify you ran `make schema-advance` after modifying schemas.

**Mock out of sync:** Regenerate with `make mocks` after changing interfaces.

**Optimistic lock errors:** This is expected behavior when concurrent updates occur - handle with retry logic or inform user.
