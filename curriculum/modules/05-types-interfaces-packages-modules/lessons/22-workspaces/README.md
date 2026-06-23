# Workspaces

## Mission

Understand and apply Workspaces in the context of professional Go software engineering.

## Prerequisites

- core-05-21

## Mental Model

A Go workspace is like a shared desk where multiple projects (modules) can be developed together. When one project changes, the others see the changes immediately. The go.work file is the desk's arrangement plan.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

go.work uses a similar format to go.mod with `go`, `use`, and `replace` directives. The `use` directive points to local module directories. When resolving imports, Go checks workspace modules before the module cache or proxy.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/22-workspaces
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/22-workspaces
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Creating a workspace when a single module would suffice.
- Forgetting to run `go work sync` after changing workspace modules.
- Committing go.work when it should be local/go.work is gitignored.
- Adding replace directives in individual go.mod files instead of using workspace.

## In Production

Workspaces are used when developing libraries alongside applications, when monorepos contain multiple modules, and when contributing changes across multiple repositories simultaneously.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-23`.
