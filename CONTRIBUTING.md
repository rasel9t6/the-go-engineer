# Contributing to The Go Engineer

Thank you for your interest in contributing! This guide will help you get started.

## Getting Started

```bash
# Clone the repository
git clone https://github.com/rasel9t6/the-go-engineer.git
cd the-go-engineer

# Verify your environment
make build
make test
make lint
```

## Code Style

### Every Go File Must Follow This Template

```go
package main

import "fmt"

// ============================================================================
// Section N: Section Name — Lesson Title
// Level: Beginner | Intermediate | Advanced
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Point 1
//   - Point 2
//
// RUN: go run ./NN-section-name/N-lesson-name
// ============================================================================

func main() {
    // Your code with inline comments explaining every concept
    fmt.Println("Hello")

    // KEY TAKEAWAY:
    // - Summary point 1
    // - Summary point 2
}
```

### Comment Guidelines

1. **Every non-obvious line gets a comment** — This is a teaching repo, not a production codebase
2. **Explain WHY, not WHAT** — `// Increment i` is useless. `// Move to the next element because...` teaches
3. **Use block comments for concepts** — Multi-line `//` comments above code blocks to explain the concept before showing the code
4. **Include KEY TAKEAWAY** — Every `main()` function should end with a KEY TAKEAWAY summary
5. **Cross-reference sections** — Use `(See Section 05: Interfaces)` when a concept is covered elsewhere

### Formatting Rules

- Run `gofmt` on every file (or use `make fmt`)
- Run `go vet ./...` before committing (or use `make vet`)
- Run `make lint` to verify all checks pass

## Section Numbering

- Sections are numbered `00-22` with descriptive names
- Lessons within a section are numbered `1-`, `2-`, etc.
- Exercises are the LAST numbered item in a section

## Adding a New Lesson

1. Create a directory: `NN-section-name/N-lesson-name/`
2. Create `main.go` following the template above
3. Update the section's `README.md` with the new lesson in the contents table
4. Update the root `README.md` if adding a new section
5. Verify: `make build && make test && make lint`

## Adding an Exercise

Exercises include both starter code and solutions:

```
NN-section-name/
└── N-exercise-name/
    ├── main.go              ← Starter code with TODO comments
    └── _solution/
        └── main.go          ← Complete solution with comments
```

## Commit Messages

Use clear, descriptive commit messages:

```
Add: Section 17 Context deep-dive (4 lessons)
Fix: go vet warning in 07-strings formatting
Update: Backfill comments in Section 02 control flow
```

## Quality Checklist

Before submitting, verify:

- [ ] `make build` passes (all packages compile)
- [ ] `make vet` passes (no suspicious code)
- [ ] `make fmt-check` passes (code is formatted)
- [ ] `make test` passes (all tests pass)
- [ ] Every new Go file has the standard header template
- [ ] Every concept has inline teaching comments
- [ ] The section README is updated
