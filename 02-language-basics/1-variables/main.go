// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: Variables and Types
// Level: Foundation
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how Go declares variables and why every type has a predictable zero value.
//
// WHY THIS MATTERS:
//   - A variable is a named slot that holds a value while the program runs. Go gives you three common declaration shapes: 1. `var name string` 2. `var na...
//
// RUN:
//   go run ./02-language-basics/1-variables
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

func main() {
	var greeting string
	fmt.Printf("Initial zero value: '%s'\n", greeting)
	greeting = "Hello, world!"
	fmt.Println(greeting)

	var count int
	fmt.Printf("Initial zero value: %d\n", count)
	count = 10
	fmt.Println(count)

	var isActive bool
	fmt.Printf("Initial zero value: %t\n", isActive)
	isActive = true
	fmt.Println(isActive)

	firstName, lastName := "John", "Doe"
	fmt.Println(firstName, lastName)

	email := "test@test.com"
	fmt.Println(email)

	age := 24
	fmt.Println(age)

	var year = 2025
	fmt.Println(year)

	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: LB.2 constants")
	fmt.Println("Current: LB.1 (variables)")
	fmt.Println("---------------------------------------------------")
}
