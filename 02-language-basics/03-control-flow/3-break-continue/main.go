// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: Break / Continue
// Level: Foundation
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how to change a loop's behavior after the loop has already started.
//
// WHY THIS MATTERS:
//   - Loop control gives you two important tools: - `continue` skips the rest of the current iteration - `break` stops the loop completely That lets one ...
//
// RUN:
//   go run ./02-language-basics/03-control-flow/3-break-continue
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

func main() {
	fmt.Println("Odd numbers until the stop point:")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 7 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println("\n---------------------------------------------------")
	fmt.Println("NEXT UP: CF.4 switch")
	fmt.Println("Current: CF.3 (break / continue)")
	fmt.Println("---------------------------------------------------")
}
