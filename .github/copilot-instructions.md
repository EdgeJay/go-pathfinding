# copilot-instructions.md

> **Audience:** GitHub Copilot / VS Code Copilot
> **Project type:** Go-based data algorithms project
> **Prime directive:** Always use `docs/PLANNING.md` as the single source of truth for scope, milestones, and priorities.

---

## 0) Always Load Project Plan First

* **Find and read** the high-level planning document at **`docs/PLANNING.md`** before proposing code, tasks, tests, or docs.
* **In every prompt you act on**, include a short reference to `docs/PLANNING.md` (e.g., “per docs/PLANNING.md: Milestone M1, Algorithm A, constraints X/Y”).
* If `docs/PLANNING.md` is missing details, **ask to refine the plan** (do not invent scope).

---

## 1) Task Management (docs/TASKS.md)

* Maintain a single project task list at **`docs/TASKS.md`** as a **Git-style checkbox list**.
* **Add, update, and tick** items there—never scatter tasks in other files.
* Organize by milestone/epic drawn from `docs/PLANNING.md`.

**Template for `docs/TASKS.md`:**

```markdown
# Project Tasks (auto-curated by Copilot; source of truth)

> All tasks must trace back to docs/PLANNING.md. Use checkboxes. Keep sections small and current.

## Milestone: M1 – Core Algorithm
- [ ] Design data model for <component> (ref: docs/PLANNING.md §M1)
- [ ] Implement algorithm A v1 in package `algo` with benchmarks
- [ ] Add unit tests for edge cases (empty input, large N, invalid data)
- [ ] Document public API functions in `algo` (GoDoc style)
- [ ] Update README usage snippet

## Milestone: M2 – I/O & Integration
- [ ] Parse input from <format> with validation
- [ ] Provide CLI command `cmd/app` with flags
- [ ] Add integration test covering end-to-end flow
- [ ] Update README run instructions

## Housekeeping
- [ ] CI: run `go test ./... -race -coverprofile=coverage.out`
- [ ] Lint: `golangci-lint run`
- [ ] Benchmarks recorded in `docs/benchmarks/`
```

---

## 2) Code Organization & Style

* Use idiomatic **Go modules** and packages; keep public surface minimal.
* Prefer clear, benchmarked implementations over cleverness.
* Enforce `go fmt`, `go vet`, and a linter (e.g., `golangci-lint`).
* For performance work, provide micro-benchmarks (`testing.B`) and capture results in `docs/benchmarks/`.

---

## 3) Public API Documentation (Required)

For **every exported (public) function, type, const, var**:

* Add **GoDoc-style** comments **immediately above** the declaration.
* Include **purpose**, **parameters**, **return values**, **error cases**, and **complexity** if relevant.
* Provide a short **usage example** in `*_test.go` or an `example_test.go`.

**Comment example:**

```go
// TopK selects the k highest-scoring elements from scores using a partial selection algorithm.
// It runs in O(n) average time and O(1) extra space.
// Params:
//   scores: non-nil slice of float64
//   k: 0 <= k <= len(scores)
// Returns:
//   indices: indices of the selected elements in descending score order
// Errors:
//   ErrInvalidK if k is out of range.
func TopK(scores []float64, k int) ([]int, error) { ... }
```

---

## 4) Testing Policy (Unit, Property, Integration)

* **Unit tests are mandatory** for all public functions and any non-trivial private helpers.
* Cover **edge cases** (empty, nil, extremes, skewed distributions) and **error paths**.
* Use **table-driven tests**; prefer **property-based tests** where applicable.
* Target **≥90%** coverage for the `algo` package; never regress without justification.
* Include **benchmarks** for performance-critical code; track results in `docs/benchmarks/`.

**Commands to support:**

```bash
go test ./... -race -coverprofile=coverage.out
go tool cover -func=coverage.out
go test ./algo -bench=. -benchmem
```

---

## 5) README Requirements (Project Root)

Ensure a **`README.md`** exists and stays current with `docs/PLANNING.md`.

**README must include:**

### a) Pre-requisites

* Supported OS / architectures
* Go toolchain version (e.g., `go >= 1.xx`)
* Any system libraries or tools (e.g., `make`, `gcc` for CGO if needed)
* Optional: `golangci-lint`, `just` or `make` for tasks

### b) Getting Started

* Clone, build, test:

  ```bash
  git clone <repo>
  cd <repo>
  go mod download
  go build ./cmd/app
  go test ./... -race
  ```
* Run the app (CLI flags, config file path, env vars):

  ```bash
  ./app --input=data/sample.json --k=10 --mode=fast
  ```
* Example input/output snippets
* Notes on performance tuning and resource usage

**README skeleton Copilot should maintain:**

````markdown
# Project Name – Data Algorithms in Go

> Scope & milestones per [../docs/PLANNING.md](../docs/PLANNING.md).

## Pre-requisites
- Go >= 1.xx
- (Optional) golangci-lint
- OS: macOS/Linux/Windows (amd64/arm64)

## Quick Start
```bash
go mod download
go build ./cmd/app
./app --help
````

## Usage

```bash
./app --input=PATH --k=10 --mode=fast
```

## Development

* Run tests: `go test ./... -race`
* Lint: `golangci-lint run`
* Benchmarks: `go test ./algo -bench=. -benchmem`

## Architecture

* Packages: `algo/`, `cmd/app/`, `internal/...`
* See [../docs/PLANNING.md](../docs/PLANNING.md) for roadmap.

## License

MIT (or as specified)

```

---

## 6) Working Protocol for Copilot

When generating or modifying code, **follow this loop**:

1. **Read `docs/PLANNING.md`** → extract goal, acceptance criteria, and constraints.  
2. **Update `docs/TASKS.md`** → add/adjust checkbox items tied to the current goal.  
3. **Propose design** (comments or short design doc in `docs/design/<feature>.md` if complex).  
4. **Implement code** in minimal increments with clear public API and comments.  
5. **Write tests** (unit + property + benchmarks where relevant).  
6. **Run & fix**: `go test ./... -race`, lint, and benchmark.  
7. **Update README** if usage or setup changed.  
8. **Mark tasks done** in `docs/TASKS.md` and reference commits/PRs.
9. **Coaching mode** → if received instructions to explain "step-by-step", always pause when adding/modifying code to provide a brief rationale for design choices, trade-offs, and testing strategies and wait for confirmation before proceeding.

---

## 7) File & Folder Conventions

```

.
├── README.md                  # Must cover pre-reqs & start instructions
├── docs/
│   ├── PLANNING.md            # Source of truth for scope/milestones
│   ├── TASKS.md               # Checkbox list, maintained continuously
│   ├── benchmarks/            # Benchmark outputs, dated
│   └── design/                # Optional mini design docs per feature
├── algo/                      # Core algorithms (exported APIs documented)
│   ├── topk.go
│   └── topk_test.go
├── cmd/
│   └── app/                   # CLI entry
│       └── main.go
└── internal/                  # Non-exported helpers

```

---

## 8) Commit / PR Hygiene

- Reference **`docs/PLANNING.md`** section or task IDs from **`docs/TASKS.md`** in commit messages.
- PR checklist:
  - [ ] Public functions documented
  - [ ] Tests added/updated; `-race` green; coverage not regressed
  - [ ] README updated if behavior or setup changed
  - [ ] Tasks checked/updated in `docs/TASKS.md`

---

## 9) Prompt Snippets for Copilot (Use/Adapt)

- “Per `docs/PLANNING.md` §M1, implement `algo.TopK`. Create table-driven tests, property tests for ordering, and a benchmark; document public API; add tasks to `docs/TASKS.md`.”
- “Using `docs/PLANNING.md` constraints, add a CLI in `cmd/app` with flags `--input`, `--k`, and `--mode`. Update README quick-start and usage. Track tasks in `docs/TASKS.md`.”
- “From `docs/PLANNING.md` performance goals, propose and implement an in-place partition strategy; add `BenchmarkTopK*`; record results in `docs/benchmarks/` and update tasks.”

---

## 10) Non-negotiables (Fast Check)

- ✅ **Always reference `docs/PLANNING.md`** in prompts and decisions  
- ✅ **Tasks live only in `docs/TASKS.md`** (checkbox list)  
- ✅ **Public functions documented** (GoDoc comments)  
- ✅ **Proper unit tests** (plus property/benchmarks where applicable)  
- ✅ **`README.md` present** with pre-reqs and start instructions

> If any of the above cannot be satisfied, pause and request plan clarification in `docs/PLANNING.md` before coding further.
```
