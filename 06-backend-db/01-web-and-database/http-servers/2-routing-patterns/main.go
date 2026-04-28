// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: Routing patterns
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how route patterns, method checks, and path parsing shape request flow.
//
// WHY THIS MATTERS:
//   - Routing is the first decision boundary in a server: find the path, check the verb, then pass the request into domain logic.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/http-servers/2-routing-patterns
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== HS.2 Routing patterns ===")
	fmt.Println("Learn how route patterns, method checks, and path parsing shape request flow.")
	fmt.Println()
	fmt.Println("- Separate path matching from business behavior.")
	fmt.Println("- Reject wrong HTTP methods early and clearly.")
	fmt.Println("- Treat path parsing as input validation, not as routing magic.")
	fmt.Println()
	fmt.Println("Route clarity matters because it keeps handlers focused and makes debugging mismatches fast.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: HS.3")
	fmt.Println("Current: HS.2 (routing patterns)")
	fmt.Println("---------------------------------------------------")
}
