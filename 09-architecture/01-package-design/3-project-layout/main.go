// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 09: Architecture & Security
// Title: Project Layout
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Decide when a Go project should stay flat, when it needs `internal/`, and when `cmd/` or other top-level directories actually earn their place. Thi...
//
// WHY THIS MATTERS:
//   - Decide when a Go project should stay flat, when it needs `internal/`, and when `cmd/` or other top-level directories actually earn their place. Thi...
//
// RUN:
//   go run ./09-architecture/01-package-design/3-project-layout
//
// KEY TAKEAWAY:
//   - Decide when a Go project should stay flat, when it needs `internal/`, and when `cmd/` or other top-level directories actually earn their place. Thi...
// ============================================================================

// Commercial use is prohibited without permission.

package main

import "fmt"

// Stage 09: Application Architecture - Package Design: Standard Go Project Layout
//
//   - The common directory conventions for Go projects
//   - When to use cmd/, internal/, pkg/, and other directories
//   - Simple vs complex project layouts
//   - Anti-patterns to avoid
//
// ENGINEERING DEPTH:
//   Go does not enforce one project structure, but the community has converged
//   on a small set of practical conventions. The most important one is
//   `internal/`, because the compiler uses it to block external imports.
//

func main() {
	fmt.Println("=== Standard Go Project Layout ===")
	fmt.Println()

	fmt.Println("ðŸ“ SIMPLE PROJECT (library or small app)")
	fmt.Println("   Perfect for: libraries, small tools, learning")
	fmt.Println()
	fmt.Println("   mylib/")
	fmt.Println("   â”œâ”€â”€ go.mod")
	fmt.Println("   â”œâ”€â”€ go.sum")
	fmt.Println("   â”œâ”€â”€ mylib.go          <- main package code")
	fmt.Println("   â”œâ”€â”€ mylib_test.go     <- tests alongside code")
	fmt.Println("   â””â”€â”€ README.md")
	fmt.Println()

	fmt.Println("ðŸ“ MEDIUM PROJECT (single binary + some packages)")
	fmt.Println("   Perfect for: web servers, CLI tools, microservices")
	fmt.Println()
	fmt.Println("   myapp/")
	fmt.Println("   â”œâ”€â”€ go.mod")
	fmt.Println("   â”œâ”€â”€ main.go           <- entry point (package main)")
	fmt.Println("   â”œâ”€â”€ internal/         <- private packages (compiler-enforced)")
	fmt.Println("   â”‚   â”œâ”€â”€ auth/")
	fmt.Println("   â”‚   â”‚   â”œâ”€â”€ auth.go")
	fmt.Println("   â”‚   â”‚   â””â”€â”€ auth_test.go")
	fmt.Println("   â”‚   â”œâ”€â”€ store/")
	fmt.Println("   â”‚   â”‚   â”œâ”€â”€ store.go")
	fmt.Println("   â”‚   â”‚   â””â”€â”€ store_test.go")
	fmt.Println("   â”‚   â””â”€â”€ handler/")
	fmt.Println("   â”‚       â”œâ”€â”€ handler.go")
	fmt.Println("   â”‚       â””â”€â”€ handler_test.go")
	fmt.Println("   â”œâ”€â”€ Makefile")
	fmt.Println("   â””â”€â”€ README.md")
	fmt.Println()

	fmt.Println("ðŸ“ LARGE PROJECT (multiple binaries, shared code)")
	fmt.Println("   Perfect for: monorepos, complex systems, multi-service apps")
	fmt.Println()
	fmt.Println("   platform/")
	fmt.Println("   â”œâ”€â”€ go.mod")
	fmt.Println("   â”œâ”€â”€ cmd/              <- each subdirectory is a binary")
	fmt.Println("   â”‚   â”œâ”€â”€ api/")
	fmt.Println("   â”‚   â”‚   â””â”€â”€ main.go   <- go build ./cmd/api")
	fmt.Println("   â”‚   â”œâ”€â”€ worker/")
	fmt.Println("   â”‚   â”‚   â””â”€â”€ main.go   <- go build ./cmd/worker")
	fmt.Println("   â”‚   â””â”€â”€ migrate/")
	fmt.Println("   â”‚       â””â”€â”€ main.go   <- go build ./cmd/migrate")
	fmt.Println("   â”œâ”€â”€ internal/         <- private shared packages")
	fmt.Println("   â”‚   â”œâ”€â”€ auth/")
	fmt.Println("   â”‚   â”œâ”€â”€ store/")
	fmt.Println("   â”‚   â””â”€â”€ email/")
	fmt.Println("   â”œâ”€â”€ pkg/              <- public shared packages")
	fmt.Println("   â”‚   â””â”€â”€ middleware/")
	fmt.Println("   â”œâ”€â”€ migrations/       <- SQL migration files")
	fmt.Println("   â”œâ”€â”€ Makefile")
	fmt.Println("   â”œâ”€â”€ Dockerfile")
	fmt.Println("   â””â”€â”€ README.md")
	fmt.Println()

	fmt.Println("=== Directory Purpose Guide ===")
	dirs := []struct {
		dir  string
		use  string
		when string
	}{
		{"cmd/", "Entry points (package main)", "Multiple binaries in one repo"},
		{"internal/", "Private packages", "Code that must not be imported externally"},
		{"pkg/", "Public shared packages", "Code designed for reuse by other modules"},
		{"migrations/", "Database migration files", "Apps with SQL databases"},
		{"testdata/", "Test fixtures", "Go test tooling ignores this directory"},
		{"scripts/", "Build/deploy/CI scripts", "Automation beyond Makefile"},
		{"docs/", "Additional documentation", "Complex projects needing detailed docs"},
	}

	for _, d := range dirs {
		fmt.Printf("  %-15s - %s\n", d.dir, d.use)
		fmt.Printf("  %15s   When: %s\n", "", d.when)
		fmt.Println()
	}

	fmt.Println("=== Anti-Patterns (Avoid These) ===")
	fmt.Println("  âŒ src/                - Go does not use src/ the way Java does")
	fmt.Println("  âŒ models/             - Too vague. Name by domain: user/, order/")
	fmt.Println("  âŒ utils/ or helpers/  - Junk drawer. Split by responsibility")
	fmt.Println("  âŒ Over-engineering    - Do not use cmd/internal/pkg/ for a 200-line app")
	fmt.Println()

	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("  1. Start simple and add structure as the project grows")
	fmt.Println("  2. cmd/ is for multiple binaries, internal/ is for private packages")
	fmt.Println("  3. pkg/ is optional - use it only for code meant to be shared publicly")
	fmt.Println("  4. Tests live next to the code they verify (user.go -> user_test.go)")
	fmt.Println("  5. Do not cargo-cult a complex layout for a small project")
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("ðŸš€ NEXT UP: ARCH.1")
	fmt.Println("   Current: PD.3 (project layout)")
	fmt.Println("---------------------------------------------------")
}
