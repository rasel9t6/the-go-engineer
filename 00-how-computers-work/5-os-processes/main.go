// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 00: How Computers Work — OS Processes
// Level: Foundation
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - How to get the Process ID (PID)
//   - The program is running as an OS process
//
// WHY THIS MATTERS:
//   Production systems track services by their PIDs.
//
// RUN: go run ./00-how-computers-work/5-os-processes
// ============================================================================

package main

import (
	"fmt"
	"os"
)

func main() {
	// The OS gives every running process a unique ID
	pid := os.Getpid()
	fmt.Printf("This Go program is running as Process ID: %d\n", pid)

	// KEY TAKEAWAY:
	// - Go programs run as standard OS processes
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("NEXT UP: GT.1 1-installation")
	fmt.Println("Run    : go run ./01-getting-started/1-installation")
	fmt.Println("Current: HC.5 (5-os-processes)")
	fmt.Println("---------------------------------------------------")
}
