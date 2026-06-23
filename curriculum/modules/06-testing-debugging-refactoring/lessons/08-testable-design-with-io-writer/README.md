# Testable design with io.Writer

## Mission

Understand and apply Testable design with io.Writer in the context of professional Go software engineering.

## Prerequisites

- core-06-07

## Mental Model

io.Writer is like a universal output socket. Your function writes data to this socket. In production, the socket is connected to stdout or a file. In tests, it's connected to a buffer you can inspect. The function doesn't know or care where the output goes.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

io.Writer is a single-method interface: `Write(p []byte) (n int, err error)`. bytes.Buffer implements it by appending to its internal byte slice. os.Stdout implements it by writing to the OS file descriptor.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/08-testable-design-with-io-writer
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/08-testable-design-with-io-writer
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Hardcoding fmt.Print in functions instead of accepting io.Writer.
- Writing to stdout directly in business logic -- impossible to test output.
- Creating test-specific output paths instead of using io.Writer abstraction.
- Passing strings around instead of io.Writer for output destinations.

## In Production

CLI tools accept io.Writer for output. HTTP handlers write to ResponseWriter (which implements io.Writer). Logging libraries accept io.Writer as output targets. It's the fundamental Go output abstraction.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-09`.
