// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: For Basics
// Level: Foundation
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how Go repeats work with its single loop keyword: `for`.
//
// WHY THIS MATTERS:
//   - A loop says, "keep doing this work while the rule allows it." Go uses one keyword for several loop shapes: - counted loops - condition-only loops -...
//
// RUN:
//   go run ./02-language-basics/03-control-flow/2-for-basics
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

func main() {
	fmt.Println("Counted loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("step %d\n", i)
	}

	fmt.Println()
	fmt.Println("Condition-only loop:")
	countdown := 3
	for countdown > 0 {
		fmt.Printf("countdown: %d\n", countdown)
		countdown--
	}

	fmt.Println()
	fmt.Println("Range preview:")
	words := []string{"go", "learn", "repeat"}
	for _, word := range words {
		fmt.Printf("word = %s\n", word)
	}

	fmt.Println("\n---------------------------------------------------")
	fmt.Println("NEXT UP: CF.3 break-continue")
	fmt.Println("Current: CF.2 (for basics)")
	fmt.Println("---------------------------------------------------")
}
