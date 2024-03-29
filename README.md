# Micro Go Template
Boilerplate template for backend project using go language for optimize and efficient resources.

## Contents
- [Getting started](#getting-started)
  - [Install and generate ORM](#install-and-generate-orm)
  - [Install and generate mock using Mockery](#install-and-generate-mock-using-mockery)
  - [Generate Dependency Injection](#generate-dependency-injection)
  - [OpenAPI Docs and Swagger](#openapi-docs-and-swagger)
  - [Migration](#migration)
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

### Generate Dependency Injection
1. Install google wire CLI
```shell
go install github.com/google/wire/cmd/wire@latest
```
2. Add wire on your $PATH, so we can use wire CLI on every project
3. Create {domains}_injector.go in your feature directory
4. Then Run wire using makefile
```shell
>> make wire
>> Enter directory: 
>> A_templates_directory
```

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
- Google wire - code gen for dependency injection
  - https://github.com/google/wire, DI code generator
  - Good tutorial for getting started with example [tutorial google DI with google wire](https://clavinjune.dev/en/blogs/golang-dependency-injection-using-wire/)
  - [Video references from PZN - golang DI with google wire](https://www.youtube.com/watch?v=dZ8Ir4Gc8D0&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=14)
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
- [x] implement DI google wire
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