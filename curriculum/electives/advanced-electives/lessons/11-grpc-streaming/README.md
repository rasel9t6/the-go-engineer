# gRPC streaming

## Mission

Understand and apply gRPC streaming in the context of professional Go software engineering.

## Prerequisites

- elective-10

## Mental Model

A gRPC stream is a series of messages exchanged over a single HTTP/2 stream. Server-streaming: the client sends one request and receives multiple responses sent by the server. Client-streaming: the client sends multiple requests and the server sends one response. Bidirectional: both sides send multiple messages independently — the stream is full-duplex.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

gRPC streams use HTTP/2 frames. The client opens a new HTTP/2 stream for each RPC. For server-streaming, the server sends multiple DATA frames (each containing a serialized protobuf message) followed by a HEADERS frame with the response status. The client reads messages from the stream with Recv() until io.EOF. For client-streaming, the client sends multiple DATA frames and the server sends a single response. Bidirectional streams interleave DATA frames from both sides.

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

- Not handling io.EOF on server-streaming Recv() — io.EOF signals the end of the stream and is not an error; forgetting to check for io.EOF causes the client to treat stream end as a failure.
- Blocking Send() on a slow client — if the client is not reading fast enough, the server's Send() blocks, potentially backing up to the entire gRPC server. Use a buffer or flow control.
- Calling Send() and Recv() from the same goroutine in bidirectional streaming — the methods block, causing a deadlock. Use separate goroutines for send and receive.
- Not setting deadlines on streaming RPCs — a streaming RPC that hangs indefinitely holds resources on both client and server. Set per-RPC or per-stream deadlines.
- Using too many concurrent streams — each stream consumes goroutines and memory on both ends. Limit concurrent streams with gRPC's MaxConcurrentStreams server option.

## In Production

gRPC streaming is used for real-time features: log streaming (server-streaming), file uploads with progress (client-streaming), chat and collaboration (bidirectional), event feeds (server-streaming), and large dataset pagination (server-streaming). Go services use server-streaming for API endpoints that return paginated results without page boundaries.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-12`.
