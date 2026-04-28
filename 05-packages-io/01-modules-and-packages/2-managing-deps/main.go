// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 05: Packages and I/O
// Title: Managing Dependencies
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Managing Dependencies fundamentals and practical application in Go.
//
// WHY THIS MATTERS:
//   - Managing Dependencies provides a structured approach to writing clean Go code.
//
// RUN:
//   go run ./05-packages-io/01-modules-and-packages/2-managing-deps
//
// KEY TAKEAWAY:
//   - Managing Dependencies fundamentals and practical application in Go.
// ============================================================================

// Commercial use is prohibited without permission.

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// Stage 05: Modules and Packages â€” Managing Dependencies
//
// This file demonstrates dependency-management workflows.
// Run the commands below in a terminal to see them in action.
//
// ADDING DEPENDENCIES:
//   go get github.com/stretchr/testify@latest        â€” latest version
//   go get github.com/stretchr/testify@v1.11.1       â€” specific version
//   go get github.com/stretchr/testify@v1.10.0       â€” downgrade
//
// REMOVING DEPENDENCIES:
//   go get github.com/some/pkg@none                  â€” remove from go.mod
//   go mod tidy                                      â€” clean up unused
//
// INSPECTING DEPENDENCIES:
//   go list -m all                                   â€” all modules
//   go list -m -versions github.com/stretchr/testify â€” available versions
//   go mod why github.com/stretchr/objx              â€” why is this needed?
//   go mod graph                                     â€” full dependency tree
//
// SECURITY:
//   go install golang.org/x/vuln/cmd/govulncheck@latest
//   govulncheck ./...                                â€” scan for known vulnerabilities

func main() {
	fmt.Println("=== Managing Dependencies ===")
	fmt.Println()

	fmt.Println("Current module dependencies:")
	fmt.Println(strings.Repeat("-", 60))

	out, err := exec.Command("go", "list", "-m", "all").Output()
	if err != nil {
		fmt.Printf("Error running 'go list -m all': %v\n", err)
		fmt.Println("(This can happen if CGO-backed dependencies are unavailable.)")
		return
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "the-go-engineer") {
			fmt.Printf("  [ROOT] %s\n", line)
			continue
		}
		fmt.Printf("  [DEP]  %s\n", line)
	}

	fmt.Println()
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("\nWhy do we have github.com/stretchr/objx?")

	whyOut, err := exec.Command("go", "mod", "why", "github.com/stretchr/objx").Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(string(whyOut))
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("ðŸš€ NEXT UP: MP.3 versioning")
	fmt.Println("   Current: MP.2 (managing deps)")
	fmt.Println("---------------------------------------------------")
}
