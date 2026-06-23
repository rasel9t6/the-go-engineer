# Trust boundaries

## Mission

Understand and apply Trust boundaries in the context of professional Go software engineering.

## Prerequisites

- core-10-01

## Mental Model

A trust boundary is any point where data crosses from less trusted to more trusted. The network edge is one boundary; every RPC call and database query is another. Each is a gate that validates trust before allowing passage.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Trust boundaries in Go are implemented by composing http.Handler with middleware. This is the 'wall and gate' pattern: the wall separates trust zones, and the gate controls passage.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/02-trust-boundaries
go test ./curriculum/modules/10-auth-security/lessons/02-trust-boundaries
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Drawing trust boundaries around deployment boundaries instead of data flow boundaries.
- Assuming the network perimeter is the only trust boundary — in modern architectures, boundaries exist between services and even between goroutines.
- Placing the trust boundary inside the request handler instead of at the middleware layer.

## In Production

Production Go services use trust boundaries to isolate public API vs admin API, user data vs system data, and staging vs production.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-03`.
