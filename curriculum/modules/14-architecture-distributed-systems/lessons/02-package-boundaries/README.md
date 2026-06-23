# Package boundaries

## Mission

Understand and apply Package boundaries in the context of professional Go software engineering.

## Prerequisites

- core-14-01

## Mental Model

A package boundary is a contract wall. Inside the wall, code can access everything (exported and unexported). Outside the wall, only the officially exported doorways (exported functions, types, constants) are accessible. The wall has a cost: every time a developer needs to cross it, they must use the official doorway. The benefit: the code inside the wall can be refactored, renamed, or rewritten as long as the doorways stay the same. A good package boundary is one where the cost of crossing is justified by the freedom it gives the code inside.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's import resolution works as follows: the compiler resolves each import path to a directory, reads all .go files in that directory with the same package declaration, and builds a symbol table of exported names. Import cycles are detected during dependency graph construction: the compiler walks the import graph depth-first, marking each package as 'visiting.' If it encounters a package that is already 'visiting,' it reports an import cycle and stops. This check runs once per build and is the fastest possible cycle detection (O(V+E) where V is packages and E is imports). The internal/ convention is enforced by go/build: during import resolution, the tool checks whether the importing package's root is an ancestor of the internal/ package's directory. If not, the import is rejected with a compile error.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/02-package-boundaries
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/02-package-boundaries
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Putting every type in one package — a single models/ or types/ package with 200 exported types has no internal cohesion. Every file imports models, creating coupling across the entire codebase. Any change to any type requires recompiling every dependent package.
- Making packages too small — a package with one 5-line type and one 3-line function adds import overhead without any benefit. The cost of understanding a package boundary (what is exported, what is internal, what does the package name mean) must be justified by the cohesion within the package.
- Circular imports between related packages — users.go imports orders.go and orders.go imports users.go. The Go compiler rejects this, forcing the developer to either merge the packages or extract shared types into a third package. The circular import is a signal that the boundary is wrong.

## In Production

Package boundaries are the most important architectural decision in Go projects. Kubernetes has 300+ packages organized by domain (pkg/api/, pkg/kubelet/, pkg/scheduler/). Docker organizes by component (daemon/, builder/, cli/). CockroachDB organizes by SQL layer (sql/, sql/parser/, sql/exec/). Every major Go project uses packages as the primary architecture primitive. The pattern is consistent: one domain per package, small exported surface, acyclic imports enforced by the compiler.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-03`.
