// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: If / Else
// Level: Foundation
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how a Go program chooses one path or another based on a condition.
//
// WHY THIS MATTERS:
//   - Branching is the ability to ask a question and choose a path. With `if`, `else if`, and `else`: - one condition is checked - one branch runs - the ...
//
// RUN:
//   go run ./02-language-basics/03-control-flow/1-if-else
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

func main() {
	temperature := 25

	if temperature > 30 {
		fmt.Println("Temperature is above 30C.")
	} else {
		fmt.Println("Temperature is 30C or below.")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	username := ""

	if username == "" {
		fmt.Println("Username is missing.")
	} else {
		fmt.Println("Username is present.")
	}

	fmt.Println("\n---------------------------------------------------")
	fmt.Println("NEXT UP: CF.2 for-basics")
	fmt.Println("Current: CF.1 (if / else)")
	fmt.Println("---------------------------------------------------")
}
