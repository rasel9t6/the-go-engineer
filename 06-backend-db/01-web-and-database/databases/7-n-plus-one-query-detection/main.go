// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: N+1 query detection
// Level: Production
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how repeated lookups inside a loop turn one logical query into unbounded database chatter.
//
// WHY THIS MATTERS:
//   - The N+1 problem happens when one outer query hides many inner queries behind a loop.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/databases/7-n-plus-one-query-detection
//
// KEY TAKEAWAY:
//   - Learn how repeated lookups inside a loop turn one logical query into unbounded database chatter.
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== DB.7 N+1 query detection ===")
	fmt.Println("Learn how repeated lookups inside a loop turn one logical query into unbounded database chatter.")
	fmt.Println()
	fmt.Println("- Count queries as part of performance debugging.")
	fmt.Println("- Spot repeated lookup patterns hidden inside loops.")
	fmt.Println("- Prefer joins, bulk fetches, or preloading when the access pattern is known.")
	fmt.Println()
	fmt.Println("N+1 bugs usually surface first as latency spikes and database saturation under real load.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: DB.8")
	fmt.Println("Current: DB.7 (n+1 query detection)")
	fmt.Println("---------------------------------------------------")
}
