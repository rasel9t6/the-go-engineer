// Copyright (c) 2026 Rasel Hossen
// See LICENSE for usage terms.

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
}
