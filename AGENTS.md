# Repository Guidelines

## Project Structure & Module Organization
- `cmd/main.go`: service entry. `cmd/docs`: generated Swagger/OpenAPI.
- `internal/adapter/rest`: centralized route wiring.
- `internal/applications/<domain>`: MVC modules: `controller`, `service`, `repository`, `dto`; DI via builder pattern in `internal/builder/services.go`.
- `configs/`: env, logger, validator, DB (Ent), Redis cache, RabbitMQ, Swagger. Runtime config from `.env` and `secret.env`.
- `ent/`: generated ORM; edit only `ent/schema/*.go`, then generate.
- `migrations/`: Goose migrations and CLI (`migrations/cmd/main.go`).
- `middleware/`, `mocks/` (Mockery), `test/` (shared helpers).

## Architecture Overview
- Request flow: Echo routes → controller → service → repository (Ent). Caching via `internal/component/cache`. RabbitMQ producer/consumer registered in `internal/component/rabbitmq/*` and toggled by config.

## Build, Test, and Development Commands
- Build/Run: `make build` (outputs `./main`), `make run` (runs binary).
- Full cycle: `make all` (schema, mocks, swagger, tests, build, run).
- Tests/Coverage: `make test` (generates `coverage.html`).
- Ent: `make schema` or `make schema-advance` (advanced features).
- Mocks: `make mocks` (generates repository mocks). DI is handled by builder pattern.
- Swagger: `make swagger` then visit `http://localhost:<port>/<appName>/swagger/index.html`.
- Migrations: `make migration-create name=<name> type=<go|sql>`, `make migration-up|down|status`.

## Coding Style & Naming Conventions
- Go 1.22; format with `go fmt`; idiomatic Go. Tabs for indentation.
- Packages lower-case; files `snake_case.go`. Exported: PascalCase; unexported: camelCase.
- Domain naming: `<domain>_controller.go`, `<domain>_routes.go`, `<domain>_service*.go`, `<domain>_repository*.go`.
- Do not edit generated files (`ent/*` except `ent/schema`). All DI is explicit in `internal/builder/`.

## Dependency Injection - Builder Pattern
When adding a new service/domain:
1. Implement service interface in `internal/applications/<domain>/service/`
2. Implement repository interface in `internal/applications/<domain>/repository/`
3. Add builder method in `internal/builder/services.go`:
   ```go
   func (c *Container) Build<Domain>Service() <domain>.Service {
       repo := repository.New<Domain>Repository(c.db)
       return service.New<Domain>Service(repo, ...)
   }
   ```
4. Register routes in `internal/adapter/rest/routes_setup.go`:
   ```go
   svc := container.Build<Domain>Service()
   controller.New<Domain>Controller(svc).AddRoutes(e, appName)
   ```
No code generation or CLI commands needed—the builder provides explicit, type-safe DI.

## Testing Guidelines
- Use `testing`, `testify`, and Mockery. Tests live next to code as `<file>_test.go` with `TestXxx`.
- Shared helpers in `test/` (Echo/Resty/Ent). Run `make test`; open `coverage.html`.
- Regenerate mocks when interfaces change: `make mocks`.

## Commit & Pull Request Guidelines
- Commits: concise, imperative, and scoped (e.g., `system_parameter: enforce optimistic lock`).
- PRs: follow `pull_request_template.md`—clear description, linked issues, tests/coverage evidence, risk level, and checklist; include Swagger updates if handlers change.

## Security & Configuration Tips
- Do not commit secrets. Use `.env` (general) and `secret.env` (credentials). Flags override: `./main -credentials-path=PATH -credentials-name=NAME`.
- RabbitMQ optional via `rabbitmq.configs.enable`; ensure Redis/DB are reachable before enabling features that depend on them.

## Local Config Quickstart
- `.env` (app/runtime):
  
  application.name=/micro-go-template
  application.port=8888
  application.cors.allowedHost=["http://localhost:3000","http://localhost:8888"]
  rabbitmq.configs.enable=false
  
- `secret.env` (credentials):
  
  db.configs.username=root
  db.configs.password=password
  db.configs.host=127.0.0.1
  db.configs.port=3306
  db.configs.database=echo_sample
  rabbitmq.configs.username=guest
  rabbitmq.configs.password=guest
  rabbitmq.configs.host=127.0.0.1
  rabbitmq.configs.port=5672

Run with overrides: `./main -credentials-path=. -credentials-name=secret`.

## CI
- GitHub Actions: `.github/workflows/go.yml` runs on `master`, `release`, `develop` for push/PR.
- Steps: checkout, setup Go via `go.mod`, build `go build ./...`, test `go test -v ./...`.
- Local coverage: `make test` generates `coverage.html` (not uploaded in CI by default).

## Run Locally
- Prereqs: set `.env` and `secret.env`. Ensure `application.name` starts with `/` (e.g., `/micro-go-template`). If you change `application.port`, update `swagger.host` to `localhost:<port>` for correct Swagger "Try it out" host.
- Start: `make build && ./main` or full cycle `make all`.
- Access: Swagger at `http://localhost:<port>/<appName>/swagger/index.html`; health check at `http://localhost:<port>/<appName>/health`.
