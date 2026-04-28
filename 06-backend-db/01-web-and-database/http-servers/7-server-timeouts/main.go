// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: Server timeouts
// Level: Production
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn why HTTP servers need read, header, write, and idle timeouts to stay safe under pressure.
//
// WHY THIS MATTERS:
//   - Timeouts are resource guards. They stop one slow or malicious client from holding the server forever.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/http-servers/7-server-timeouts
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== HS.7 Server timeouts ===")
	fmt.Println("Learn why HTTP servers need read, header, write, and idle timeouts to stay safe under pressure.")
	fmt.Println()
	fmt.Println("- ReadHeaderTimeout protects header parsing.")
	fmt.Println("- ReadTimeout limits the full body-read window.")
	fmt.Println("- WriteTimeout and IdleTimeout protect response and keep-alive resources.")
	fmt.Println()
	fmt.Println("Missing timeout settings creates real attack surfaces, especially slowloris-style header drips and slow response consumers.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: HS.8")
	fmt.Println("Current: HS.7 (server timeouts)")
	fmt.Println("---------------------------------------------------")
}
