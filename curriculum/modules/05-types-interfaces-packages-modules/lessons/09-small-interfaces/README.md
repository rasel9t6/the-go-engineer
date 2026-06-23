# Small interfaces

## Mission

Understand and apply Small interfaces in the context of professional Go software engineering.

## Prerequisites

- core-05-08

## Mental Model

A small interface is like a power outlet -- it defines one specific capability. Any device providing that capability can plug in. The interface doesn't care how -- only that the capability is provided.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Interface values are two-word structures: type pointer (itab) and data pointer. Method calls look up the function in the itab's function table. Single indirection for dispatch.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/09-small-interfaces
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/09-small-interfaces
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Defining interfaces with 10+ methods -- they should be 1-3 methods at most.
- Creating an interface before it has two concrete implementations -- premature abstraction.
- Putting setters and getters in an interface -- use exported fields or methods that do something.
- Making interfaces that mirror a concrete type's methods exactly -- that's indirection, not abstraction.

## In Production

Production Go defines interfaces like `Storer { Store(key string, data []byte) error }`, `Notifier { Notify(msg Message) error }`, `Cache { Get(key string) ([]byte, bool) }`. Each is easily mocked.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-10`.
