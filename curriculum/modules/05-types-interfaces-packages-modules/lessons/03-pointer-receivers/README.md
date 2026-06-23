# Pointer receivers

## Mission

Understand and apply Pointer receivers in the context of professional Go software engineering.

## Prerequisites

- core-05-02

## Mental Model

A pointer receiver is like giving someone the key to your house instead of a photo. With the key (pointer), they can rearrange furniture (modify fields). With a photo (value), they can only describe what they see.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A pointer receiver method receives a pointer (8 bytes on 64-bit). Go auto-converts `value.Method()` to `(&value).Method()` when addressable. This works for variables, struct fields, array elements, but not map elements or function returns.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/03-pointer-receivers
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/03-pointer-receivers
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using a pointer receiver when the method doesn't modify the receiver -- be consistent but don't over-use pointers.
- Forgetting that pointer receiver methods are not in the method set of a non-pointer type variable.
- Calling a pointer receiver method on a non-addressable value (e.g., function return value).
- Modifying the pointer receiver itself inside the method -- only changes the local pointer.

## In Production

Every Go type needing mutation uses pointer receivers: `*http.Request`, `*os.File`, `*bytes.Buffer`, `*sql.DB`, `*net.Conn`. Factory functions return pointers. Pointer receivers are standard for stateful types.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-04`.
