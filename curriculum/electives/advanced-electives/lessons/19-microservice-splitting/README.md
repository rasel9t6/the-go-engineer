# Microservice splitting

## Mission

Understand and apply Microservice splitting in the context of professional Go software engineering.

## Prerequisites

- elective-18

## Mental Model

A monolith is a single deployable unit with all features and data in one codebase and database. Microservice splitting extracts bounded contexts (groups of related features) into separate services. Each service has its own codebase, its own database, its own CI/CD pipeline, and its own on-call team. Services communicate via APIs (HTTP, gRPC, message queues). The monolith shrinks as services are extracted.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Bounded contexts are identified by analyzing the monolith's domain: which features change together? Which features share data? Which features have different scaling or latency requirements? Each bounded context becomes a candidate for extraction. The extraction process: (1) identify the bounded context, (2) define the API contract (proto file or OpenAPI spec), (3) extract the data (new database with migrated schema), (4) extract the code (new service with its own repo), (5) route traffic from monolith to new service.

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

- Splitting too early — the bounded context is not well understood; frequent API changes couple the services more tightly than the monolith did.
- Sharing a database between services — services share data via API calls, not database queries; a shared database creates coupling that prevents independent deployment.
- Creating a distributed monolith — services that call each other synchronously in a chain; a failure in one service cascades through all services.
- Not handling network failures — a service call can fail, timeout, or return garbage; every cross-service call must have retries, timeouts, and circuit breakers.
- Splitting by technical layer instead of business capability — having a 'data service', 'logic service', and 'presentation service' instead of 'order service', 'payment service', 'inventory service'.

## In Production

Microservice splitting is used by every large-scale organization: Netflix (500+ services), Amazon (1000+ services), Uber (2200+ services). Go is a popular choice for microservices due to its small binary size, fast startup, and efficient resource usage. Opslane uses Go microservices for orders, payments, inventory, and notifications — each with its own PostgreSQL database and independent deployment pipeline.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-20`.
