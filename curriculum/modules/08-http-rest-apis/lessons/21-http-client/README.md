# HTTP Client

## Mission

Understand and apply gRPC and Protocol Buffers in the context of professional Go software engineering.

## Prerequisites

- core-08-20

## Mental Model

An HTTP client sends requests to external services and processes their responses. Go's http.Client manages connection pooling, timeouts, and transport configuration. Always close response bodies, configure timeouts, and reuse a single client across requests.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

gRPC uses HTTP/2 as its transport. Each RPC call maps to an HTTP/2 stream, allowing multiple concurrent calls over a single TCP connection. Protocol Buffers encode messages as binary using tag-length-value encoding, where each field is identified by its unique tag number rather than its name — this is why renaming a field is safe but renumbering breaks compatibility.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/21-http-client
go test ./curriculum/modules/08-http-rest-apis/lessons/21-http-client
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Defining a new proto message type for every RPC call instead of reusing common message types across the service boundary.
- Treating gRPC as a drop-in replacement for REST without designing for streaming, backpressure, or connection reuse.
- Forgetting to set client-side deadlines and server-side timeouts — a stuck gRPC stream can hold resources indefinitely.

## In Production

gRPC is the standard for inter-service communication in cloud-native architectures. Companies like Uber, Netflix, and Docker use gRPC for service-to-service RPC, replacing REST in high-throughput paths where contract enforcement and performance matter.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
