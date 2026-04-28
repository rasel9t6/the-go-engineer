// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: REST API
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Build a small REST-shaped surface that combines routing, middleware, validation, and consistent responses.
//
// WHY THIS MATTERS:
//   - A real API is the composition of many smaller rules: routing, parsing, error handling, and stable payload contracts.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/http-servers/10-rest-api-exercise
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== HS.10 REST API ===")
	fmt.Println("Build a small REST-shaped surface that combines routing, middleware, validation, and consistent responses.")
	fmt.Println()
	fmt.Println("- Combine middleware, validation, and JSON responses in one small service.")
	fmt.Println("- Use consistent status codes for create, read, and failure paths.")
	fmt.Println("- Treat the exercise as a proof that each earlier lesson solved part of the final shape.")
	fmt.Println()
	fmt.Println("Exercise surfaces are where transport choices become habits that carry into real service work.")
}
