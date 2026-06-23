# Parameters

## Mission

Understand and apply Parameters in the context of professional Go software engineering.

## Prerequisites

- core-04-01

## Mental Model

A parameter is a local variable that receives its initial value from the caller's argument. The function signature declares what types the caller must supply, and the function body uses those values.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

In Go's calling convention, each parameter is allocated space in the caller's stack frame. The compiler ensures the argument expression is evaluated and copied into the parameter slot before the function call. For interface parameters, the compiler generates code to create an `iface` or `eface` struct (type pointer + data pointer) and passes that two-word value.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/02-parameters
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/02-parameters
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Declaring a function with parameters but never using them inside the body -- results in a compile error in Go.
- Confusing arguments and parameters -- the parameter is the variable in the function signature, the argument is the value passed at the call site.
- Modifying a parameter inside the function expecting the caller to see the change -- Go passes by value, so the original is unchanged unless a pointer is used.

## In Production

Every HTTP handler, database query function, and API endpoint in Go uses parameters. A typical handler signature `func GetUser(w http.ResponseWriter, r *http.Request)` takes the response writer and request as parameters. Parameters make functions testable by allowing different inputs to produce different outputs.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-03`.
