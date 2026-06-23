# Receiver sets

## Mission

Understand and apply Receiver sets in the context of professional Go software engineering.

## Prerequisites

- core-05-04

## Mental Model

Every type T has a method set containing all methods declared with a T receiver. Every pointer type *T has a method set containing all methods declared with either a T receiver or a *T receiver. Interface satisfaction checks the method set of the concrete type: if the variable is T, only T's method set counts; if the variable is *T, *T's method set (which includes T's methods) counts. Addressability determines whether you can call a pointer receiver method on a value — you can if and only if the value's address can be taken.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The Go compiler maintains two method tables per type: one for T (methods with T receivers) and one for *T (methods with T + *T receivers). Interface satisfaction checks whether the method set of the concrete type (T or *T) contains all methods of the interface. This is resolved at compile time: the compiler generates an interface dispatch table (itab) for each concrete type-interface pair, stored in the binary. At runtime, the itab pointer is stored alongside the value in the interface struct (interface{typecode, value}). The method call is a pointer lookup through the itab — one indirect call, no reflection, no type switching.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/05-receiver-sets
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/05-receiver-sets
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming a value type automatically satisfies an interface that requires pointer receiver methods — T has access to T's method set (value receivers only), *T has access to *T's method set (all methods including pointer receivers). A value can only call pointer receiver methods if it is addressable (e.g., a map element is not addressable).
- Passing a value to a function that expects an interface with pointer receiver methods — the value does not satisfy the interface; the compiler rejects it with 'does not implement'.
- Not understanding that a slice of T must be converted to a slice of the interface type manually — even if T satisfies io.Writer, []T does not automatically satisfy []io.Writer.
- Assuming that a value receiver method on *T also appears in T's method set — only *T's method set includes both value and pointer receiver methods; T's method set includes only value receiver methods.
- Calling pointer receiver methods on non-addressable values (results of function calls, map index expressions) — the compiler rejects them because the value is not addressable and cannot be implicitly referenced.

## In Production

Every Go program that uses interfaces encounters method set rules. When passing http.Handler implementations, the handler type must have ServeHTTP in its method set. When implementing database/sql Scanner, the Scan method must use a pointer receiver because it modifies the receiver. When implementing sort.Interface on a slice type, value receivers are idiomatic. When implementing json.Marshaler, value receivers are common because serialization does not mutate the receiver.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-06`.
