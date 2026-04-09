// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0
// Commercial use is prohibited without permission.

package main

import (
	"fmt"
	"slices"
)

// ============================================================================
// Section 03: Data Structures — Advanced Slicing
// Level: Intermediate
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Advanced slice expressions and capacity behavior
//   - How sub-slices share memory (and when they stop sharing)
//   - How "slicing a slice" can isolate or accidentally couple data
//   - A first look at the Go 1.21 `slices` package
//
// ENGINEERING DEPTH:
//   When you sub-slice (for example `s[2:5]`), Go usually does not allocate
//   new memory. It creates a new slice header that points into the same backing
//   array. That makes slicing cheap, but it also means tiny sub-slices can keep
//   much larger arrays alive and can sometimes mutate shared data unexpectedly.
//
// RUN: go run ./03-data-structures/5-slices-2
// ============================================================================

func main() {
	fmt.Println("=== Advanced Slicing Operations ===")

	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v, len: %d, cap: %d\n", original, len(original), cap(original))

	// --- SLICING SYNTAX EXPLAINED ---
	s1 := original[2:5]
	fmt.Printf("\n  s1 (original[2:5]): %v\n", s1)
	fmt.Printf("  s1 len: %d, cap: %d\n", len(s1), cap(s1))

	s2 := original[:4]
	fmt.Printf("\n  s2 (original[:4]): %v\n", s2)

	s3 := original[6:]
	fmt.Printf("  s3 (original[6:]): %v\n", s3)

	s4 := original[:]
	fmt.Printf("  s4 (original[:]):  %v\n", s4)

	// --- THE SLICES STANDARD PACKAGE ---
	found := slices.Contains(s4, 8)
	fmt.Printf("\n  s4 contains 8? %v\n", found)

	// --- THE APPEND SHARED-MEMORY TRAP ---
	s4 = append(s4, 1000)

	fmt.Println("\n=== The Re-Allocation Trap ===")
	fmt.Printf("s4 (after append): %v\n", s4)
	fmt.Printf("Original:          %v\n", original)

	// KEY TAKEAWAY:
	// - Sub-slicing is fast and cheap, but shares memory.
	// - The capacity of a sub-slice extends to the end of the backing array.
	// - Appending to a sub-slice can overwrite data in the original slice if capacity allows it.
	// - If capacity is exceeded, append allocates new memory and breaks the link.
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("🚀 NEXT UP: DS.6 contact-manager")
	fmt.Println("   Current: DS.5 (slices-2)")
	fmt.Println("---------------------------------------------------")
}
