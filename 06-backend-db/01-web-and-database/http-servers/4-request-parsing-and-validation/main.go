// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: Request parsing and validation
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how to decode request data deliberately and reject malformed input early.
//
// WHY THIS MATTERS:
//   - Parsing turns transport bytes into domain values. Validation decides whether those values are safe to trust.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/http-servers/4-request-parsing-and-validation
//
// KEY TAKEAWAY:
//   - Learn how to decode request data deliberately and reject malformed input early.
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== HS.4 Request parsing and validation ===")
	fmt.Println("Learn how to decode request data deliberately and reject malformed input early.")
	fmt.Println()
	fmt.Println("- Parse one request surface at a time.")
	fmt.Println("- Limit body size before decoding it.")
	fmt.Println("- Return validation feedback before business logic runs.")
	fmt.Println()
	fmt.Println("Most API bugs start at the boundary. Size limits, decode errors, and validation checks should fail fast.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: HS.5")
	fmt.Println("Current: HS.4 (request parsing and validation)")
	fmt.Println("---------------------------------------------------")
}
