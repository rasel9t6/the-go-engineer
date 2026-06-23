# Deeper distributed systems

## Mission

Understand and apply Deeper distributed systems in the context of professional Go software engineering.

## Prerequisites

- elective-19

## Mental Model

A distributed system is a collection of independent nodes that communicate over an unreliable network. Each node can fail independently, and the network can partition (some nodes cannot reach others). The fundamental challenge is making progress despite partial failures. Patterns: timeouts detect failures, retries handle transient errors, circuit breakers prevent cascading failures, bulkheads isolate failures, and idempotency ensures safe retries.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The network is not reliable: packets can be lost, delayed, duplicated, or reordered. Timeouts detect failures but cannot distinguish between 'node crashed' and 'node is slow but will respond eventually'. Consensus protocols (Raft) use a quorum of nodes (majority) to agree on state changes — a node cannot commit a change without approval from the majority. The CAP theorem states that in a partition, you must choose between consistency (all nodes see the same data) and availability (every request gets a response).

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

- Assuming the network is reliable — every network call can fail; always have retry logic with exponential backoff and jitter.
- Not setting timeouts — a network call without a timeout blocks the goroutine indefinitely, leaking memory and connections.
- Using synchronous calls in a chain — a failure in any service cascades to all upstream services; use async patterns or circuit breakers.
- Ignoring clock skew — distributed timestamps are unreliable; use monotonic clocks for duration and logical clocks for ordering.
- Not testing for network failures — a system that has never been tested with network partitions will fail when a partition occurs.

## In Production

Distributed systems patterns are used in every multi-service architecture: etcd provides distributed consensus for Kubernetes, Cassandra uses eventual consistency for high availability, Kafka uses distributed logs for event streaming, and Redis Sentinel uses failure detection for high availability. Go services use circuit breakers (gobreaker), retries (go-resiliency), and rate limiters (golang.org/x/time/rate) for resilience.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-21`.
