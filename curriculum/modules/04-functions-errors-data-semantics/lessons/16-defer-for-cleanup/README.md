# defer for cleanup

## Mission

Understand and apply defer for cleanup in the context of professional Go software engineering.

## Prerequisites

- core-04-15

## Mental Model

Defer is like writing a 'to-do on exit' list. When you acquire a resource, you immediately add a cleanup task. As you acquire more, new tasks go to the top. When the function exits, Go reads the list from top to bottom and does each task.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The compiler registers a deferred function in a linked list on the goroutine. Each `defer` adds a `_defer` struct containing the function pointer and arguments. On return or panic, the runtime walks this chain in LIFO order.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/16-defer-for-cleanup
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/16-defer-for-cleanup
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming deferred functions run after the function returns -- they run when the enclosing function reaches its end, including on panic.
- Putting `defer` after the resource acquisition -- if the acquire fails, `defer` on a nil resource causes panic.
- Deferring a function that modifies a loop variable -- all defer calls see the final value of the variable.
- Using `defer` in a tight loop -- each defer incurs overhead and runs at function return, not loop iteration end.

## In Production

Every file operation, mutex lock, HTTP response body read, and database connection in production Go uses defer. `sync.Mutex` is always paired: `mu.Lock(); defer mu.Unlock()`.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-17`.
