# go generate

## Mission

Understand and apply go generate in the context of professional Go software engineering.

## Prerequisites

- elective-06

## Mental Model

go generate is a project-wide command runner. A //go:generate directive in a Go source file tells go generate to run a specific command. Running go generate ./... finds all directives in the project and executes them in order. Each directive is associated with a Go source file, making it clear which generators produce which files. The generated files are checked into version control so consumers do not need to run go generate.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

go generate scans .go files for lines matching //go:generate followed by a command and arguments. It executes the command in the file's directory. The command's stdout and stderr are printed to the terminal. go generate does not parse Go syntax — it only scans for the //go:generate prefix. Generated files typically have a _gen.go suffix or are written to a specified output directory.

## Run Instructions

```bash
Read the lesson and complete the practice task.
No automated test is required for this lesson.
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Running go generate from the wrong directory — commands are executed relative to the file's directory, not the project root.
- Forgetting to check generated files into version control — consumers (CI, teammates) cannot build without running go generate first.
- Not making go generate idempotent — running go generate twice should produce the same output and not fail on the second run.
- Putting complex logic in the go:generate line — the directive should be a simple command invocation; complex logic belongs in a script.
- Not documenting which go:generate directives to run — some generates depend on others; document the execution order in a Makefile or README.

## In Production

go generate is used in every Go project that uses code generation: protobuf (protoc --go_out), OpenAPI (oapi-codegen), mocking (mockgen, counterfeiter), stringer (stringer), and database (sqlc, ent). Production Go projects use go generate as a standard build step documented in CONTRIBUTING.md. CI runs go generate ./... and then git diff --exit-code to ensure generated files are up to date.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-08`.
