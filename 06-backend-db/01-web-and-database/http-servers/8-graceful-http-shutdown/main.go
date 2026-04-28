// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: Graceful HTTP shutdown
// Level: Production
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how to stop accepting new work while allowing in-flight requests to finish cleanly.
//
// WHY THIS MATTERS:
//   - Graceful shutdown is a drain: stop the front door, wait for active requests, then release the process.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/http-servers/8-graceful-http-shutdown
//
// KEY TAKEAWAY:
//   - Learn how to stop accepting new work while allowing in-flight requests to finish cleanly.
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== HS.8 Graceful HTTP shutdown ===")
	fmt.Println("Learn how to stop accepting new work while allowing in-flight requests to finish cleanly.")
	fmt.Println()
	fmt.Println("- Stop accepting new requests before terminating the process.")
	fmt.Println("- Use contexts and deadlines to bound the drain window.")
	fmt.Println("- Design handlers so cancellation and shutdown signals can be respected.")
	fmt.Println()
	fmt.Println("Rolling deploys, autoscaling, and restarts all depend on shutdown behavior that does not drop useful work unnecessarily.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: HS.9")
	fmt.Println("Current: HS.8 (graceful http shutdown)")
	fmt.Println("---------------------------------------------------")
}
