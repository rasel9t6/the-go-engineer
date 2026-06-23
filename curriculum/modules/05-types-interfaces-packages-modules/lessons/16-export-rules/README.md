# Export rules

## Mission

Understand and apply Export rules in the context of professional Go software engineering.

## Prerequisites

- core-05-15

## Mental Model

Export rules are like building access control. Exported identifiers are the public lobby -- anyone enters. Unexported identifiers are back offices -- only employees (same package) enter.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Compiler checks export rules at name resolution. When resolving `pkg.FuncName`, verifies FuncName is capitalized. Also enforces `internal/` package restrictions.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/16-export-rules
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/16-export-rules
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming 'exported' means 'public' -- exported is accessible from other packages.
- Using ALL_CAPS for exported constants -- Go doesn't use ALL_CAPS.
- Thinking unexported functions cannot be tested -- test from same package.
- Forgetting types with unexported fields cannot be constructed cross-package.

## In Production

Every Go package uses export rules. Exported types, functions, constants form the API. Helper functions, internal types remain unexported. `internal/` enforces package-group visibility.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-17`.
