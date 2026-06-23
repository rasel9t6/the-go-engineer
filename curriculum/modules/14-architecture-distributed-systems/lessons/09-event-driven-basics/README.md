# Event-driven basics

## Mission

Understand and apply Event-driven basics in the context of professional Go software engineering.

## Prerequisites

- core-14-08

## Mental Model

Event-driven architecture is a party where services talk through a megaphone (the broker) instead of private phone calls (direct API calls). Service A shouts 'Order Created!' into the megaphone. Services B, C, and D hear it and do their thing. Service A does not know or care who is listening. If Service B is in the bathroom (down), it will hear the message when it comes back (event persistence). The megaphone ensures every message is heard at least once. Each service must be prepared to hear the same message twice (at-least-once delivery) and handle it gracefully (idempotency).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Message brokers use a publish-subscribe model. The publisher sends a message to a topic/exchange. The broker stores the message (persistently for Kafka/PubSub, in-memory for RabbitMQ without persistence). Consumers subscribe to the topic with a consumer group. Each message is delivered to one consumer in the group (queue semantics) or all consumers (pub-sub semantics). After processing, the consumer sends an acknowledgment. If the consumer crashes before acknowledging, the broker redelivers the message to another consumer in the group. This is at-least-once delivery. To achieve exactly-once semantics, the consumer must be idempotent (processing the same message twice produces the same result). In Go, the consumer library (e.g., Sarama for Kafka, amqp for RabbitMQ) runs the consumer loop in a goroutine and calls the handler function for each message.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/09-event-driven-basics
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/09-event-driven-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Publishing events directly from the handler — the handler publishes an 'order.created' event during HTTP request processing. If the event publish fails (broker is down), the request returns an error even though the database write succeeded. Events should be published after the transaction commits, not during.
- Assuming events are delivered exactly once — message brokers guarantee at-least-once delivery, not exactly-once. An event may be delivered twice if the broker crashes between publishing and acknowledging. Consumers must be idempotent: processing the same event twice produces the same result.
- Coupling event format to internal types — publishing a Go struct as JSON means changing the struct field name breaks all consumers. Events are a contract: define a schema (protobuf, Avro, or a versioned JSON schema) and map between internal types and event types.

## In Production

Event-driven architecture is the backbone of every large-scale distributed system. Kafka processes 1M+ events/second at Netflix for real-time streaming. RabbitMQ handles 10K+ messages/second for task distribution at Docker Hub. Google Pub/Sub processes events for Google Ads at planetary scale. In Go production services, event-driven communication is used for: order processing pipelines, notification systems, data replication, analytics pipelines, and workflow orchestration.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-10`.
