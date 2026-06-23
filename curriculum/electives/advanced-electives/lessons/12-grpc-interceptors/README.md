# gRPC interceptors

## Mission

Understand and apply gRPC interceptors in the context of professional Go software engineering.

## Prerequisites

- elective-11

## Mental Model

An interceptor is a function that wraps an RPC handler. The server creates an interceptor chain at startup. For each incoming RPC, gRPC calls the interceptor chain instead of the handler directly. Each interceptor can inspect the request, modify the context, call the next interceptor in the chain, inspect the response, and modify or replace the response. Client interceptors wrap outgoing RPC calls with the same pattern.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A unary server interceptor has the signature: func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error). The interceptor calls handler(ctx, req) to invoke the next interceptor or the final handler. Stream interceptors wrap streaming RPCs with a similar pattern using grpc.StreamServerInfo and grpc.StreamHandler. Client interceptors have analogous signatures for outgoing calls.

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

- Forgetting to call handler(ctx, req) in the interceptor — the handler never executes, and the client gets no response.
- Modifying the request after calling handler — the request is already processed; the interceptor should modify the request before calling handler.
- Using a unary interceptor for streaming RPCs — stream RPCs must use stream interceptors, not unary interceptors.
- Not propagating context values set by interceptors — context values are passed via context.WithValue and must be extracted by downstream interceptors and handlers.
- Blocking in an interceptor — an interceptor that calls time.Sleep or blocks on I/O blocks the entire gRPC server's handler goroutine.

## In Production

gRPC interceptors are used in every production gRPC service: auth interceptors validate JWT tokens from gRPC metadata, logging interceptors log every RPC with duration and status, recovery interceptors catch panics and return gRPC Internal errors, rate limit interceptors reject requests when the rate limit is exceeded, and metrics interceptors record request counts and latency histograms.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-13`.
