// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0
// Commercial use is prohibited without permission.

package main

import (
	"fmt"
	"time"
)

// ============================================================================
// Section 11: Concurrency â€” Buffered Channels
// Level: Intermediate
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Buffered vs unbuffered channels â€” the critical distinction
//   - make(chan T, capacity) â€” creating a channel with buffer space
//   - When buffered channels block (only when FULL or EMPTY)
//   - Use cases: batch processing, rate limiting, producer-consumer
//   - When to use buffered vs unbuffered
//
// ANALOGY:
//   Unbuffered = a phone call. Both parties must be on the line at the
//   same time. The sender WAITS until the receiver picks up.
//
//   Buffered = a mailbox with N slots. The sender can drop messages
//   and keep going â€” until the mailbox is full. Then the sender waits.
//   The receiver can pick up messages whenever ready.
//
// ENGINEERING DEPTH:
//   Buffered channels use a circular ring buffer internally.
//   - Send blocks only when the buffer is FULL
//   - Receive blocks only when the buffer is EMPTY
//   - Buffer size is set at creation and CANNOT be resized
//   - cap(ch) returns the buffer capacity
//   - len(ch) returns the number of items currently in the buffer
//
// RUN: go run ./11-concurrency/goroutines/4-channels-buffered
// ============================================================================

// logEvent represents a system event to be processed asynchronously.
type logEvent struct {
	Level   string // "INFO", "WARN", "ERROR"
	Message string // Event description
}

func main() {
	fmt.Println("=== Buffered Channels ===")
	fmt.Println()

	// =====================================================================
	// 1. Basic Buffered Channel
	// =====================================================================
	// make(chan T, N) creates a buffered channel with capacity N.
	// The sender can put up to N values WITHOUT a receiver being ready.
	events := make(chan logEvent, 3) // Buffer holds up to 3 events

	// Because the buffer has space, these sends DON'T block.
	// With an unbuffered channel, these would deadlock (no receiver).
	events <- logEvent{"INFO", "Server started on :8080"}
	events <- logEvent{"INFO", "Connected to database"}
	events <- logEvent{"WARN", "Cache miss rate above 50%"}
	// events <- logEvent{"ERROR", "timeout"} â† This 4th send would BLOCK (buffer full!)

	fmt.Printf("  Buffer: %d/%d items\n\n", len(events), cap(events))

	// Receive all events
	fmt.Println("  1ï¸âƒ£  Basic Buffered Channel (capacity=3):")
	for i := 0; i < 3; i++ {
		e := <-events // Each receive removes one item from the buffer
		fmt.Printf("     [%s] %s\n", e.Level, e.Message)
	}
	fmt.Println()

	// =====================================================================
	// 2. Producer-Consumer Pattern
	// =====================================================================
	// The most common use of buffered channels: decouple a fast producer
	// from a slower consumer. The buffer absorbs bursts.
	fmt.Println("  2ï¸âƒ£  Producer-Consumer Pattern:")

	jobs := make(chan int, 5) // Buffer holds 5 jobs

	// Producer: generates jobs fast
	go func() {
		for i := 1; i <= 8; i++ {
			fmt.Printf("     ðŸ“¤ Producing job #%d\n", i)
			jobs <- i // Blocks only when buffer is full
		}
		close(jobs) // Signal: no more jobs coming
	}()

	// Consumer: processes jobs slowly
	// range over a channel reads until the channel is CLOSED.
	time.Sleep(50 * time.Millisecond) // Let producer fill buffer first
	for job := range jobs {
		fmt.Printf("     ðŸ“¥ Processing job #%d\n", job)
		time.Sleep(30 * time.Millisecond) // Simulate slow processing
	}
	fmt.Println()

	// =====================================================================
	// 3. Comparison: Buffered vs Unbuffered
	// =====================================================================
	fmt.Println("  3ï¸âƒ£  Buffered vs Unbuffered:")
	fmt.Println("     â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¬â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”")
	fmt.Println("     â”‚   Unbuffered    â”‚         Buffered               â”‚")
	fmt.Println("     â”œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¤")
	fmt.Println("     â”‚ make(chan T)     â”‚ make(chan T, N)                â”‚")
	fmt.Println("     â”‚ Send blocks     â”‚ Send blocks only when FULL     â”‚")
	fmt.Println("     â”‚ until received  â”‚ Receive blocks only when EMPTY â”‚")
	fmt.Println("     â”‚ Synchronization â”‚ Async with bounded queue       â”‚")
	fmt.Println("     â”‚ Phone call      â”‚ Mailbox with N slots           â”‚")
	fmt.Println("     â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”´â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜")

	fmt.Println()
	fmt.Println("KEY TAKEAWAY:")
	fmt.Println("  - Buffered: make(chan T, N) â€” N items can be sent without blocking")
	fmt.Println("  - Unbuffered: make(chan T) â€” sender waits for receiver (synchronous)")
	fmt.Println("  - Use buffered channels to decouple fast producers from slow consumers")
	fmt.Println("  - Buffer size should be tuned based on throughput needs")
	fmt.Println("  - When in doubt, start unbuffered â€” add buffer only for performance")
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("ðŸš€ NEXT UP: GC.5 closing channels")
	fmt.Println("   Current: GC.4 (buffered channels)")
	fmt.Println("---------------------------------------------------")
}
