// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: Error handling middleware
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how central error translation keeps handlers thin and HTTP behavior consistent.
//
// WHY THIS MATTERS:
//   - Handlers produce domain errors; middleware maps those errors into status codes and response bodies.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/http-servers/6-error-handling-middleware
//
// KEY TAKEAWAY:
//   - Learn how central error translation keeps handlers thin and HTTP behavior consistent.
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== HS.6 Error handling middleware ===")
	fmt.Println("Learn how central error translation keeps handlers thin and HTTP behavior consistent.")
	fmt.Println()
	fmt.Println("- Treat validation, system, and bug failures differently.")
	fmt.Println("- Translate domain errors at the HTTP boundary.")
	fmt.Println("- Keep handlers focused on returning useful errors, not formatting responses.")
	fmt.Println()
	fmt.Println("HTTP APIs become easier to audit when one layer decides how user, system, and bug errors are rendered.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: HS.7")
	fmt.Println("Current: HS.6 (error handling middleware)")
	fmt.Println("---------------------------------------------------")
}
