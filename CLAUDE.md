# CLAUDE.md

Behavioral contract for AI-assisted work on this repository (`myapp` — a Go/Echo/Ent
backend boilerplate, module `myapp`, currently instantiated as a small admin/system-parameter
service with Auth0-backed login, Redis caching, and optional RabbitMQ).

---

## Working agreement

- Act as a senior Go engineer: thoughtful, skeptical, thorough. Not eager to please — push
  back on requests that would introduce bugs, break layering, or add unneeded complexity,
  and say so plainly.
- Do not assume. If a requirement is ambiguous, or there are several reasonable ways to do
  something, ask, incorporate the answer, then continue. Silently picking one and moving on
  is not acceptable on anything non-trivial.
- Plan first for anything non-trivial (new domain, schema change, cross-cutting change,
  anything touching auth or money-adjacent logic here would be system parameters/config).
  Small, direct edits (typo fixes, adding a field, a one-line bug fix) may skip planning.
- **No git write operations.** Never run `commit`, `push`, `reset`, `rebase`, `tag`, or
  `branch -D`/delete. Print the exact commands for a human to run instead. This is enforced
  by a PreToolUse hook (see `.claude/settings.json`) and is not negotiable regardless of how
  the request is phrased.
- Keep reasoning concise. Long internal deliberation dumped into chat is noise; state
  conclusions and the one or two facts that justify them.

## General code rules

- Preserve existing architecture and conventions. Read the surrounding domain (an existing
  `internal/applications/<domain>` folder) before writing a new one — copy its shape, don't
  invent a new one.
- Follow existing patterns before introducing new ones. If two patterns already coexist in
  the codebase, ask which is canonical rather than adding a third.
- Extend rather than abstract. Do not introduce a new interface, generic helper, or
  base-struct to "avoid repetition" until at least three real call sites need it in this repo.
- Minimal diffs. No unrelated refactoring, renaming, or reformatting bundled into a feature
  or bugfix change. If you spot something else worth fixing, mention it, don't fix it inline.
- Explicitness over magic. This codebase already favors explicit, hand-wired code (the
  builder-pattern DI container is a deliberate example) over reflection-heavy or
  code-generated wiring outside of Ent/mockery/swag. Match that.
- Comment the *why*, never the *what*. Most existing code has no comments; the handful that
  exist explain a non-obvious constraint (e.g. why `AdminRoleID` is hardcoded to `1`). Match
  that density — don't add narration comments.

## Go rules

Target: **Go 1.22** (`go.mod`). Note `go.mod` pins `labstack/echo v4.12.0` with an inline
comment that later Echo versions require Go 1.25 — do not bump Echo without also confirming
the Go toolchain bump is wanted.

- Wrap errors with `%w` and check with `errors.Is`/`errors.As`. The codebase's own
  `exceptions.BusinessLogicError` already implements `Unwrap()`/`errors.As`-compatible
  digging (`exceptions/business_logic_error.go`) — use `exceptions.NewBusinessLogicError(code, err)`
  at the service layer rather than inventing a new error type.
- `context.Context` is the first parameter on every service/repository/outbound-client
  method that does I/O. Never store a `context.Context` in a struct field — every example in
  this repo threads it through as a parameter.
- No `panic` outside `main`, with one existing, deliberate exception:
  `internal/helper/response.Success`/`Created` panic if `data == nil`, as a programmer-error
  assertion (a controller forgot to pass a response body). Don't add new panics elsewhere;
  don't "fix" that existing one as a drive-by change — it's intentional.
- No new package-level mutable state and no side-effecting `init()`. One existing exception:
  `internal/vars/cache_duration.go` initializes a package-level `ttl` var from viper in
  `init()`. This is pre-existing, not a pattern to copy into new code, and not something to
  refactor unless you're asked to touch that file specifically.
- Interfaces in this codebase are declared **in the producer's package, next to the
  implementation** — e.g. `service.SystemParameterService` lives in
  `.../system_parameter/service/system_parameter_service.go` beside
  `SystemParameterServiceImpl`, and `db.SystemParameterRepository` lives beside
  `SystemParameterRepositoryImpl`. This is the opposite of the textbook "consumer defines the
  interface" idiom — follow the repo's actual convention, not the textbook one, when adding a
  new service/repository.
- Every goroutine has a stated exit path. The two long-running goroutines in this codebase
  (the HTTP server in `cmd/main.go`, and RabbitMQ consumers under
  `internal/component/rabbitmq`) are stopped via the graceful-shutdown signal handler and
  connection `Close()` in `cmd/main.go` — new goroutines need the same kind of explicit,
  traceable shutdown, not a bare `go func(){ for {...} }()`.
- Avoid `any`/`interface{}` and unnecessary generics. The one broad use of `interface{}` is
  the cache layer (`internal/component/cache`), which genuinely needs to serialize arbitrary
  types — don't take that as license to use `any` elsewhere.

## Layering

Request flow is strictly one direction:

```
Echo route (internal/adapter/rest, internal/adapter/web)
   -> Controller  (internal/applications/<domain>/controller)
   -> Service     (internal/applications/<domain>/service)
   -> Repository  (internal/applications/<domain>/repository/{db,outbound,inbound,auth0})
   -> Ent client (ent/) or external system (Redis, RabbitMQ, Auth0, HTTP outbound)
```

Rules that follow from the actual code (not an idealized version of it):

- **Imports only flow downward.** Controllers import services; services import repositories
  and `ent`; repositories import `ent` and nothing above them. A repository must never import
  a controller or service package, and a service must never import a controller package.
- **DTOs live per-domain** in `internal/applications/<domain>/dto`. Request DTOs
  (`*CreateRequest`, `*UpdateRequest`) carry `validate` tags for `go-playground/validator`.
  Response DTOs (`*Response`) are the only shape that ever leaves the service layer toward
  JSON.
- **The DTO <-> entity crossing is asymmetric by convention, follow it as-is:**
  - Request DTO -> `ent.<Entity>` is built by hand in the service method (field-by-field
    struct literal), because the request shape is intentionally narrower than the entity.
  - `ent.<Entity>` -> Response DTO is done through the named converter
    `internal/helper.Mapper` (wraps `github.com/dranikpg/dto-mapper`), called from the
    **controller**, not the service. Services return `*ent.<Entity>` (or slices of it)
    directly — they do not return response DTOs.
  - Do not replace either direction with a direct struct assignment/type-cast across the
    boundary; do not move the `Mapper` call into the service layer without discussing it
    first — that would be a cross-cutting convention change, not a local fix.
- **Repositories operate on `*ent.<Entity>`**, never on DTOs. A repository interface is
  declared in `.../repository/<subpkg>/<domain>_repository.go` beside its implementation
  (see Go rules above on interface placement).
- **Cross-domain repository subpackages**: `repository/db` for direct Ent-backed persistence,
  `repository/outbound` for calling external HTTP APIs (see `quotes`), `repository/inbound`
  for RabbitMQ consumers (see `example/rabbitmq`), `repository/auth0` for the Auth0 identity
  provider client (see `auth`). Pick the one matching what the new repository actually talks
  to — don't put an HTTP client inside `repository/db`.
- **Wiring is manual, not generated.** New services are wired in
  `internal/builder/services.go` as `Build<Domain>Service()` methods on `*Container`, and
  routes are registered in `internal/adapter/rest/routes_setup.go` by calling
  `container.Build<Domain>Service()` then `<domain>Controller.AddRoutes(e, appName, mw...)`.
  There is no DI code generation (no wire, no fx) — follow the existing builder methods as
  the template.
- **`make template name=<domain>`** scaffolds a new domain from
  `internal/applications/A_templates_directory` (controller/service/repository/dto stubs with
  placeholders substituted). Use it for a brand-new domain instead of hand-copying an existing
  one; then wire it into `internal/builder/services.go` and `routes_setup.go` per the
  Makefile's own printed next-steps.

## Database

- ORM is **Ent** (`entgo.io/ent`). Generated code lives under `ent/` and is regenerated from
  `ent/schema/*.go` via `make schema` (standard) or `make schema-advance` (adds
  `sql/modifier`/`sql/execquery` features). Only `ent/schema/*.go` and `ent/generate.go` are
  hand-edited; everything else under `ent/` is regenerated output (also enforced by hook, see
  `.claude/settings.json`).
- Every entity mixes in `BaseFieldMixin` (`created_by`/`created_at`/`updated_by`/`updated_at`/
  `deleted_by`/`deleted_at`) and `VersionMixin` (`versions int64`, an optimistic-concurrency
  token) — see `ent/schema/base_mixin.go`. New entities should mix these in the same way
  unless there's a specific reason not to.
- **Optimistic locking is load-bearing.** Repository `Update` methods must filter on both the
  entity's ID and its current `.Versions()` value (see
  `system_parameter_repository_impl.go`'s `Where(systemparameter.ID(...), systemparameter.Versions(...))`),
  and must treat `affected < 1` as an error ("no records were updated") rather than silently
  succeeding — that's how a stale-version conflict is detected. Don't drop the version filter
  to "simplify" an update.
- Both hard delete (`Delete`, `client.<Entity>.DeleteOneID(...).Exec(ctx)`) and soft delete
  (`SoftDelete`, sets `deleted_by`/`deleted_at`) exist side by side on the same repository.
  Which one a domain's routes actually expose varies — check the controller's routes file,
  don't assume soft delete is always the "safe default" that's wired up.
- Driver is switchable: `db.configs.driver` in `.env` is `mysql` (default) or `postgres`, read
  by `configs/database/dsn.go`. **The committed migration files under `migrations/` are
  MySQL-specific DDL** (`AUTO_INCREMENT`, `bigint unsigned`, etc.) — targeting Postgres for
  the first time needs new/ported migration files, not just a config flip.
- Migrations use **goose** (`pressly/goose/v3`) via the custom CLI in `migrations/cmd/main.go`
  (built as `./migration` by `make migration-build`/`make migration-up`/etc.). Create new
  migrations with `make migration-create name=<name> type=<go|sql>` — never hand-write a new
  timestamped file. Existing, already-numbered migration files under `migrations/` are applied
  history and must not be edited (hook-enforced; see `.claude/settings.json`) — write a new
  migration to change behavior instead.
- Seeders are separate from migrations: plain, idempotent `.sql` files under `seeders/`
  (`WHERE NOT EXISTS (...)` guards), run manually via the **DB Seeder** GitHub Actions
  workflow, not part of the normal migration-up path.
- Tests never touch a real database: `test.DbConnection(t)` opens an in-memory SQLite client
  via `ent/enttest` (`mattn/go-sqlite3`). Repository tests run against real Ent code against
  that in-memory DB, not mocks.

## HTTP

- Framework is **Echo v4**. Every domain controller has a `<domain>_routes.go` implementing
  `AddRoutes(e *echo.Echo, appName string, mw ...echo.MiddlewareFunc)`, and all routes are
  registered centrally in `internal/adapter/rest/routes_setup.go` — don't call `e.GET/POST/...`
  from anywhere else.
- Global middleware (CORS, 5-minute timeout, gzip, logger, recover, request-id, secure
  headers, in-memory per-IP rate limit) is set up once in `middleware.SetupMiddlewares`
  (`middleware/echo_middlewares.go`). Auth middleware (`RequireAuth`, `RequireAuthWeb`,
  `RequireAdmin`) is applied selectively per route group in `routes_setup.go`, not globally —
  check an existing domain's route file (e.g. `system_parameter`, which is readable by both
  roles but Admin-only for writes) before wiring a new protected route.
- Every controller method carries `swaggo` declarative comments (`@summary`, `@param`,
  `@success`, `@failure`, `@router`, etc.). Add/update them on any handler change and run
  `make swagger` to regenerate `cmd/docs/` — don't hand-edit the generated docs files.
- Responses go through `internal/helper/response` (`response.Success`, `response.Created`,
  `response.Base`) — a fixed envelope `{code, message, data, error, serverTime}`. Don't build
  ad hoc `ctx.JSON(...)` responses in a controller; use the existing helpers so the envelope
  stays consistent.
- Request binding + validation is combined via `helper.BindAndValidate(ctx, dto)`
  (`ctx.Bind` + `ctx.Validate`, the latter backed by `go-playground/validator` with custom
  validators for password/email in `configs/validator`). Business-logic errors from a service
  become HTTP responses through the single global handler
  (`configs/validator/response_error_handler.go` -> `exceptions.BusinessLogicReason`) — a
  service should return an `exceptions.BusinessLogicError` with the right error code from
  `exceptions/err_enum.go`, not format its own HTTP response.

## Observability

- Logging is `github.com/labstack/gommon/log`, wired into Echo. There is no tracing or
  metrics stack in this repo (no OpenTelemetry, no Prometheus) — don't assume one exists or
  reference spans/metrics that aren't there.
- `middleware.RequestID()` is enabled globally; propagate/log it if you add new logging, don't
  invent a separate correlation-ID scheme.

## Security

- Auth is **Auth0-backed** (password grant + Google OAuth), issuing an app-level **JWT
  (HS256)** stored in an httpOnly session cookie (`middleware.SessionCookieName`), verified
  per-request by `RequireAuth` (`middleware/auth_middleware.go`). `RequireAdmin` gates by
  `AdminRoleID = 1`, a constant tied directly to the row inserted by
  `migrations/20260630120005_seed_roles_and_admin_user.go` — if you ever touch that seed
  migration's role IDs, `AdminRoleID` (and `defaultAutoProvisionRoleID = 2` in
  `auth_service_impl.go`) must change with it.
- Passwords are bcrypt-hashed (`golang.org/x/crypto/bcrypt`, default cost). Auth0-provisioned
  users get a random local password that is never used to authenticate directly — Auth0 is
  the credential source of truth.
- Rate limiting is in-memory, per-IP, via Echo's built-in limiter (`middleware.GlobalRateLimiter`,
  `middleware.LoginRateLimiter`) — there's no Redis-backed distributed limiter, so it resets
  per-instance.
- Secrets (`jwt.secret`, `auth0.clientSecret`, DB/Redis/RabbitMQ credentials) live in
  `secret.env` locally and in GitHub Environment secrets for CI/CD workflows — never hardcode
  a credential in application code. `secret.env` and `.env` are committed with placeholder
  values today (`CHANGEME_...`, a documented bootstrap admin password) — do not replace those
  placeholders with anything that looks like a real credential.

## Testing

- `testify` (`assert`/`require`) + `mockery`-generated mocks. Mocks live under `mocks/`,
  mirroring the `internal/` tree with a `mock_` package prefix (`make mocks` — never hand-edit
  a file under `mocks/`, hook-enforced).
- Repository tests exercise real Ent code against the in-memory SQLite helper
  (`test.DbConnection`); service tests use the generated mocks for the repository/cache
  interfaces they depend on. Match whichever pattern the layer you're testing already uses.
- Tests live beside the code as `<file>_test.go`, named `Test<Type>_<Method>`.
- **Going forward, order test cases negative-first:** failing/error-path assertions in their
  own block at the top of a test function (or as their own `Test..._<Scenario>` functions),
  then the happy-path assertions after. The existing suite in this repo is happy-path-only —
  don't mass-rewrite existing tests just to satisfy this ordering, but new tests and any test
  you're already touching should follow it.
- Run only the affected package's tests while iterating (`go test ./internal/applications/<domain>/...`);
  reserve the full `make test` run for before you consider the change done, and ask before
  running it if it's expensive in context (it runs and re-covers the whole module).

## Commands

| Purpose | Command |
|---|---|
| Build (runs tests first) | `make build` |
| Run the built binary | `make run` |
| Full cycle (schema, mocks, swagger, test, build, run) | `make all` |
| Unit/integration tests + coverage | `make test` |
| Generate Ent models (standard) | `make schema` |
| Generate Ent models (sql/modifier, sql/execquery) | `make schema-advance` |
| Generate mockery mocks | `make mocks` |
| Generate/refresh Swagger docs | `make swagger` |
| Scaffold a new domain | `make template name=<domain>` |
| Build migration CLI | `make migration-build` |
| Create a migration | `make migration-create name=<name> type=<go\|sql>` |
| Run/rollback/status migrations | `make migration-up` / `make migration-down` / `make migration-status` (add `DRIVER=postgres` to target Postgres) |
| Lint | `make lint` (`golangci-lint run ./...`) |
| Security scan | `make security-scan` (`gosec ./...`) |
| Coverage report | `make coverage-report` |
| Dependency check | `make deps-check` |
| All CI checks locally | `make ci-all` |
| Docker build / compose up / compose down | `make docker-build` / `make docker-compose-up` / `make docker-compose-down` |

Drive these Makefile targets rather than reimplementing what they do — if there's a generator
(`make schema`, `make mocks`, `make swagger`, `make migration-create`, `make template`),
hand-writing the files it produces is forbidden. Ask before running `make lint`,
`make security-scan`, or the full `make test`/`make ci-all` suite, since they're
whole-repo operations. Never run `make clean` without explicit confirmation — it deletes
generated Ent code, all mocks, and coverage reports. Prefer `rg` over `grep` for searching.

## Cost and context rules

- Read only the files relevant to the change. Don't scan the whole repo "just in case" —
  the domain-per-folder layout under `internal/applications/` means the files for a given
  change are almost always confined to one domain folder plus, at most, `internal/builder/`
  and `internal/adapter/rest/routes_setup.go`.
- Run only the tests affected by a change while iterating; run the full suite once, at the
  end, as a final check.
- A smaller, targeted context produces more reliable output here than a broad one — this
  isn't just a cost optimization.

## Documentation rules

Update these when a change actually affects them — skip all of them for internal refactors
with no observable effect on architecture, API, schema, or workflow:

- **Swagger annotations + `make swagger`** — any change to a controller's request/response
  shape, route, or status codes.
- **`README.md`** — any change to setup steps, `make` targets, environment variables, or the
  driver-switching / migration / seeder workflow it documents.
- **`AGENTS.md`** — any change to the actual project structure, build commands, or
  conventions it documents (its "Repository Guidelines" section mirrors this file at a
  higher level — keep them consistent if you touch one).
- **`MEMORY.md`** — any new non-obvious, durable fact you discover or introduce (a load-bearing
  constant, an intentional-looking piece of dead code, a workaround). Delete an entry the
  moment it stops being true.
- **This file (`CLAUDE.md`)** — any change to layering, the DTO/entity conversion convention,
  or anything else asserted here as ground truth. A stale rule here is worse than no rule.

## Never touch

These are hook-enforced (`.claude/settings.json`), listed here so the denial makes sense
instead of looking like a bug:

- **`ent/**` except `ent/schema/**` and `ent/generate.go`** — generated by `entgo`; hand
  edits are silently overwritten by the next `make schema`/`make schema-advance` and rot the
  moment they drift from the schema.
- **`cmd/docs/**`** — generated by `swag`; edit the source controller's Swagger comments and
  run `make swagger` instead.
- **`mocks/**` and `internal/component/rabbitmq/mocks/**`** — generated by `mockery`; edit the
  interface and run `make mocks`.
- **`go.sum`** — derived from `go.mod`; hand-editing it breaks module checksum verification.
  Use `go get`/`go mod tidy`.
- **`.env`, `secret.env`, and any `.env.*`** — credential/config files. Real values belong in
  a local override or GitHub Environment secret, never in a code change.
- **Already-numbered files under `migrations/`** — applied migration history is immutable in
  goose; create a new migration instead of editing an old one.
- **`coverage.html`, `coverage.out`** — regenerated by `make test`; not meaningful to hand-edit.
