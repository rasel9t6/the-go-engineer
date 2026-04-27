// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: gRPC streaming
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how server, client, and bidirectional streams change the contract from one message to many.
//
// WHY THIS MATTERS:
//   - Streaming turns one call into a long-lived message flow rather than a single request/response pair.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/apis/6-grpc-streaming
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== API.6 gRPC streaming ===")
	fmt.Println("Learn how server, client, and bidirectional streams change the contract from one message to many.")
	fmt.Println()
	fmt.Println("- Server streams push many updates from one request.")
	fmt.Println("- Client streams batch or trickle many inputs into one response.")
	fmt.Println("- Bidirectional streams allow both sides to exchange messages over one call.")
	fmt.Println()
	fmt.Println("Streaming is powerful but should be chosen for ongoing flows, not just because it looks more advanced.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: API.7")
	fmt.Println("Current: API.6 (grpc streaming)")
	fmt.Println("---------------------------------------------------")
}
