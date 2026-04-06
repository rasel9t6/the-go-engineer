// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0
// Commercial use is prohibited without permission.

package main

import (
	"context"
	"fmt"
	"time"
)

// ============================================================================
// Section 11: Context Ã¢â‚¬â€ WithCancel
// Level: Intermediate
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - context.WithCancel creates a cancellable context
//   - The cancel function: how and when to call it
//   - Listening for cancellation with ctx.Done()
//   - Using select to handle cancellation in goroutines
//   - Why you MUST always call cancel() (preventing goroutine leaks)
//
// ENGINEERING DEPTH:
//   When you invoke `context.WithCancel(parent)`, Go creates a `cancelCtx` struct
//   under the hood. This struct contains a Mutex and a `done` channel. It also
//   recursively climbs up the Context tree until it finds the first cancellable
//   parent, and appends itself to that parent's internal array of children!
//   When you call `cancel()`, Go closes the `done` channel, broadcasts the close
//   signal to all children in the array, and then removes itself from the parent
//   to allow the Garbage Collector to sweep the dead goroutines.
//
// RUN: go run ./11-concurrency/context/2-with-cancel
// ============================================================================

func main() {
	fmt.Println("=== Context: WithCancel ===")
	fmt.Println()

	// --- CREATING A CANCELLABLE CONTEXT ---
	// WithCancel returns two values:
	//   1. ctx Ã¢â‚¬â€ a new context that can be cancelled
	//   2. cancel Ã¢â‚¬â€ a function that triggers the cancellation
	//
	// CRITICAL RULE: You MUST call cancel() when you're done with the context.
	// If you don't, the context's resources (goroutines) are never freed.
	// Use defer cancel() immediately after creation.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Always defer cancel Ã¢â‚¬â€ even if you call it explicitly later

	// Start a background worker that listens for cancellation
	results := make(chan string)
	go worker(ctx, results)

	// Collect some results from the worker
	for i := 0; i < 3; i++ {
		fmt.Printf("  Received: %s\n", <-results)
	}

	// Now cancel the context Ã¢â‚¬â€ this signals the worker to stop
	fmt.Println("\n  Calling cancel()...")
	cancel()

	// Give the worker a moment to notice the cancellation
	time.Sleep(100 * time.Millisecond)

	// After cancellation, ctx.Err() returns a non-nil error
	fmt.Printf("  Context error after cancel: %v\n", ctx.Err())
	// ctx.Err() returns context.Canceled

	fmt.Println()

	// --- NESTED CANCELLATION ---
	// When a parent context is cancelled, ALL children are cancelled too.
	fmt.Println("=== Nested Cancellation ===")
	parentCtx, parentCancel := context.WithCancel(context.Background())

	// Create child contexts from the parent
	childCtx1, childCancel1 := context.WithCancel(parentCtx)
	childCtx2, childCancel2 := context.WithCancel(parentCtx)
	defer childCancel1()
	defer childCancel2()

	fmt.Printf("  Before cancel Ã¢â‚¬â€ parent err: %v, child1 err: %v, child2 err: %v\n",
		parentCtx.Err(), childCtx1.Err(), childCtx2.Err())

	// Cancel the PARENT Ã¢â‚¬â€ both children are cancelled automatically
	parentCancel()

	fmt.Printf("  After parent cancel Ã¢â‚¬â€ parent err: %v, child1 err: %v, child2 err: %v\n",
		parentCtx.Err(), childCtx1.Err(), childCtx2.Err())

	fmt.Println()
	fmt.Println("KEY TAKEAWAYS:")
	fmt.Println("  1. WithCancel returns (ctx, cancel) Ã¢â‚¬â€ ALWAYS defer cancel()")
	fmt.Println("  2. Listen for cancellation with <-ctx.Done() in a select")
	fmt.Println("  3. After cancellation, ctx.Err() returns context.Canceled")
	fmt.Println("  4. Cancelling a parent cancels ALL children automatically")
	fmt.Println("  5. Forgetting cancel() causes GOROUTINE LEAKS (memory leak)")
	fmt.Println()
	fmt.Println("   Next: go run ./11-concurrency/context/3-with-timeout")
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("Ã°Å¸Å¡â‚¬ NEXT UP: CT.3 WithTimeout")
	fmt.Println("   Current: CT.2 (WithCancel)")
	fmt.Println("---------------------------------------------------")
}

// worker simulates a long-running task that checks for cancellation.
// This is the standard pattern for cancellation-aware goroutines.
//
// The select statement waits for EITHER:
//   - ctx.Done() Ã¢â‚¬â€ the context was cancelled (stop working)
//   - default/time Ã¢â‚¬â€ normal work continues
func worker(ctx context.Context, results chan<- string) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			// The context was cancelled Ã¢â‚¬â€ clean up and return
			fmt.Printf("  Worker stopped: %v\n", ctx.Err())
			return
		default:
			// Context is still active Ã¢â‚¬â€ do work
			i++
			results <- fmt.Sprintf("result-%d", i)
			time.Sleep(50 * time.Millisecond) // Simulate work
		}
	}
}
