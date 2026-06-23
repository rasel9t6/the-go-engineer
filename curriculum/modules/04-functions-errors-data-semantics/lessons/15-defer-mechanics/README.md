# defer mechanics

## Mission

Understand and apply defer mechanics in the context of professional Go software engineering.

## Prerequisites

- core-04-14

## Mental Model

Defer is a stack of cleanup functions attached to the current function call. When the function returns (normally or via panic), the stack is popped LIFO. Each deferred call's arguments are evaluated at the defer site, but the function body executes at return time.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's runtime maintains a per-goroutine linked list of _defer structs. When the compiler encounters a defer statement, it calls runtime.deferproc to allocate a _defer node, populate it with the function pointer and arguments, and prepend it to the goroutine's defer list. When the function returns, the epilogue calls runtime.deferreturn, which pops _defer nodes from the list and executes them. In Go 1.14+, non-defer-panic paths are optimized so that deferred calls are inlined at each return site rather than going through the runtime — this makes the common case much faster.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/15-defer-mechanics
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/15-defer-mechanics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming deferred calls run immediately when the function exits — they run when the surrounding function returns, not when the block ends.
- Capturing loop variables in a defer inside a loop — the defer captures the variable's address, and by the time the defer runs, the loop variable has advanced to its final value.
- Using defer in a tight loop — each deferred call allocates a stack entry that is not freed until the function returns, causing O(n) memory growth.
- Assuming deferred recover catches panics in a different goroutine — recover only works in the same goroutine as the panic.

## In Production

Every Go HTTP handler defers response body close. Every mutex lock is paired with defer unlock. Every database transaction defers rollback on error and commit on success. Defer is the standard Go pattern for resource lifecycle management.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-16`.
