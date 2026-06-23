# CQRS

## Mission

Understand and apply CQRS in the context of professional Go software engineering.

## Prerequisites

- elective-15

## Mental Model

CQRS splits the system into two sides. The command side receives write operations, validates them, and appends events to an event store. The query side reads from denormalized read models that are optimized for specific queries. Events propagate from the command side to the query side via an event bus. The read model is eventually consistent with the write model — there is a propagation delay.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Commands are validated, applied to the domain model, and persisted as events in an event store (append-only log). The event store is the source of truth. A projection (event handler) subscribes to events and updates a read model (typically a denormalized database table or a separate read-optimized store). The read model is rebuilt by replaying all events from the beginning. Read models can be specialized: one for order listing (denormalized with customer name and total), another for order statistics (aggregated counts).

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

- Applying CQRS to CRUD apps — CQRS adds complexity (event bus, separate stores, eventual consistency) that is not justified for simple CRUD.
- Sharing the same database between command and query sides — the database is a coupling point; use separate databases or schemas.
- Ignoring eventual consistency — users see stale data after a write; communicate the propagation delay and use UI patterns (optimistic updates).
- Using synchronous event propagation — the command handler waits for the read model update before returning, eliminating the benefits of CQRS.
- Not having a rebuild mechanism — read models can become corrupt or out of sync; provide a way to rebuild them from the event store.

## In Production

CQRS is used in high-scale systems where read and write workloads diverge: e-commerce (order writes are transactional, product catalog reads are denormalized and cached), banking (transaction writes require strict validation, balance queries need fast lookups), and analytics (write events to a log, read aggregated projections). Go services use CQRS with NATS for event propagation and PostgreSQL for read models.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-17`.
