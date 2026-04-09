# LB.4 Application Logger

## Mission

Build a small application logger that turns early language fundamentals into one readable,
runnable program.

This is the Section 01 milestone.
It is where variables, constants, `iota`, named types, formatting, and basic defensive checks come
together.

## Prerequisites

Complete these first:

- `GT.4` development environment
- `LB.1` variables
- `LB.2` constants
- `LB.3` enums with `iota`

## What You Will Build

Implement a small logger that:

1. defines a `LogLevel` type based on `int`
2. creates level constants with `iota`
3. maps each level to a readable name
4. implements a `String()` method with bounds checking
5. prints readable log level output through a helper function

## Files

- [main.go](./main.go): complete solution with teaching comments
- [_starter/main.go](./_starter/main.go): starter file with TODOs and requirements

## Run Instructions

Run the completed solution:

```bash
go run ./01-core-foundations/language-basics/4-application-logger
```

Run the starter:

```bash
go run ./01-core-foundations/language-basics/4-application-logger/_starter
```

## Success Criteria

Your finished solution should:

- define a named `LogLevel` type
- use `iota` for ordered log level constants
- convert levels into readable text safely
- avoid crashing on invalid levels
- print a few example log levels in `main()`

## Common Failure Modes

- using raw integers everywhere instead of a named type
- forgetting the bounds check before indexing the `levelNames` slice
- assuming Go has a built-in enum keyword
- building too much structure for what should stay a simple first milestone

## Next Step

After this milestone, continue to Section 02: Control Flow.
