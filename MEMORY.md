# MEMORY.md

Durable, non-obvious facts about this repository that the code alone won't tell you fast
enough. If an entry stops being true, delete it in the same change that makes it untrue —
don't leave it to rot.

---

## Admin role ID is a hardcoded literal tied to seed-migration row order

**What:** `middleware.AdminRoleID = uint64(1)` (`middleware/auth_middleware.go`) and
`defaultAutoProvisionRoleID = uint64(2)` (`internal/applications/auth/service/auth_service_impl.go`)
are bare integer literals, not looked up by name at runtime.

**Why:** They match the exact row IDs `migrations/20260630120005_seed_roles_and_admin_user.go`
inserts explicitly (`INSERT INTO roles (id, ...) VALUES (1, ..., 'Admin', ...)` /
`VALUES (2, ..., 'User', ...)`), which only works because that migration assumes the `roles`
table is empty when it runs.

**How to apply:** If you ever change the seed migration's role IDs, add a new role before
"Admin"/"User", or re-run seeding against a non-empty table, update both constants in the same
change. `RequireAdmin()` silently and permanently locks out real admins (or, worse, grants
admin to the wrong role) if these three numbers ever disagree — there's no runtime check that
they still match.

---

## Ent `Update` methods must keep the `.Versions()` filter and the affected-rows check

**What:** Every repository `Update` (see `system_parameter_repository_impl.go`) filters on
both `.ID(...)` and `.Versions(...)`, then treats `affected < 1` from the update as an error
("no records were updated"), not as a no-op success.

**Why:** `VersionMixin` (`ent/schema/base_mixin.go`) is this repo's optimistic-concurrency
mechanism. Dropping the `.Versions()` predicate — even though it still "compiles" and the
update still "succeeds" — turns every write into a blind last-writer-wins, silently discarding
concurrent changes instead of surfacing a conflict.

**How to apply:** Any new `Update` method on a `VersionMixin` entity must copy this exact
shape: filter on current version, check affected rows, return an explicit error on zero rows.

---

## The seeded admin password is a real, working placeholder — not a dummy value

**What:** `migrations/20260630120005_seed_roles_and_admin_user.go` bcrypt-hashes and inserts
the literal password `ChangeMe123!` for `admin@example.com` on every fresh install. The
migration's own comment says to rotate it immediately in any non-local environment, but
nothing in the code enforces that.

**Why:** It's how a fresh `make migration-up` gives you a working admin login out of the box,
same idea as the `CHANGEME_*` Auth0 placeholders in `.env`/`secret.env`.

**How to apply:** Never remove the rotate-immediately comment, never "fix" it by making the
password more complex (that just changes what needs rotating), and flag it if you ever see
this migration referenced in a production/staging deployment plan without a password-rotation
step next to it.

---

## `application.cors.allowedHost` only ever registers as a single CORS origin

**What:** `middleware.CorsConfig()` reads `application.cors.allowedHost` with
`viper.GetString(...)` and passes it straight into `AllowOrigins: []string{corsAllowedHost}` —
one element, whatever the raw string is. The in-code default
(`configs/initialize_general_env.go`) sets it as a real `[]string` of three hosts, but once
it's read back out through `GetString` for CORS, it collapses to a single literal string.

**Why:** Unclear whether this is intentional (single-origin deployments only) or a latent gap
between the multi-host default/documentation and what `CorsConfig()` actually does with it —
see Open Questions below.

**How to apply:** Don't assume setting a comma-separated or JSON-array-looking value in
`application.cors.allowedHost` enables multiple allowed origins — as written, it doesn't. If
you need multi-origin CORS, `CorsConfig()` needs an actual split/parse step, not just a
`GetString`.

---

## Migrations are MySQL-specific DDL even though the driver is switchable

**What:** `db.configs.driver` (`.env`) toggles `mysql`/`postgres` cleanly in
`configs/database/dsn.go`, but every committed file under `migrations/` uses MySQL-only DDL
(`AUTO_INCREMENT`, `bigint unsigned`, `tinyint(1)`, `datetime`).

**Why:** The driver switch was added for the *application's* DB connection, not retrofitted
onto the migration history.

**How to apply:** Don't flip `db.configs.driver` to `postgres` against a fresh database and
expect `make migration-up` to work — Postgres-compatible migration files need to be written
(or the existing ones ported) first. `make test` won't catch this either (see below).

---

## `make test` never exercises MySQL or PostgreSQL — only in-memory SQLite

**What:** `test.DbConnection` (`test/ent_test_helper.go`) opens every repository test against
`enttest`/`mattn/go-sqlite3` in-memory. There is no CI job or local target that runs the test
suite against a real MySQL or Postgres instance.

**Why:** Keeps tests fast and hermetic, which is the right tradeoff for unit/integration
coverage — but it means driver-specific SQL (including the MySQL-only migration DDL above)
gets zero automated coverage.

**How to apply:** A green `make test` does not confirm a change works against the real MySQL
target (or would work against Postgres). Anything touching raw SQL, driver-specific Ent
features, or the migration files needs a manual check against an actual `docker-compose up`
database, not just the test suite.

---

## The example RabbitMQ producer route is always registered, even when RabbitMQ is disabled

**What:** `internal/adapter/rest/routes_setup.go` unconditionally calls
`container.BuildProducer()` and registers the example RabbitMQ controller's routes.
`rabbitmq.configs.enable` (checked in `cmd/main.go`) only gates whether the RabbitMQ
*connection* is established and consumers/producers are registered against it — it does not
gate route registration.

**Why:** `Container.rabbit` stays `nil` when the flag is off; `BuildProducer()` still
constructs a `WrappedChannelServiceImpl` wrapping that `nil` connection, which only fails when
`PublishMessage` actually calls `wc.connection.GetChannel()`.

**How to apply:** With RabbitMQ disabled, hitting the example producer endpoint panics
(recovered into a 500 by `middleware.Recover()`) instead of 404ing or failing at startup. If
you add a new RabbitMQ-backed route, either gate its registration on
`viper.GetBool("rabbitmq.configs.enable")` in `routes_setup.go`, or accept and document the
same lazy-failure behavior — don't assume the route silently disappears when the flag is off.

---

## Open questions

Found during exploration; couldn't confirm intent from the code or history, so recorded here
instead of guessed at.

- **`response.Error()`** (`internal/helper/response/response_json_builder.go`) is unused
  anywhere in the codebase and carries a `//goland:noinspection GoUnusedExportedFunction`
  annotation — someone deliberately suppressed the unused-export warning rather than deleting
  it. Is this kept as public API surface for future direct-error-response call sites, or is it
  safe to delete? Don't delete it as a drive-by cleanup without checking first.
- **`.mockery.yaml`** at the repo root is a 0-byte file. `make mocks` passes all its config via
  CLI flags (`mockery --all --dir internal --output mocks --packageprefix mock_ --keeptree`)
  and never reads the yaml file. Is the empty file a leftover from an earlier setup, or
  reserved for a planned migration to yaml-based config? If mock generation ever needs to
  change, the Makefile target is the one that actually has effect today.
- **`golangci-lint` version differs between local and CI**: `.pre-commit-config.yaml` pins
  `v1.55.2`; `.github/workflows/ci.yml` uses `golangci-lint-action@v6` with `version: latest`.
  They can silently drift apart (a rule that fails locally may pass in CI, or vice versa).
  Which one is meant to be authoritative wasn't clear from the repo — worth pinning both to
  the same version the next time either is touched, but not changing unprompted.
