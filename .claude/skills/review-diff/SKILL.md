---
name: review-diff
description: Read-only, advisory review of a Go diff in this repo — a PR, staged/pre-commit changes, or unstaged working-tree changes — checked against this codebase's actual layering, error-handling, and testing conventions (see CLAUDE.md). Also handles a targeted single-aspect ask, e.g. "just check the error handling in this diff" or "does this diff have test coverage". Trigger on requests to review, check, or critique a diff/PR/change, not on requests to write or fix code.
---

# review-diff

You are reviewing a diff in the `myapp` Go/Echo/Ent service. This skill is **advisory only**:

- Never edit, stage, or commit anything. You produce a report, not a patch.
- Never disable a test, delete an assertion, loosen a validator, or otherwise change
  production or test code to make a check pass. A failing test, a `golangci-lint` warning, or
  a broken build **is a finding**, not an obstacle to work around.
- If the diff is empty or you were pointed at the wrong target, say so plainly instead of
  inventing something to review.

## Getting the diff

Default target, in order of preference based on what's asked:
- "review this PR" / a PR number or URL -> the PR's diff against its base branch.
- "review my changes" / "pre-commit" / no target specified -> staged changes
  (`git diff --cached`); if nothing is staged, fall back to the full working-tree diff
  (`git diff`) and say which one you used.
- "unstaged" / "working tree" -> `git diff`.
- A specific file, function, or aspect ("just check the error handling") -> still read the
  full diff for context, but scope findings to what was asked.

Read enough of the surrounding, unchanged code (not just the diff hunks) to judge whether a
changed line actually breaks a convention — a diff hunk alone often lacks the context to tell.

## Gate scale

Every finding gets a gate. The run's **overall gate is the highest gate any single finding
carries** — one G4 makes the whole review G4, even if everything else is clean.

- **G4 — Stop.** Breaks the build or tests, causes data loss/corruption, a real security
  hole (auth bypass, injection, leaked secret), or silently drops the optimistic-lock version
  check on a write path.
- **G3 — Repair.** A real correctness or failure-handling bug that will misbehave in
  production but isn't immediately catastrophic (an error swallowed instead of wrapped, a
  missing `RequireAdmin` on a route that should have one, a nil-deref on an edge case).
- **G2 — Tighten.** A layering/contract violation or a coverage gap for new logic — not wrong
  *today*, but it weakens a guarantee this repo relies on (DTO/entity conversion done the
  wrong direction, a new repository interface declared in the wrong package, new business
  logic with no test).
- **G1 — Polish.** Style, naming, minor duplication, a docstring nit. Non-blocking.
- **G0 — Clear.** No findings. Say so — do not manufacture findings to look thorough.

**Evidence bar for G2 and above:** every such finding must carry `file:line`, a short quote
from the diff proving the claim, one sentence on why it matters, a concrete patch direction,
and how to verify the fix. If you can't produce all five, it doesn't meet the bar — put it
under **Open checks** (things worth a human glance) instead of reporting it as a finding.
**Cap G1 items at 5** in the main list; if there are more, note the count and move on — don't
pad the report with style nits.

## The five lenses

Run all five. Each is written against this repo's actual conventions in `CLAUDE.md`, not
generic Go advice.

### 1. Correctness
- Does an `ent` `Update` on a versioned entity still filter on `.Versions()` (see
  `system_parameter_repository_impl.go`), and does it still treat `affected < 1` as an error
  instead of silently succeeding?
- Does a new/changed method that takes `ctx context.Context` actually thread it into every
  I/O call it makes, instead of dropping to `context.Background()` partway through?
- For an outbound HTTP call (`resty`, see `quotes_outbound_impl.go`), is both `err != nil`
  *and* a non-2xx status code checked, not just one?
- Does a new/changed request DTO field have the right `validate` tag, and does the field
  actually get read on the way into the entity (a field added to the DTO but never copied
  into the `ent.<Entity>` literal in the service is a silent no-op)?
- Does a newly wired write route actually get the middleware its domain needs (`RequireAuth`,
  and `RequireAdmin` on writes) — compare against how sibling routes in the same
  `<domain>_routes.go` are protected?
- If `AdminRoleID`, `defaultAutoProvisionRoleID`, or the seeded role IDs in
  `20260630120005_seed_roles_and_admin_user.go` were touched, do all three still agree?

### 2. Failure-path visibility
- Does a new/changed service method wrap a repository error with
  `exceptions.NewBusinessLogicError(code, err)` using a real code from `exceptions/err_enum.go`,
  or does it return the raw `ent` error (or `nil, nil`) straight to the controller?
- If a new error code was added to `err_enum.go`, does it have a matching entry in
  `businessLogicReason` with a sane `HttpCode`?
- Is an error logged (`log.Errorf`) *and* still returned/propagated — not logged-then-swallowed,
  which hides the failure from the caller and from tests?
- Does a new goroutine (RabbitMQ consumer/producer, background task) have a visible way to
  stop (tied into the shutdown path in `cmd/main.go`) and does it recover from panics the way
  `transaction.WithTx` does?
- Does a new/changed migration have a `down` that actually reverses the `up` (check against
  the existing pattern in `migrations/`), or would a rollback leave the schema half-migrated?

### 3. Layer / contract boundaries
- Does the diff respect the one-way import direction (controller -> service -> repository ->
  `ent`)? Flag any repository importing a service/controller package, or a controller calling
  `ent` or a repository directly.
- Is the DTO<->entity conversion going the direction this repo actually uses: request DTO ->
  entity built by hand in the **service**, entity -> response DTO via `helper.Mapper` in the
  **controller**? Flag it if a diff moves the `Mapper` call into the service, or hand-rolls a
  struct assignment where `Mapper` should be used, without that being the stated intent of
  the change.
- If a new service/repository interface was added, is it declared in the same package as its
  implementation (this repo's actual convention — see Go rules in `CLAUDE.md`), not split out
  into a separate "consumer" package?
- Is validation/business logic staying out of the controller, and persistence logic staying
  out of the service?
- Is a new domain wired the existing way — `internal/builder/services.go` +
  `internal/adapter/rest/routes_setup.go` — rather than constructed ad hoc in `cmd/main.go`?

### 4. Test coverage mapping
- For every new or changed exported method in the diff, is there a corresponding new/changed
  test? Point to the specific method with no test, don't just assert "coverage is low."
- Do mock expectations (`mockRepo.On("Method", ...)`) in a changed test still match the real
  method signature after the production change — a stale mock expectation compiles fine but
  fails at runtime with a mockery "I don't know what to return" panic.
- Does a repository-layer test exercise real `ent` code against the in-memory SQLite helper
  (`test.DbConnection`), consistent with how every other repository test in this repo is
  written — flag a repository test that mocks `ent` instead, since that's not this repo's
  pattern.
- Does a new error/negative branch in production code have a negative-path test, not just a
  happy-path one?
- (G1 only) Do new tests order negative cases before positive ones per the convention in
  `CLAUDE.md`? This is a style preference, not a correctness issue — never above G1, and
  don't flag it on tests the diff didn't touch.

### 5. Behaviour-preserving simplification
- Is the diff hand-rolling something a helper already does — DTO mapping instead of
  `helper.Mapper`, manual bind+validate instead of `helper.BindAndValidate`, a bespoke JSON
  error body instead of `internal/helper/response`?
- Does the diff introduce a new interface/abstraction with exactly one implementation and no
  second call site in sight? That's premature per this repo's "extend rather than abstract"
  rule — name the concrete type it could just be instead.
- Does the diff carry unrelated reformatting, renaming, or refactoring alongside the actual
  change? Point at the specific unrelated hunks — that's a minimal-diff violation, not a bug,
  but it should be called out so it can be split out.
- Is there a simpler, equally-correct version of the same logic already elsewhere in the
  codebase (e.g. another domain's controller/service pair) that this diff should have mirrored
  instead of solving differently?

## Reporting

Structure the report as:

1. **Gate: G<n>** — one line, the overall gate and why (point at the single highest finding).
2. **Findings**, most severe first, each with lens, gate, `file:line`, the quoted proof, why it
   matters, patch direction, and how to verify.
3. **Open checks** — things that smelled off but didn't clear the evidence bar; phrase these
   as questions, not accusations.
4. If everything is clean: say so in one or two sentences and stop. G0 is a valid, common, and
   honest outcome — don't stretch a nit into a finding just to have something to report.
