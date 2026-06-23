# Pointer and value mutation behavior

## Mission

Understand and apply Pointer and value mutation behavior in the context of professional Go software engineering.

## Prerequisites

- core-04-07

## Mental Model

Go is pass-by-value, always. Every function parameter is a copy of the caller's argument. If the argument is a 100-byte struct, 100 bytes are copied onto the callee's stack. If the argument is a pointer (always 8 bytes on 64-bit), the pointer is copied — but both the original and the copy point to the same memory. Mutation through a copied pointer modifies the original memory. This is not pass-by-reference — it is pass-by-value where the value happens to be an address.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's calling convention (Go 1.17+) passes function arguments and return values in registers (up to a certain size), then on the stack for larger structs. For a struct passed by value, the compiler emits a memcpy of the struct's bytes from the caller's frame to the callee's frame (or registers). For a pointer, only the 8-byte address is copied. Maps and channels are implemented as pointers to runtime descriptors (hmap for maps, hchan for channels) — when you pass a map to a function, you copy the pointer, not the entire descriptor. Slices are structs {ptr, len, cap} — the header is copied, but the ptr field points to the same backing array. This is why map modifications and slice element changes inside a function are visible to the caller, but slice append (which may change len/cap) is not.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/08-pointer-and-value-mutation-behavior
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/08-pointer-and-value-mutation-behavior
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Thinking a function that receives a struct value can modify the caller's original — Go copies the entire struct, so modifications inside the function are lost when the function returns.
- Using pointers everywhere to avoid copying without considering mutation semantics — passing a pointer allows the callee to modify the original, which may be unintended and cause hard-to-find bugs.
- Mixing pointer and value receivers on the same type — some methods mutate (pointer receiver) and some do not (value receiver), making the API inconsistent and surprising callers.
- Taking the address of a loop variable and storing it — the address is the same for every iteration, and by the time the stored pointer is used, the loop variable holds its final value.
- Assigning a pointer to an interface variable — the interface holds a copy of the pointer, not a copy of the value, so mutations through the interface mutate the original.

## In Production

Every Go codebase depends on understanding value vs pointer mutation semantics. Misunderstanding causes data races (multiple goroutines holding pointers to the same struct), silent data corruption (modifying a slice backing array shared with another caller), and logic bugs (trying to append to a slice inside a function and expecting the caller's length to update).

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-09`.
