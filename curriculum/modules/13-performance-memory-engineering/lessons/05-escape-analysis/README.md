# Escape analysis

## Mission

Understand and apply Escape analysis in the context of professional Go software engineering.

## Prerequisites

- core-13-04

## Mental Model

Escape analysis is the compiler's answer to: 'Does this value's address outlive the function that creates it?' If yes -> heap (GC must track it). If no -> stack (cheap allocation, freed when function returns). The compiler traces every pointer from creation to final use: if the pointer is returned, stored in a global, put into a heap-allocated struct, or captured by a closure — it escapes. If the pointer is only used locally or passed to functions that do not store it — it stays on the stack.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Escape analysis runs on the compiler's SSA form. It builds a directed graph of value flows: allocations, stores, loads, and function calls. For each allocation (call to runtime.newobject or stack allocation), the analysis walks forward through all uses of the allocated value. If any use stores the address in a location that outlives the allocating function (global, heap-allocated struct field, returned value, closure variable), the allocation is marked 'heap'. The analysis is interprocedural: it also analyzes callees to determine if they store the value in escaping locations. The compiler prints the reasoning with -gcflags=-m -m (two -m flags for more detail). Go's implementation is in cmd/compile/internal/escape.

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/05-escape-analysis
go test ./curriculum/modules/13-performance-memory-engineering/lessons/05-escape-analysis
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Optimizing allocation by guessing — changing &T{} to T{} or vice versa based on superstition rather than running -gcflags=-m or profiling.
- Assuming all function parameters are heap-allocated — small structs passed by value live on the stack; only values that escape (returned, stored in interface, captured by closure) go to the heap.
- Returning pointers to local variables unknowingly — the compiler detects this and allocates on heap, but the developer may think it is a stack allocation because the code 'looks local.'
- Using pointers everywhere to 'avoid copying' — pointers cause escape to heap more often than values, and copying small structs (up to a few words) is cheaper than heap allocation + GC.
- Believing that 'make' always allocates on heap — make for small slices created in a function and returned may allocate; but slices that never escape stay on stack if their backing array is small enough.

## In Production

Escape analysis optimization is critical for high-throughput Go services. The standard library's net/http server uses object pools (sync.Pool) for buffer reuse precisely because escape analysis shows that response buffers would otherwise escape to the heap. Uber's goleak detection and Datadog's trace-agent both optimize hot paths to avoid escape. In production profiling, the first question about unexpected GC overhead is: 'what is escaping to the heap?'

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-13-06`.
