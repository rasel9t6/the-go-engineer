// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: Constants
// Level: Foundation
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how Go represents values that should never change at runtime.
//
// WHY THIS MATTERS:
//   - A variable can change while the program runs. A constant cannot. Constants communicate: - this value is fixed by design - the compiler can treat it...
//
// RUN:
//   go run ./02-language-basics/2-constants
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

const (
	Host = "127.0.0.1"
	Port = ":8080"
	User = "root"
)

var (
	isRunning bool
)

func main() {
	AppName := "Go"
	fmt.Println(AppName)

	const pi float64 = 3.1415926
	fmt.Println(pi)

	const rate float32 = 5.2
	fmt.Println(rate)

	fmt.Printf("Server: %s%s (User: %s)\n", Host, Port, User)
	fmt.Printf("Running: %t\n", isRunning)

	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: LB.3 enums")
	fmt.Println("Current: LB.2 (constants)")
	fmt.Println("---------------------------------------------------")
}
