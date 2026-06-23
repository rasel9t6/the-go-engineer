# When to split services

## Mission

Understand and apply When to split services in the context of professional Go software engineering.

## Prerequisites

- core-14-14

## Mental Model

Splitting a monolith is like moving from a shared house to separate apartments. The shared house (monolith) has one kitchen, one bathroom, one front door — everything is shared, coordination is simple, but everyone must agree on everything. Separate apartments (services) have their own kitchen, bathroom, and door — each team can paint their walls any color, but now they need to visit each other (network calls) to borrow things. The key question: when does the cost of visiting (network calls, data duplication, deployment complexity) become less than the cost of coordination in the shared house?

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Service extraction in Go follows a mechanical process: (1) identify the package to extract — it already has a well-defined exported interface (the Go package API). (2) Create a new Go module (go mod init service-name). (3) Copy the package and its dependencies into the new module. (4) Add a server that exposes the package's exported methods over gRPC or HTTP. (5) In the original monolith, replace the direct package call with a gRPC/HTTP client call. (6) Add observability (tracing, metrics, logging) to the network boundary. (7) Verify that the behavior is identical — run integration tests that cover both in-process and network paths.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/15-when-to-split-services
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/15-when-to-split-services
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Splitting a service because a single binary feels 'unclean' — this is the most common mistake. A monolith that is well-structured (modular monolith with domain packages) handles 98% of use cases. Splitting adds: network latency (1-10ms per call), distributed tracing, data consistency challenges, deployment coordination, and operational complexity. Only split when there is a concrete, measured need.
- Splitting by layer instead of domain — creating separate services for 'handlers', 'services', and 'repositories'. Each service is tightly coupled to the others (handlers need services, services need repositories). The split increases complexity without any of the benefits of independence. Split by domain boundary: UserService, OrderService, PaymentService — each owns its full stack.
- Splitting before the data access pattern is understood — a monolith makes JOINs and transactions easy. After splitting, the same queries require API calls between services or an API composition layer. If the team does not know which queries need to JOIN across domains, they will create a chatty, high-latency architecture.

## In Production

The standard advice for software architecture is 'start with a monolith, extract services when you have a proven need.' Amazon ran a monolith for years before splitting. eBay ran a monolith for a decade. Etsy still runs a monolith. Shopify extracted its first service after 8 years. The key insight: premature splitting creates a distributed monolith with all the complexity and none of the benefits. The right time to split is when the cost of coordination in the monolith exceeds the cost of network calls between services.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-16`.
