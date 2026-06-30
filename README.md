# Micro Go Template
Boilerplate template for backend project using go language for optimize and efficient resources.

## Contents
- [Getting started](#getting-started)
  - [Install and generate ORM](#install-and-generate-orm)
  - [Install and generate mock using Mockery](#install-and-generate-mock-using-mockery)
  - [Dependency Injection](#dependency-injection)
  - [OpenAPI Docs and Swagger](#openapi-docs-and-swagger)
  - [Migration](#migration)
  - [Switching database driver (MySQL / PostgreSQL)](#switching-database-driver)
  - [Seeders](#seeders)
  - [Running with config](#run-config)
- [Example code](#example-code)
- [Build project](#build-project)
- [References project](#references-project)
- [List of project features](#list-of-project-features)

## Getting started

### Install and generate ORM
First before we run the application, lets fresh create model schema using ent:
1. Get dependency for golang ent
```shell
go get entgo.io/ent/cmd/ent
```
2. Create new model schema, the generated model located on `ent/schema/model_name.go` </br>
   https://entgo.io/docs/schema-def for schema documentation
```shell
go run entgo.io/ent/cmd/ent init {model_name}
```
3. Generate assets
```shell
go generate ./ent
```

### Install and generate mock using Mockery
generate mock file for all interface in domains. Install mockery first
```shell
go install github.com/vektra/mockery/v2@v2.32.0
```
then generate mock using below command
```shell
 mockery --all --dir internal/applications --output mocks --keeptree --packageprefix mock_
``` 

### Dependency Injection

This project uses a **fluent builder pattern** for dependency injection instead of code generation.

**Setup** - The builder container is initialized in `cmd/main.go`:
```go
container := builder.NewBuilder().
    WithDatabase(dbClient).
    WithCache(redisClient).
    WithRabbit(rabbitConn)

restApi.SetupRouteHandler(e, container)
```

**Adding a New Service** - When creating a new feature domain:
1. Create your service interface and implementation in `internal/applications/{domain}/service/`
2. Create your repository interface and implementation in `internal/applications/{domain}/repository/`
3. Add a builder method in `internal/builder/services.go`:
```go
func (c *Container) BuildMyDomainService() mydomain.Service {
    repo := repository.NewMyDomainRepository(c.db)
    return service.NewMyDomainService(repo)
}
```
4. Register your controller routes in `internal/adapter/rest/routes_setup.go`:
```go
myService := container.BuildMyDomainService()
controller.NewMyDomainController(myService).AddRoutes(e, appName)
```

No code generation or CLI commands needed - the builder pattern provides type-safe, explicit dependency injection.

### OpenAPI Docs and Swagger
Steps to generate OpenAPI Docs and use via Swagger UI:
1. Install `swag` command into your local machine:
    ```shell
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
2. Annotate each controller with [Declarative Comment Format](https://github.com/swaggo/swag#declarative-comments-format)
3. Update and synchronize the `cmd/docs` module with your update:
    ```shell
    make swagger
    ```
4. Restart the service and access Swagger UI at `http://localhost:8888/:app-name/swagger/index.html`


### Migration
generate migration, up, down and status
1. Run the command to download and install the migration library.
```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
``` 
2. Run the command for migration creation, this will create the migration file version
```shell
make migration-create name={name_migration} type={go|sql}
```
3. Run the command for migration up, this command will execute the generated migration files
```shell
make migration-up
```
4. Run the command for migration down, this command will execute rollback migration based on version of your migration file
```shell
make migration-down
```
5. Run the command for migration status, we can check status the migration using this command
```shell
make migration-status
```
6. All `migration-*` targets default to the `mysql` driver. Override with `DRIVER=postgres` to target PostgreSQL instead:
```shell
make migration-up DRIVER=postgres
```
7. Migrations can also be run/rolled back from GitHub Actions via the **DB Migration** workflow (`.github/workflows/migration.yml`, `workflow_dispatch`). Pick the action (`up`, `up-to`, `down`, `down-to`, `redo`, `status`), the `db_driver`, and the target GitHub Environment holding the `DB_HOST` / `DB_PORT` / `DB_USERNAME` / `DB_PASSWORD` / `DB_NAME` secrets.

> Note: the bundled migration files under `migrations/` are written with MySQL-specific DDL (`AUTO_INCREMENT`, `bigint unsigned`, `tinyint(1)`, `datetime`). When targeting PostgreSQL for the first time, add Postgres-compatible migration files (or port the existing ones) before running `up`.

### Switching database driver

The ent client (`configs/database/connection_sqlent.go`, `connection_ent.go`) and the migration runner build their DSN from `configs/database/dsn.go`, which reads:
- `db.configs.driver` (`.env`) - `mysql` (default) or `postgres`
- `db.configs.username` / `db.configs.password` / `db.configs.host` / `db.configs.port` / `db.configs.database` (`secret.env`)
- `db.configs.sslmode` (`.env`, PostgreSQL only) - defaults to `disable`

To switch the running application to PostgreSQL, set `db.configs.driver = "postgres"` in `.env` (and adjust credentials in `secret.env`). No code changes are required.

### Seeders
Plain, idempotent `.sql` files under `seeders/` (numbered, e.g. `0001_seed_default_roles.sql`) that can be re-run safely - each statement guards itself with `WHERE NOT EXISTS (...)`. Run them via the **DB Seeder** GitHub Actions workflow (`.github/workflows/seeder.yml`, `workflow_dispatch`):
- `seeder_file`: pick a specific file under `seeders/`, or `all` to run every file in order
- `db_driver`: `mysql` or `postgres`
- `target_environment`: the GitHub Environment holding the `DB_HOST` / `DB_PORT` / `DB_USERNAME` / `DB_PASSWORD` / `DB_NAME` secrets

When adding a new seeder file, also add it to the `seeder_file` choice list in `.github/workflows/seeder.yml`.

### Run config
- Default general config is located in `root-project-path/.env`
- For credential the default config is located in `root-project-path/secret.env`
  - We can use custom path on running the executable application `./main -credentials-path="ABC" -credentials-name="XYZ"`


## Build project
using Makefile to run the project. `make all` will execute schema ent generation, mockery for mock test, swagger docs, testing project and build then run project
```shell
make all
```
## Example code
- health check example using MVC pattern - [health](internal%2Fapplications%2Fhealth)
- System parameter for feature flag as CRUD example using MVC & repository pattern - [system_parameter](internal%2Fapplications%2Fsystem_parameter)
- Transaction CRUD example using ent go - [service](internal%2Fapplications%2Fuser%2Fservice)
- Publish and subscribe using rabbitmq - [add link]

## References project
- Golang 1.19
  - https://go.dev/doc/effective_go
  - https://github.com/golovers/effective-go
  - https://go.dev/play/
  - Video references from PZN bahasa indo : [PZN-golang-playlist](https://www.youtube.com/watch?v=JOXbresHhIk&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ)
  - Tutorial bahasa indonesia [dasar golang noval agung](https://dasarpemrogramangolang.novalagung.com/1-berkenalan-dengan-golang.html)
- echo framework v4
  - https://echo.labstack.com/, Not the fastest, but on par with GIN with better documentation [benchmark discussion](https://github.com/labstack/echo/discussions/2143)
  - There are also a lot of tutorial on the net using bahasa indo, and easy for beginner ! [noval agung echo framework rest api](https://dasarpemrogramangolang.novalagung.com/C-echo-routing.html) 
- Viper 
  - https://github.com/spf13/viper, commonly used and powerfull configuration libs
- Entgo
  - https://entgo.io/
    - New kids on the block, developed by facebook team. Not the fastest, but better than gorm and have generated query builder! 
    - See benchmark : [ent benchmark](https://github.com/efectn/go-orm-benchmarks/blob/master/results.md)
- Dependency Injection - Builder Pattern
  - This project uses a fluent builder pattern for dependency injection
  - Service builders are registered in `internal/builder/services.go`
  - All dependencies are explicitly visible and type-safe
- testify : [assert test](https://github.com/stretchr/testify) 
- mockery : [mock test](https://vektra.github.io/mockery/latest/)
- Cache : {to add explanation later}
- Logging : using echo log 
- Message brooker : {to add explanation later}
- API documentation : [swagger using swaggo](https://github.com/swaggo/swag) 
- Containerization : {to add explanation later}

## List of project features
- [x] create clean code structure
- [x] create interface with example domains system param
- [x] create manual DI on hello worlds example domains
- [x] implement repository + database connection using ent in system param example domains
- [x] implement optimistic locking https://github.com/ent/ent/blob/master/examples/version/README.md
- [x] implement global error handling
- [x] implement builder pattern dependency injection
- [x] implement redis cache
- [x] implement migration files, instead of using ent / atlas we decide to use https://pkg.go.dev/github.com/golang-migrate/migrate/v4 for easiness and simplicity
- [x] Example - implement test for CRUD example
  - [x] positive test case
  - great talk by imre ! [Writing Better Test in Go | Go Srilanka Meetup June 25th, 2021](https://www.youtube.com/watch?v=xTQI_4EKB8Y)
  - using testify for the sake concise code, no needs write many if for assertion
  - using mockery
- [x] Example - implement outbound http calls - support config timeout, retry, circuit breaker
- [x] Example - implement transaction examples - https://entgo.io/docs/transactions
- [x] Example - implement migration files - https://entgo.io/docs/data-migrations, https://atlasgo.io/
- [x] implement logger
- [ ] integrate live reload
- [ ] implement message broker integration - rabbitmq
- [x] integrate swagger or API docs - https://github.com/swaggo/echo-swagger
- [ ] dockerize project / Colima 