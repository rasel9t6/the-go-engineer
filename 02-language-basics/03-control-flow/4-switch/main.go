// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: Switch
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn how to choose among several possible paths without building long, hard-to-scan branch chains.
//
// WHY THIS MATTERS:
//   - `switch` is a multi-way branch. It is useful when: - one value may match several known cases - several conditions need a clean top-to-bottom table ...
//
// RUN:
//   go run ./02-language-basics/03-control-flow/4-switch
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend mode.")
	case "Monday":
		fmt.Println("Start-of-week mode.")
	default:
		fmt.Println("Regular workday mode.")
	}

	score := 82

	switch {
	case score >= 90:
		fmt.Println("Excellent result.")
	case score >= 80:
		fmt.Println("Strong result.")
	case score >= 70:
		fmt.Println("Passing result.")
	default:
		fmt.Println("Needs more work.")
	}

	fmt.Println("\n---------------------------------------------------")
	fmt.Println("NEXT UP: CF.5 defer-basics")
	fmt.Println("Current: CF.4 (switch)")
	fmt.Println("---------------------------------------------------")
}
