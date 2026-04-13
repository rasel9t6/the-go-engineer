// Copyright (c) 2026 Rasel Hossen
// See LICENSE for usage terms.

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
}
