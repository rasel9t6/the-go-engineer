# PD.3 Project Layout

## Mission

Decide when a Go project should stay flat, when it needs `internal/`, and when `cmd/` or other
top-level directories actually earn their place.

This surface is the package-design track output for Stage 09.

## Files

- [main.go](./main.go): layout guide with small, medium, and large project examples

## Run Instructions

```bash
go run ./09-architecture/01-package-design/3-project-layout
```

## Success Criteria

You should be able to:

- explain when a flat layout is enough
- describe what `cmd/`, `internal/`, and `pkg/` are for
- call out common layout anti-patterns like `utils/`, `helpers/`, and premature folder sprawl

## In Production

Project layout directly impacts team velocity and code maintainability. A common mistake in Go is importing directory structures from other languages (like Java's deep nesting or Rails' strict MVC separation) instead of following idiomatic Go patterns. In production, the `cmd/` directory cleanly separates multiple executable binaries (e.g., a server, a CLI tool, a worker) that share the same core logic. The `internal/` directory is a crucial compiler-enforced boundary that prevents other repositories from importing your private packages, allowing you to refactor without breaking external consumers. Flat structures are preferred until complexity demands separation, because every package boundary introduces cognitive overhead and potential import cycles. Teams that embrace `internal/` and avoid grab-bag folders like `utils/` build codebases that remain navigable even as they scale to hundreds of thousands of lines.

## Thinking Questions

1. Why does the Go compiler enforce the `internal/` directory boundary, and what problem does it solve for large organizations with many repositories?
2. If you put common helper functions in a package named `utils`, why does that eventually lead to cyclic import errors as the project grows?
3. When transitioning a project from a single flat package to a structured layout, what is the most reliable indicator that it is time to split a package?
4. How does separating the `main.go` file in `cmd/` from the business logic in `internal/` make testing the application easier?

## Next Step

After `PD.3`, continue to [Stage 09 overview](../README.md) or move into
[Stage 10 Production Operations](../../../10-production).
