// Copyright (c) 2026 Rasel Hossen
// Licensed under The Go Engineer License v1.0

// ============================================================================
// Section 02: Language Basics
// Title: Application Logger
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Build a small logger that combines variables, constants, `iota`, and methods into one readable program.
//
// WHY THIS MATTERS:
//   - This exercise turns separate language pieces into one compact system: - a named type models the log level - `iota` creates ordered constants - a me...
//
// RUN:
//   go run ./02-language-basics/4-application-logger
//
// KEY TAKEAWAY:
//   - [TODO: Summarize the core takeaway]
// ============================================================================

package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}
	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)

	printLogLevel(10)

	fmt.Println()
	fmt.Println("Section 02 complete! Ready for Control Flow.")
}
