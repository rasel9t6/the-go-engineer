# go.sum

## Mission

Understand and apply go.sum in the context of professional Go software engineering.

## Prerequisites

- core-05-19

## Mental Model

go.sum is like a fingerprint database for your dependencies. Each module version has a unique fingerprint (checksum). Before using a dependency, Go checks that its fingerprint matches the recorded one. This prevents accidentally or maliciously modified dependencies.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

go.sum contains lines like `module version h1:hash`. The hash is a cryptographic checksum using the Go checksum format (h1:). The `go mod verify` command checks that locally cached modules match go.sum entries.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/20-go-sum
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/20-go-sum
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Committing go.sum without go.mod -- they must go together.
- Manually editing go.sum -- it's auto-generated and must not be hand-edited.
- Ignoring go.sum checksum mismatches -- indicates tampering or corruption.
- Not committing go.sum -- breaks reproducible builds for other developers.

## In Production

CI pipelines check go.sum consistency. Module proxies cache checksummed modules. go.sum is co-committed with go.mod in every Go project.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-21`.
