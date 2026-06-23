# iota

## Mission

Understand and apply iota in the context of professional Go software engineering.

## Prerequisites

- core-03-06

## Mental Model

Iota is a compile-time counter within const blocks. Like an auto-numbering machine - each line prints next number. Formulas (1 << iota) change what each number means.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

iota is a predeclared identifier representing untyped integer constant incrementing by 1 per line in const declaration. Compiler resets iota to 0 before each const block.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/07-iota
go test ./curriculum/modules/03-programming-fundamentals/lessons/07-iota
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming iota always starts at 0 - resets per const block, increments by 1 per line.
- Thinking iota works outside const declarations - only defined within const blocks.
- Confusing iota's line-based counting with semantic values - reordering lines changes iota values.

## In Production

Iota defines HTTP methods, DB isolation levels, permission flags, byte-size constants. Production Go relies on iota for compact enum definitions.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-08`.
