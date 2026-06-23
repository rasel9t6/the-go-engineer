# Strings

## Mission

Understand and apply Strings in the context of professional Go software engineering.

## Prerequisites

- core-03-18

## Mental Model

A string is a read-only view of a byte sequence. Sealed envelope: peek inside (read bytes), but cannot change message. Copy to whiteboard ([]byte) to edit.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

String is struct{ ptr *byte, len int } - 16 bytes. Pointer references immutable byte array. Multiple strings can share same array (slicing). Compiler ensures immutability.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/19-strings
go test ./curriculum/modules/03-programming-fundamentals/lessons/19-strings
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Treating strings as mutable - strings are immutable byte sequences.
- Using len(string) to count characters - len returns bytes, not runes.
- Assuming string indexing yields a character - s[i] returns byte, not rune.

## In Production

Every HTTP header, JSON key, file path, log line is a string. String processing is most common non-trivial Go task.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-20`.
