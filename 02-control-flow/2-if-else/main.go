// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0
// Commercial use is prohibited without permission.

package main

import "fmt"

// ============================================================================
// Section 02: Control Flow — If/Else Statements
// Level: Beginner
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Basic if/else and if/else-if chains
//   - Go's unique "if with init statement" pattern
//   - The comma-ok idiom for map lookups
//   - Why Go doesn't need parentheses around conditions
//
// ENGINEERING DEPTH:
//   The `if init; condition` pattern keeps temporary values scoped to the branch
//   where they are needed. That makes Go code easier to read and reduces accidental
//   variable leakage into the outer function.
//
// RUN: go run ./02-control-flow/2-if-else
// ============================================================================

func main() {
	// --- BASIC IF/ELSE ---
	tmp := 25
	if tmp > 30 {
		fmt.Println("Temperature is above 30°C — hot!")
	} else {
		fmt.Println("Temperature is 30°C or below — comfortable")
	}

	// --- IF/ELSE-IF CHAIN ---
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A — Excellent")
	} else if score >= 80 {
		fmt.Println("Grade: B — Good")
	} else if score >= 70 {
		fmt.Println("Grade: C — Average")
	} else {
		fmt.Println("Grade: F — Failed")
	}

	// --- IF WITH INIT STATEMENT ---
	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	if hasAccess, ok := userAccess["john"]; ok && hasAccess {
		fmt.Println("John can access the system")
	} else if ok && !hasAccess {
		fmt.Println("John exists but access is denied")
	} else {
		fmt.Println("User not found")
	}

	if _, exists := userAccess["admin"]; !exists {
		fmt.Println("admin user not found in the system")
	}

	// KEY TAKEAWAY:
	// - No parentheses around conditions
	// - Braces are mandatory
	// - `if init; condition { }` keeps variables scoped tightly
	// - The comma-ok pattern (`value, ok := map[key]`) is used everywhere in Go
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("🚀 NEXT UP: CF.3 switch")
	fmt.Println("   Current: CF.2 (if / else)")
	fmt.Println("---------------------------------------------------")
}
