# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Philosophy

### THIS IS AN OPINIONATED BOILERPLATE

This boilerplate makes strong architectural decisions for you. These opinions are based on production experience and reflect best practices for building maintainable Go microservices. If you use this boilerplate, you're expected to follow these opinions:

- **MVC + Repository Pattern**: Strict separation of concerns across controller, service, and repository layers
- **Dependency Injection**: Google Wire for compile-time DI (no runtime reflection)
- **ORM with Schema-First**: Ent ORM with code generation from schemas
- **Testing is Mandatory**: Not optional, not negotiable (see below)

### TESTING IS A FIRST-CLASS CITIZEN

**Testing is not optional in this boilerplate. It is MANDATORY and enforced at build time.**

Key principles:
- **Build depends on tests**: The Makefile enforces `build: test` - you CANNOT build without passing tests
- **Every layer must be tested**: Repository (integration), Service (unit with mocks), Controller (HTTP integration)
- **Tests live alongside code**: `*_test.go` files in the same directory as implementation
- **35+ example tests**: Learn from existing patterns in health, system_parameter, user, quotes, and template modules
- **CI enforces testing**: GitHub Actions runs tests on every PR/push to master/release/develop branches

**What this means for you:**
- When you create a new module, you MUST write tests for all three layers before considering it "done"
- The scaffolding template (`make template`) includes test stubs - you're expected to fill them
- Pull requests should include test coverage for new functionality
- `make all` pipeline explicitly runs: schema → mocks → swagger → **TEST** → build → run

This testing-first approach is not negotiable. If you want a boilerplate without mandatory testing, this is not the right choice.

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

**This is an OPINIONATED boilerplate** with mandatory testing and strict architectural patterns. See [Project Philosophy](#project-philosophy) for core principles.

## Common Development Commands

### Build and Run

**IMPORTANT:** Build automatically runs tests first (`build: test` in Makefile). Tests must pass to build.

```bash
make test           # Run tests with coverage (ALWAYS run this first during development)
make build          # Build binary to ./main (depends on 'make test' passing)
./main              # Run the application
make run            # Run the built binary
make all            # Full pipeline: schema → mocks → swagger → TEST → build → run
```

**Pipeline breakdown for `make all`:**
1. `schema-advance`: Generate Ent code from schemas
2. `mocks`: Generate Mockery mocks for testing
3. `swagger`: Generate OpenAPI documentation
4. **`test`**: Run all tests with coverage (MUST PASS)
5. `build`: Compile binary (only if tests pass)
6. `run`: Execute the binary

If tests fail at step 4, the pipeline stops. No build, no run.

### Testing (Mandatory Before Build)

**IMPORTANT:** The build process depends on tests passing. You cannot build without passing tests.

```bash
make test                    # Run all tests with coverage (REQUIRED before build)
go test ./...                # Run tests without coverage report
go test ./internal/applications/health/...  # Run tests for specific module
go test -v ./...             # Verbose output showing each test
```

After running `make test`, two files are generated:
- `coverage.out`: Raw coverage data for tooling
- `coverage.html`: **Visual HTML report - OPEN THIS to review your coverage**

**Build enforcement:**
```bash
make build          # This internally runs 'make test' FIRST
make all            # Full pipeline: schema → mocks → swagger → TEST → build → run
```

If any test fails, the build will abort. This is intentional and ensures code quality.

**Test count:** This boilerplate includes 35+ test files covering all layers. Browse `internal/applications/` for examples.

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

**TESTING REQUIREMENT:** The scaffolded module includes test stubs in:
- `{name}/controller/{name}_controller_test.go`
- `{name}/service/{name}_service_impl_test.go`
- `{name}/repository/db/{name}_repository_impl_test.go`

These are NOT optional placeholders. You MUST implement these tests before considering your module complete. See the [Testing Strategy](#testing-strategy-mandatory) section for patterns and examples.

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

### Testing Strategy (MANDATORY)

Testing is enforced at multiple levels in this boilerplate:

#### Build-Time Enforcement
```makefile
# From Makefile - build explicitly depends on test
build: test
    go build -o main cmd/main.go
```

You CANNOT build the application without all tests passing. This is intentional.

#### Three-Layer Testing Approach

**1. Repository Layer (Integration Tests with Real DB)**
- Use in-memory SQLite via `test.DbConnection(t)`
- Test actual database operations without mocking
- Example: `internal/applications/system_parameter/repository/db/system_parameter_repository_test.go`
- Tests: Create, Update, Delete, SoftDelete, GetById, GetAll, GetByKey operations
- Each test includes cleanup via `t.Cleanup()`

**2. Service Layer (Unit Tests with Mocks)**
- Mock repository and cache dependencies using Mockery-generated mocks
- Use `testify/assert` and `testify/require` for assertions
- Example: `internal/applications/system_parameter/service/system_parameter_service_impl_test.go`
- Pattern: Setup mocks → Configure expectations → Call service → Assert results
- Always test business logic in isolation from infrastructure

**3. Controller Layer (HTTP Integration Tests)**
- Use `httptest.NewRecorder()` and Echo test context
- Mock service layer dependencies
- Example: `internal/applications/health/controller/health_controller_test.go`
- Test HTTP request/response handling, status codes, JSON responses
- Use helper functions like `helper.GetFieldBytes()` to assert response structure

#### Test Infrastructure

**Shared Test Helpers** (`test/` directory):
- `ent_test_helper.go`: In-memory SQLite database setup with `DbConnection(t)` and `DbConnectionTx(t)` for transaction tests
- `echo_test_helper.go`: Echo server initialization for controller tests
- `resty_test_helper.go`: HTTP client setup for integration tests

**Testing Tools:**
- **testify**: Assertions (`assert.Equal`, `assert.NoError`, `assert.NotNil`)
- **Mockery v2.32.0**: Auto-generated mocks for all interfaces (run `make mocks`)
- **In-memory SQLite**: Fast, isolated database tests with `file:ent?mode=memory&cache=shared&_fk=1`
- **miniredis**: In-memory Redis for cache testing
- **httpmock**: HTTP client mocking for outbound API tests

#### Coverage Reports
After running `make test`, two files are generated:
- `coverage.out`: Raw coverage data
- `coverage.html`: Visual HTML report (open in browser)

The Makefile command:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

#### When to Write Tests

**MANDATORY scenarios** (you must write tests for these):
- After scaffolding a new module with `make template name=mymodule`
- When implementing repository layer (`*_repository_impl.go`)
- When implementing service layer (`*_service_impl.go`)
- When implementing controller layer (`*_controller.go`)
- When modifying existing business logic
- When fixing bugs (write a failing test first, then fix)

**Test file naming convention:**
- Repository: `{domain}_repository_impl_test.go`
- Service: `{domain}_service_impl_test.go`
- Controller: `{domain}_controller_test.go`

#### CI/CD Testing
GitHub Actions workflow (`.github/workflows/go.yml`) runs on every push/PR to master, release, and develop branches:
```yaml
- name: Test
  run: go test -v ./...
```

PRs without passing tests will fail CI. This is intentional and enforced.

#### Module Lifecycle Testing Requirements
From the Module Lifecycle section (step 13): **"Write tests for repository, service, and controller"**

This is not optional. A module is not complete until all three layers have test coverage.

### Testing Examples Reference

This boilerplate includes comprehensive test examples across multiple domains. **Study these before writing your own tests:**

#### Repository Layer Tests (Integration with In-Memory DB)
- **System Parameter:** `internal/applications/system_parameter/repository/db/system_parameter_repository_test.go`
  - Covers: Create, Update, Delete, SoftDelete, GetById, GetAll, GetByKey
  - Pattern: Real DB operations, no mocking, cleanup with `t.Cleanup()`

- **Health Check:** `internal/applications/health/repository/health_repository_test.go`
  - Covers: Basic health check repository operations

#### Service Layer Tests (Unit Tests with Mocks)
- **System Parameter:** `internal/applications/system_parameter/service/system_parameter_service_impl_test.go`
  - Pattern: Mock repository + cache, test business logic in isolation
  - Shows: Mock expectations with `.On()`, assertions with `assert.NoError()` and `assert.Equal()`

- **Health Service:** `internal/applications/health/service/health_service_impl_test.go`
  - Pattern: Service layer with multiple dependencies

#### Controller Layer Tests (HTTP Integration)
- **Health Controller:** `internal/applications/health/controller/health_controller_test.go`
  - Pattern: Mock service, use `httptest.NewRecorder()`, assert HTTP responses
  - Shows: Success cases (200), failure cases (500), query parameter handling

- **User Controller:** `internal/applications/user/controller/user_controller_test.go`
  - Pattern: CRUD operations with various HTTP methods

#### Component Tests
- **Cache:** `internal/component/cache/cache_impl_test.go`
  - Pattern: Testing shared components with table-driven tests

- **RabbitMQ Producer:** `internal/component/rabbitmq/producer/producer_service_impl_test.go`
  - Pattern: Testing message broker integration

#### Template Module (Scaffolding Reference)
- **Template Tests:** `internal/applications/A_templates_directory/`
  - Contains test stubs that are copied when you run `make template name=mymodule`
  - These show the expected structure for new modules

**Test file count:** 35+ test files covering all layers and components.

**Best practice:** When implementing a new feature, find the most similar existing module and use its test structure as a template.

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

### Testing Conventions (MANDATORY)

**Test File Placement:**
- Tests MUST live in the same package and directory as the code they test
- Naming: `{implementation_filename}_test.go`
- Example: `user_service_impl.go` → `user_service_impl_test.go`

**Test Function Naming:**
- Format: `Test{TypeName}_{MethodName}` or `Test{TypeName}_{MethodName}_{Scenario}`
- Examples:
  - `TestSystemParameterRepositoryImpl_Create`
  - `TestHealthController_Success`
  - `TestHealthController_Failed`

**Test Structure Pattern:**
```go
func TestMyType_MyMethod(t *testing.T) {
    // 1. Setup: Create test data, mocks, context
    // 2. Execute: Call the method being tested
    // 3. Assert: Verify expected behavior with testify assertions
    // 4. Cleanup: Use t.Cleanup() for resource cleanup
}
```

**Assertion Library:**
- Use `testify/assert` for non-critical assertions
- Use `testify/require` for critical assertions (stops test immediately if fails)
- Examples:
  - `assert.NoError(t, err)` - continue even if fails
  - `require.NoError(t, err)` - stop test if fails
  - `assert.Equal(t, expected, actual)`
  - `assert.NotNil(t, result)`

**Mock Usage:**
- Generate mocks: `make mocks` (after changing interfaces)
- Import from: `myapp/mocks/applications/{domain}/{layer}`
- Setup expectations: `mockRepo.On("MethodName", args...).Return(returnValues...)`
- Verify: Mocks automatically verify expectations were met

**Coverage Expectations:**
While there's no enforced minimum coverage percentage, the expectation is:
- All happy path scenarios are tested
- Critical error paths are tested
- Business logic edge cases are covered

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
13. **WRITE TESTS FOR ALL THREE LAYERS (MANDATORY):**
    - Repository tests: Use `test.DbConnection(t)` for integration tests with in-memory SQLite
    - Service tests: Mock repository and cache using Mockery-generated mocks
    - Controller tests: Use `httptest.NewRecorder()` and mock service layer
    - Run `make test` to verify all tests pass and generate coverage report
    - Open `coverage.html` to review coverage
    - **Your module is NOT complete until tests are written and passing**

## CI/CD

### GitHub Actions Workflow

**File:** `.github/workflows/go.yml`

**Triggers:** Push or Pull Request to `master`, `release`, `develop` branches

**Steps:**
1. Checkout code
2. Setup Go (version from `go.mod`)
3. **Build:** `go build -v ./...`
4. **Test:** `go test -v ./...` (MUST PASS for workflow success)

If tests fail, the workflow fails, and the PR cannot be merged. This is intentional.

**Coverage upload:** Not configured by default. Consider adding coverage reporting (e.g., Codecov, Coveralls) for visibility.

### Pull Request Template

**File:** `pull_request_template.md`

The PR template includes a testing checklist:
```markdown
## How have I tested this
- [ ] I have covered the code with unit tests
- [ ] I have performed manual testing
```

**Both checkboxes should be checked** before requesting review. Code without tests should not be merged.

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
