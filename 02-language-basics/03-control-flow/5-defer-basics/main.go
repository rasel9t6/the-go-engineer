// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: Defer — mechanics & order
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how to schedule work to happen at the very end of a function, ensuring cleanup always runs.
//
// WHY THIS MATTERS:
//   - Think of `defer` like a "sticky note" you put on the exit door of a room. No matter what you do in the room or which way you leave, you must perfor...
//
// RUN:
//   go run ./02-language-basics/03-control-flow/5-defer-basics
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

//   Defer is Go's primary mechanism for resource cleanup and safety.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("CF.5: Defer Mechanics")
	fmt.Println("--------------------------------")

	defer fmt.Println("  (Defer 1) This runs LAST")
	defer fmt.Println("  (Defer 2) This runs SECOND")

	fmt.Println("Performing main work...")
	fmt.Println("Work complete. Returning now.")

	// - defer schedules a function call to run when the surrounding function returns.
	// - Multiple defers run in Last-In-First-Out (LIFO) order.

	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: CF.6 defer-use-cases")
	fmt.Println("Current: CF.5 (defer-basics)")
	fmt.Println("---------------------------------------------------")
}
