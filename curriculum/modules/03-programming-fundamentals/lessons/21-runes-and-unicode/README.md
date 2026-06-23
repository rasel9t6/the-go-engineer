# Runes and Unicode

## Mission

Understand and apply Runes and Unicode in the context of professional Go software engineering.

## Prerequisites

- core-03-20

## Mental Model

A rune is a Unicode code point - 32-bit integer representing one 'character'. Bytes = ink on page, runes = letters they form. Some letters use 1 byte (ASCII), others 2-4 bytes.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Rune is Go's int32 representing Unicode code point. String is byte sequence (not necessarily UTF-8). But for range, []rune, utf8 package assume UTF-8. Invalid UTF-8 produces replacement rune U+FFFD.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/21-runes-and-unicode
go test ./curriculum/modules/03-programming-fundamentals/lessons/21-runes-and-unicode
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing 'character' in Go with single byte - rune can be 1-4 bytes in UTF-8.
- Using string indexing for characters - s[i] gives byte, use for range for runes.
- Assuming len(string) gives characters - it gives UTF-8 bytes.

## In Production

Internationalization, emoji handling (1-4 bytes), URL encoding, text processing depend on rune-aware operations.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-22`.
