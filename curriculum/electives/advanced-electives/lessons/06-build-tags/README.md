# Build tags

## Mission

Understand and apply Build tags in the context of professional Go software engineering.

## Prerequisites

- elective-05

## Mental Model

A build tag is a boolean expression evaluated at compile time. Files whose build tag expression evaluates to false are excluded from the compilation. Common constraints include: GOOS (linux, darwin, windows), GOARCH (amd64, arm64), Go version (go1.22), custom tags set via -tags flag, and cgo. The expression syntax supports &&, ||, and ! operators for complex constraints.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's go/build package evaluates build constraints during package scanning. The //go:build comment (Go 1.17+) replaces the older // +build syntax. The expression is: //go:build linux && amd64 means the file only compiles on linux/amd64. Multiple //go:build lines are ORed. Filename-based constraints use the pattern *_GOOS.go, *_GOARCH.go, or *_GOOS_GOARCH.go — these are checked before build tag comments.

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

- Using the old // +build syntax without the new //go:build syntax — Go 1.17+ requires both for migration; Go 1.21+ ignores // +build alone.
- Placing the build tag comment after the package declaration — the build tag must be at the top of the file, before the package statement.
- Forgetting that filenames with _linux.go are automatically constrained to linux — adding a //go:build windows to a _linux.go file creates an impossible constraint (file never compiles).
- Using custom build tags inconsistently across files — defining a feature in some files but not others causes partial compilation.
- Not testing all build tag combinations — a build tag combination with a missing file causes a compilation error that is only caught in CI.

## In Production

Build tags are used in every Go project targeting multiple platforms: system-level libraries (database drivers with CGo), platform-specific implementations (file paths, process signals), testing utilities (integration tests with //go:build integration), and debug/logging features (//go:build debug). Production Go code uses build tags for OS-specific system monitoring, Windows service management, and Linux cgroup integration.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-07`.
