// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: net/http basics
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn the smallest useful shape of an HTTP server in Go.
//
// WHY THIS MATTERS:
//   - A handler is the unit of HTTP work. A mux decides which handler should receive each request.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/http-servers/1-net-http-basics
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== HS.1 net/http basics ===")
	fmt.Println("Learn the smallest useful shape of an HTTP server in Go.")
	fmt.Println()
	fmt.Println("- Handlers are ordinary functions or interface implementations.")
	fmt.Println("- ServeMux routes requests by pattern.")
	fmt.Println("- ListenAndServe owns the accept loop for incoming connections.")
	fmt.Println()
	fmt.Println("The stdlib server is enough for most internal services and a strong baseline for external APIs too.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: HS.2")
	fmt.Println("Current: HS.1 (net/http basics)")
	fmt.Println("---------------------------------------------------")
}
