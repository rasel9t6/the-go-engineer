// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0
// Commercial use is prohibited without permission.

package main

import (
	"fmt"
	"time"
)

// ============================================================================
// Section 02: Control Flow — Switch Statements
// Level: Beginner → Intermediate
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Basic switch with multiple case values
//   - Switch with init statement (like if with init)
//   - Tagless switch (switch without a value)
//   - Type switch for inspecting the dynamic type of an interface value
//   - Why Go's switch does not fall through by default
//
// ENGINEERING DEPTH:
//   In C and Java, forgetting `break` inside a switch can silently fall through to
//   the next case. Go flips that behavior: every case stops unless you explicitly
//   opt into `fallthrough`, which makes switch-heavy code safer by default.
//
// RUN: go run ./02-control-flow/3-switch
// ============================================================================

func main() {
	// --- BASIC SWITCH ---
	day := "Monday"
	fmt.Println("Today is", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend! No work")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings")
	default:
		fmt.Println("Mid-week")
	}

	fmt.Println()

	// --- SWITCH WITH INIT STATEMENT ---
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	fmt.Println()

	// --- TYPE SWITCH ---
	checkType := func(i any) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T with value: %v\n", v, v)
		}
	}

	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(312.23)

	// KEY TAKEAWAY:
	// - Go switch does not fall through by default
	// - Use comma-separated values for multiple matches
	// - Tagless switch works like an if/else-if chain
	// - Type switches help when data can hold different underlying types
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("🚀 NEXT UP: CF.4 pricing-calculator")
	fmt.Println("   Current: CF.3 (switch)")
	fmt.Println("---------------------------------------------------")
}
