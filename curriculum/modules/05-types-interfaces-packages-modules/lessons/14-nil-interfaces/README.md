# Nil interfaces

## Mission

Understand and apply Nil interfaces in the context of professional Go software engineering.

## Prerequisites

- core-05-13

## Mental Model

An interface value is a box with two compartments: one holds the type tag, one holds the data. A nil interface has both compartments empty. An interface holding a nil pointer has the type compartment filled and the data compartment empty — the box is NOT empty because the type tag is present. The '== nil' check looks at the whole box, not just the data.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

An interface value in Go is implemented by the runtime as an iface (for non-empty interfaces) or eface (for empty interface{}). Both contain a _type pointer and an unsafe.Pointer for data. When you assign a nil *T to an interface, the _type pointer is set to the type descriptor for *T, and the data pointer is set to nil. The runtime's nil check (iface == nil) checks that BOTH the _type and data are nil — if either is non-nil, the interface is non-nil. This is by design: an interface that knows its type but has no data is not the same as an interface that doesn't know its type at all.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/14-nil-interfaces
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/14-nil-interfaces
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Writing return nil, err in a function that returns (MyCustomError, error) — the nil is typed as *MyCustomError(nil), making the error interface non-nil. The caller sees err != nil even though there was no error.
- Checking if err == nil after calling a function that returns (*MyError)(nil) — the condition is false because the interface has a type even though the pointer is nil.
- Using if err != nil { return err } in a loop where err is an interface — the first non-nil error is returned, but subsequent iterations may see a nil concrete value wrapped in a non-nil interface.
- Comparing two interface values that each contain a nil *T — they compare equal. Comparing a nil interface with an interface containing nil *T — they compare NOT equal.

## In Production

Every Go production service that defines custom error types (API errors, database errors, domain errors) must handle the typed-nil interface pattern. Bugs from this pattern have caused production outages in Kubernetes, Docker, and etcd. The fix is always: return nil as the interface type directly (return nil, nil) rather than returning a typed nil (return (*MyError)(nil), nil).

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-15`.
