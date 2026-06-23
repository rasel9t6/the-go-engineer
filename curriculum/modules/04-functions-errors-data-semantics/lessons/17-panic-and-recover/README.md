# panic and recover

## Mission

Understand and apply panic and recover in the context of professional Go software engineering.

## Prerequisites

- core-04-16

## Mental Model

Panic is an emergency stop, not a return value. When a function panics, the call stack unwinds, running deferred cleanups, until either a recover catches it or the program terminates. Recover is the emergency brake: it stops the unwinding and lets the goroutine continue.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The Go runtime maintains a per-goroutine panic stack. When panic is called, runtime.gopanic creates a _panic struct and begins unwinding. For each frame, it iterates the defer list (which is also per-goroutine) and executes each deferred function. If a deferred function calls runtime.gorecover, it marks the panic as recovered and stops unwinding. If the deferred function itself panics (a double panic), runtime.gopanic checks for the double-panic flag and, if set, exits with a fatal error and the second panic's stack trace. The net/http server leverages this by having each connection handler run in its own goroutine with a deferred recover that converts panics to HTTP 500 responses.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/17-panic-and-recover
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/17-panic-and-recover
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using panic for routine error handling instead of returning errors — panic should be for exceptional conditions, not regular error flow.
- Calling recover outside a deferred function — recover() returns nil if not called directly inside a defer.
- Assuming recover catches panics from other goroutines — recover only works in the goroutine that panicked.
- Panicking in a library function that the caller cannot recover from — library code should return errors, not panic.

## In Production

Production Go HTTP servers use recover in middleware to convert handler panics into 500 responses instead of crashing. CLI tools use recover to print a stack trace and exit gracefully. Batch processing pipelines use recover per-item to skip a single malformed record without aborting the entire batch.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-18`.
