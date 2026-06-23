# Interfaces

## Mission

Understand and apply Interfaces in the context of professional Go software engineering.

## Prerequisites

- core-05-07

## Mental Model

An interface is a contract: 'anything with these methods satisfies me.' At runtime, an interface value is a pair: (type, pointer). The type tells Go which concrete method to call. The pointer holds the data (either a pointer to the value or the value itself boxed on the heap). Checking interface equality checks both the type and the data pointer.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's interface value is a runtime struct: iface for non-empty interfaces (has methods) and eface for empty interface{} (no methods). Both contain a type pointer and a data pointer. The type pointer points to an itable (interface table) that maps interface method offsets to concrete method addresses — computed once at the assignment point and cached. When a concrete value fits in one machine word (e.g., bool, int, pointer), it is stored directly in the data pointer slot without additional allocation. Larger values are heap-allocated and the data pointer points to the copy.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/08-interfaces
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/08-interfaces
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Declaring an interface and implementing it on a pointer receiver, then calling methods through a value type that does not satisfy the interface — the compiler error says 'does not implement' but the learner does not see why the method set is empty.
- Assigning a nil *T to an interface variable — the interface value is non-nil (it has a type and a nil pointer), so 'if err == nil' evaluates to false even though the underlying value is nil.
- Assuming io.Reader closes after EOF because Read returns (0, io.EOF) — the learner forgets that Read can return (n, io.EOF) with n > 0, requiring the caller to process n bytes before handling EOF.

## In Production

Every Go production service uses interfaces for HTTP handlers, database drivers, mock testing, middleware decoupling, and structured logging backends. Understanding interface nil semantics prevents a class of production bugs where error checks silently pass nil-as-non-nil.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-09`.
