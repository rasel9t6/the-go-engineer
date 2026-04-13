// Copyright (c) 2026 Rasel Hossen
// See LICENSE for usage terms.

package main

import (
	"fmt"
	"strings"
)

func main() {
	cart := []string{"TSHIRT", "MUG_SALE", "HAT", "BOOK", "KEYBOARD"}

	var subtotal float64

	fmt.Println("Processing checkout:")

	for _, item := range cart {
		isSale := strings.HasSuffix(item, "_SALE")
		baseCode := strings.TrimSuffix(item, "_SALE")

		var price float64

		switch baseCode {
		case "TSHIRT":
			price = 20.00
		case "MUG":
			price = 12.50
		case "HAT":
			price = 18.00
		case "BOOK":
			price = 25.99
		}

		if price == 0 {
			fmt.Printf("skip %s: unknown item\n", item)
			continue
		}

		if isSale {
			originalPrice := price
			price = price * 0.80
			fmt.Printf("%s sale: %.2f -> %.2f\n", baseCode, originalPrice, price)
		} else {
			fmt.Printf("%s: %.2f\n", baseCode, price)
		}

		subtotal += price
	}

	fmt.Printf("subtotal: %.2f\n", subtotal)
}
