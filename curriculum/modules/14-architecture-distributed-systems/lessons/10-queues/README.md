# Queues

## Mission

Understand and apply Queues in the context of professional Go software engineering.

## Prerequisites

- core-14-09

## Mental Model

A queue is a to-do list. The producer adds tasks to the list. Workers pick up tasks, do the work, and mark the task as done. If a worker crashes while working, the task goes back on the list for another worker. The list has a maximum size — if it fills up, the producer must slow down (backpressure). Tasks that keep failing go to a 'lost and found' bin (dead-letter queue) for manual handling.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

An AMQP message (RabbitMQ) has: body (payload), properties (content-type, delivery-mode, priority, expiration), and delivery info (exchange, routing-key, redelivered, delivery-tag). The consumer receives messages via a channel (AMQP channel, not Go channel). Each message must be acknowledged or rejected. If the consumer disconnects without acknowledging, unacked messages are redelivered to another consumer. Kafka uses a pull model: consumers poll for new messages and commit offsets. The offset is the position in the partition. If the consumer crashes, it resumes from the last committed offset. At-least-once delivery is guaranteed by: the broker stores messages persistently and consumers acknowledge after processing.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/10-queues
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/10-queues
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using a message queue as a database — storing state in queue messages and expecting to query them. Queues are for transient work distribution, not persistent storage. Messages should be treated as ephemeral: once processed, they are gone. If you need to query past events, use a database or event store.
- Blocking the consumer goroutine on a slow downstream call — consumer reads a message, calls a 30-second external API, and blocks the consumer goroutine. If the consumer has 10 goroutines and all are blocked, the queue backs up. Fix: use async processing or increase goroutine count with backpressure.
- Not setting a visibility timeout or TTL — a consumer takes a message but crashes without acknowledging it. The message stays invisible forever (or until the broker restarts), causing work to be lost. Fix: always set visibility timeout and message TTL. The broker redelivers after the timeout expires.

## In Production

Queues are used in every production Go service that processes background work. SendGrid uses queues to send millions of emails/day. Stripe uses queues for payment processing and webhook delivery. Docker Hub uses RabbitMQ for image push notifications. The pattern is universal: any work that should not block the HTTP response goes through a queue.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-11`.
