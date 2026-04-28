// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 03: Functions and Errors
// Title: First-class functions
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn that functions are ordinary values in Go, which makes callbacks and higher-order helpers possible.
//
// WHY THIS MATTERS:
//   - A function value is just another tool you can store, pass, and call later.
//
// RUN:
//   go run ./03-functions-errors/8-first-class-functions
//
// KEY TAKEAWAY:
//   - Learn that functions are ordinary values in Go, which makes callbacks and higher-order helpers possible.
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== FE.8 First-class functions ===")
	fmt.Println("Learn that functions are ordinary values in Go, which makes callbacks and higher-order helpers possible.")
	fmt.Println()
	fmt.Println("- Assign functions to variables without calling them.")
	fmt.Println("- Pass behavior into other functions with callback parameters.")
	fmt.Println("- Read function signatures as contracts for what the caller must provide.")
	fmt.Println()
	fmt.Println("Callback-driven APIs stay readable only when function signatures are narrow and the names reveal the job each callback performs.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: FE.9")
	fmt.Println("Current: FE.8 (first-class functions)")
	fmt.Println("---------------------------------------------------")
}
