# REST vs gRPC tradeoffs

## Mission

Understand and apply REST vs gRPC tradeoffs in the context of professional Go software engineering.

## Prerequisites

- elective-12

## Mental Model

REST and gRPC are different tools for the same job (service-to-service communication) with different tradeoffs. REST: resources as URLs, JSON payloads, HTTP methods, cacheable, browser-native. gRPC: services as .proto contracts, protobuf binary, HTTP/2, streaming, code-generated clients. REST prioritizes accessibility and simplicity; gRPC prioritizes performance and contract rigor.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

REST over HTTP/1.1 uses text-based JSON, one request per TCP connection (without keep-alive), and manual documentation of endpoints. gRPC over HTTP/2 uses binary protobuf, multiplexed streams over one connection, automatic code generation, and built-in service discovery. Performance difference: protobuf is 3-10x faster to serialize and 30-70% smaller payloads than JSON. HTTP/2 multiplexing eliminates connection overhead for multiple concurrent requests.

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

- Using gRPC for browser-facing APIs without gRPC-Web — browsers do not natively support gRPC; use gRPC-Web proxy or REST for browser clients.
- Using REST for high-throughput internal microservice APIs — JSON serialization overhead adds latency and bandwidth costs at scale.
- Not leveraging HTTP caching for REST APIs — REST responses can be cached at the CDN or client level; gRPC responses are not cacheable without a proxy.
- Forgetting that gRPC requires protobuf schema management — .proto files must be shared between services; breaking changes require coordinated deployments.
- Choosing REST just because it is simpler — REST's simplicity comes at the cost of performance, contract enforcement, and streaming support.

## In Production

Most Go microservices use both REST and gRPC: REST for external APIs (public-facing, browser clients, third-party integrations) and gRPC for internal APIs (service-to-service, high-throughput, streaming). Companies like Google, Uber, and Netflix use gRPC internally and REST externally. The grpc-gateway project generates REST endpoints from .proto files, allowing a single service contract for both protocols.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-14`.
