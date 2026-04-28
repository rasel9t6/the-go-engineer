// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 07: Concurrency
// Title: Race conditions
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how data races happen, why the race detector matters, and how to read its reports.
//
// WHY THIS MATTERS:
//   - A race is unsynchronized shared access where at least one access is a write.
//
// RUN:
//   go run ./07-concurrency/01-concurrency/sync-primitives/4-race-conditions
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== SY.4 Race conditions ===")
	fmt.Println("Learn how data races happen, why the race detector matters, and how to read its reports.")
	fmt.Println()
	fmt.Println("- Shared mutable state needs synchronization.")
	fmt.Println("- Lost updates and stale reads come from interleaving, not just from high load.")
	fmt.Println("- The race detector is a normal part of verification, not a last resort.")
	fmt.Println()
	fmt.Println("The race detector is one of the highest-value debugging tools in Go. Use it routinely before you trust concurrent code.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: SY.5")
	fmt.Println("Current: SY.4 (race conditions)")
	fmt.Println("---------------------------------------------------")
}
