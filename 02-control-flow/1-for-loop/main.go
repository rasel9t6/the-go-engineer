// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0
// Commercial use is prohibited without permission.

package main

import "fmt"

// ============================================================================
// Section 02: Control Flow — For Loops
// Level: Beginner
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Go has only one loop keyword: "for" (no while, no do-while)
//   - The four common forms of the for loop: C-style, while-style, infinite, range
//   - break and continue flow control
//   - The range keyword for iterating over collections
//
// ENGINEERING DEPTH:
//   Go deliberately removed separate `while` and `do-while` keywords. A while loop
//   is just a `for` loop with the init and post statements omitted, so one keyword
//   keeps the language smaller and the code more visually consistent.
//
// RUN: go run ./02-control-flow/1-for-loop
// ============================================================================

func main() {
	// --- FORM 1: C-style for loop ---
	fmt.Println("=== C-Style Loop ===")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  i = %d\n", i)
	}

	// --- FORM 2: While-style loop ---
	fmt.Println("\n=== While-Style Loop ===")
	k := 3
	for k > 0 {
		fmt.Printf("  k = %d\n", k)
		k--
	}

	// --- FORM 3: Infinite loop ---
	fmt.Println("\n=== Infinite Loop (with break) ===")
	counter := 0
	for {
		fmt.Printf("  counter = %d\n", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	// --- CONTINUE: Skip the current iteration ---
	fmt.Println("\n=== Continue (odd numbers only) ===")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("  %d\n", i)
	}

	// --- FORM 4: Range loop ---
	fmt.Println("\n=== Range Loop ===")
	items := [3]string{"Go", "Python", "Java"}
	for index, value := range items {
		fmt.Printf("  items[%d] = %s\n", index, value)
	}

	fmt.Println("\n  Values only:")
	for _, value := range items {
		fmt.Printf("  %s\n", value)
	}

	// KEY TAKEAWAY:
	// Go has one loop keyword (`for`) with four common forms:
	//   1. for i := 0; i < n; i++ { }     — C-style
	//   2. for condition { }              — while-style
	//   3. for { }                        — infinite loop
	//   4. for index, value := range x { } — iterate collections
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("🚀 NEXT UP: CF.2 if / else")
	fmt.Println("   Current: CF.1 (for loop)")
	fmt.Println("---------------------------------------------------")
}
