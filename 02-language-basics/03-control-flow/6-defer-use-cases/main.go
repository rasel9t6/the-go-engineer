// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: Defer in real use cases
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - See how `defer` is used in production for file cleanup, mutex unlocking, and logging.
//
// WHY THIS MATTERS:
//   - Think of `defer` as a "safety harness". Before you start a dangerous task (like opening a file), you put on the harness (`defer file.Close()`) so t...
//
// RUN:
//   go run ./02-language-basics/03-control-flow/6-defer-use-cases
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

//   Proper resource management is critical for production stability.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("CF.6: Real-World Defer Patterns")
	fmt.Println("--------------------------------")

	simulateFileOperation()

	// - Defer cleanup immediately after a resource is successfully acquired.
	// - This ensures that even if subsequent logic fails, resources are released.

	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: CF.7 pricing-checkout")
	fmt.Println("Current: CF.6 (defer-use-cases)")
	fmt.Println("---------------------------------------------------")
}

func simulateFileOperation() {
	fmt.Println("[Step 1] Opening simulated file 'data.txt'...")

	// Idiomatic pattern: defer the cleanup right after the "open" succeeds
	defer fmt.Println("[Step 4] File 'data.txt' closed automatically by defer.")

	fmt.Println("[Step 2] Reading data from file...")

	// Simulate some work or potential early return
	fmt.Println("[Step 3] Processing data complete.")
}
