// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: API versioning strategies
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn the trade-offs between URL versioning, header versioning, and evolutionary compatibility.
//
// WHY THIS MATTERS:
//   - Versioning is a change-management tool. The goal is controlled evolution, not version numbers everywhere.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/apis/2-api-versioning-strategies
//
// KEY TAKEAWAY:
//   - Learn the trade-offs between URL versioning, header versioning, and evolutionary compatibility.
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== API.2 API versioning strategies ===")
	fmt.Println("Learn the trade-offs between URL versioning, header versioning, and evolutionary compatibility.")
	fmt.Println()
	fmt.Println("- URL versions are explicit and easy to route.")
	fmt.Println("- Header versions keep URLs stable but are easier to miss.")
	fmt.Println("- The best versioning strategy is the one your clients can follow reliably.")
	fmt.Println()
	fmt.Println("Versioning policy should be boring, documented, and easy for clients to discover.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: API.3")
	fmt.Println("Current: API.2 (api versioning strategies)")
	fmt.Println("---------------------------------------------------")
}
