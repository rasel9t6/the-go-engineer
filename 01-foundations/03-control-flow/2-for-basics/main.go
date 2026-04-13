// Copyright (c) 2026 Rasel Hossen
// See LICENSE for usage terms.

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
	for index, word := range words {
		fmt.Printf("words[%d] = %s\n", index, word)
	}
}
