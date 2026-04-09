// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0
// Commercial use is prohibited without permission.

package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// Section 02: Control Flow — Pricing Calculator (Exercise)
// Level: Beginner
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Combining maps, loops, if/else, and functions
//   - Real-world data processing: looking up prices and applying discounts
//   - The comma-ok pattern for safe map access
//   - String manipulation with the strings package
//
// ENGINEERING DEPTH:
//   Notice how we strip suffixes using `strings.TrimSuffix()`. In Go, strings are
//   immutable read-only byte slices. When you "modify" a string, Go actually
//   allocates a new block of memory and copies the bytes over. For hot paths,
//   that cost is one reason byte slices (`[]byte`) are often preferred.
//
// RUN: go run ./02-control-flow/4-pricing-calculator
// ============================================================================

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.80
				fmt.Printf("  ✅ %s (Sale! $%.2f -> $%.2f)\n",
					originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}

		fmt.Printf("  ❌ %s (not found)\n", itemCode)
		return 0.0, false
	}

	fmt.Printf("  📦 %s — $%.2f\n", itemCode, basePrice)
	return basePrice, true
}

func main() {
	fmt.Println("=== Sales Order Processor ===")
	fmt.Println()

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK", "KEYBOARD",
	}

	var subtotal float64
	fmt.Println("Processing order:")

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Println()
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Println()
	fmt.Println("🎉 Section 02 complete! You're ready for Data Structures.")
	fmt.Println("   Next: Section 03 Data Structures")
}
