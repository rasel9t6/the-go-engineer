// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 06: Backend, APIs & Databases
// Title: Pagination and filtering
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Learn why large result sets need stable pagination and explicit filtering rules.
//
// WHY THIS MATTERS:
//   - Pagination is resource budgeting. It prevents one query from turning one endpoint into unbounded work.
//
// RUN:
//   go run ./06-backend-db/01-web-and-database/apis/3-pagination-and-filtering
//
// KEY TAKEAWAY:
//   - Learn why large result sets need stable pagination and explicit filtering rules.
// ============================================================================

package main

import "fmt"

//

func main() {
	fmt.Println("=== API.3 Pagination and filtering ===")
	fmt.Println("Learn why large result sets need stable pagination and explicit filtering rules.")
	fmt.Println()
	fmt.Println("- Always define an ordering when paginating.")
	fmt.Println("- Offset pagination is simple but weak on large, changing datasets.")
	fmt.Println("- Cursor pagination is harder up front but often safer at scale.")
	fmt.Println()
	fmt.Println("Pagination bugs often become latency bugs, memory bugs, or duplicated work for clients.")
	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println("NEXT UP: API.4")
	fmt.Println("Current: API.3 (pagination and filtering)")
	fmt.Println("---------------------------------------------------")
}
