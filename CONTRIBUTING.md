# Contributing to The Go Engineer

Thank you for your interest in contributing. This guide explains the required workflow and quality bar.

## Quick Links

- [ARCHITECTURE.md](./ARCHITECTURE.md) — the locked v2.1 12-section source of truth
- [CURRICULUM-BLUEPRINT.md](./CURRICULUM-BLUEPRINT.md) — teaching and lesson contract
- [CODE-STANDARDS.md](./CODE-STANDARDS.md) — Go style, headers, comments, and review standards
- [TESTING-STANDARDS.md](./TESTING-STANDARDS.md) — test and verification expectations
- [AGENTS.md](./AGENTS.md) — repository operating contract for repeatable maintenance work

## Architecture Rule

The v2.1 architecture is locked. Contributions should deepen the existing 12 sections, not redesign the public structure.

Do not add, remove, rename, or reorder public root sections unless the maintainer explicitly approves reopening architecture.

## Strict GitHub Workflow

All contributors and maintainers follow the same workflow.

The workflow is closed-loop: implementation is not complete until review findings, CI failures, and PR conversation threads are handled. The final squash-merge remains maintainer-only.

### 1. Create or confirm an issue first

Never start implementation work without an approved issue.

The issue must include:

- affected section: `s00` through `s11`, if applicable
- affected lesson IDs, if applicable
- expected files or folders
- validation plan
- risk notes

Issue title format:

```text
[TYPE] short description
```

Allowed issue types:

```text
[FEAT] [FIX] [DOCS] [TEST] [CHORE] [REFACTOR] [SECURITY] [RELEASE]
```

Examples:

```text
[FEAT] add SY.1 sync.Mutex lesson
[FIX] correct SQL injection example in DB.3
[DOCS] align ROADMAP.md with v2.1 final status
[TEST] add TE.4 benchmark coverage
[CHORE] tighten curriculum validator
```

### 2. Branch from the correct line

Long-lived branches:

- `main`: post-v2.1 implementation line
- `release/v1`: stable v1 maintenance line
- `release/v2`: stable v2.1.x maintenance line

For normal v2.1 work:

```bash
git switch main
git pull origin main
git switch -c <branch-name>
```

Branch naming:

| Pattern | Use |
| --- | --- |
| `feat/sNN-lesson-slug` | new lesson or lesson expansion |
| `fix/sNN-description` | bug fix in a section |
| `docs/document-or-section` | documentation-only change |
| `test/sNN-description` | test-only improvement |
| `chore/tooling-description` | tooling, CI, dependencies |
| `release/vX.Y.Z-prep` | release metadata work |

### 3. Open a draft PR early

Open a draft PR after the first push. Link the issue with `Closes #<issue>`.

The PR should include:

- scope
- affected section and lesson IDs
- validation commands run
- risk notes
- screenshots or output when useful

### 4. Keep commits logical

Use bracketed commit messages:

```text
[TYPE] short imperative description
```

Examples:

```text
[FEAT] add SY.1 sync.Mutex lesson
[TEST] add SY.1 race-focused tests
[DOCS] update s07 section map
[FIX] scope DB.3 example query by tenant id
[CHORE] align local verification with CI
```

Keep unrelated formatting changes out of feature commits.

### 5. Verify locally before review

Run the CI-equivalent checks:

```bash
go build ./...
go vet ./...
unformatted=$(gofmt -l .); test -z "$unformatted" || (echo "$unformatted" && exit 1)
go mod tidy
git diff --exit-code -- go.mod go.sum
go test ./...
go test -race ./...
go test -coverprofile=coverage.out ./...
go run ./scripts/validate_curriculum.go
```

For benchmark-related changes:

```bash
go test -bench=. -benchmem -count=1 ./08-quality-test/01-quality-and-performance/testing/benchmarks/
```

### 6. Final review and merge

When ready:

- mark the PR ready for review
- wait for maintainer review
- use **Squash and Merge**
- do not self-merge unless the maintainer explicitly approves it for that PR

The final squash commit should also use the bracketed format.

## Lesson File Template

Every Go lesson file must follow the standard header and footer contract.

```go
// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section NN: Section Name — Lesson Title
// Level: Foundation | Core | Stretch
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Concept one
//   - Concept two
//
// WHY THIS MATTERS:
//   One sentence on why this exists in a real production Go codebase.
//
// RUN: go run ./NN-section-slug/N-lesson-slug
// ============================================================================

package main

import "fmt"

func main() {
	// Your code with inline comments explaining every concept.
	fmt.Println("Hello")

	// KEY TAKEAWAY:
	// - Summary point 1
	// - Summary point 2
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: XY.N next-lesson-slug")
	fmt.Println("Run    : go run ./NN-section-slug/N-next-lesson-slug")
	fmt.Println("Current: XY.N current-lesson-slug")
	fmt.Println("---------------------------------------------------")
}
```

`NEXT UP:` must use this exact text with no emoji. It must match `curriculum.v2.json`.

## Comment Guidelines

1. Explain why, not obvious what.
2. Use comments to teach non-obvious behavior, traps, and runtime implications.
3. Cross-reference lessons with lesson IDs where helpful.
4. Keep source files readable; the README remains the primary teaching surface.
5. End `main()` with a `KEY TAKEAWAY` and the standard footer where applicable.

## Adding a Lesson

1. Confirm an approved issue.
2. Create the lesson directory inside the existing section.
3. Write `README.md` first using the lesson contract.
4. Write `main.go` using the standard source contract.
5. Add tests and starter code when the item is an exercise.
6. Register the item in `curriculum.v2.json`.
7. Update the section README.
8. Run validation.

## Adding an Exercise

Exercises usually include:

```text
NN-section-slug/
└── N-exercise-name/
    ├── README.md
    ├── main.go
    ├── main_test.go
    └── _starter/
        └── main.go
```

The `_starter/main.go` file must compile successfully and contain clear TODO markers.

## Quality Checklist

Before final review:

- [ ] Architecture v2.1 remains intact.
- [ ] `curriculum.v2.json` is updated for lesson changes.
- [ ] Section README is updated when needed.
- [ ] Lesson README follows the required section order.
- [ ] Run commands work exactly as printed.
- [ ] Tests prove behavior where behavior should be provable.
- [ ] Starter code compiles.
- [ ] `NEXT UP:` footer matches curriculum metadata.
- [ ] `go build ./...` passes.
- [ ] `go vet ./...` passes.
- [ ] gofmt check passes.
- [ ] `go mod tidy` produces no `go.mod` or `go.sum` diff.
- [ ] `go test ./...` passes.
- [ ] `go test -race ./...` passes.
- [ ] `go run ./scripts/validate_curriculum.go` passes.
- [ ] CI passes on GitHub.
