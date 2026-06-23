# Service layer

## Mission

Understand and apply Service layer in the context of professional Go software engineering.

## Prerequisites

- core-14-02

## Mental Model

The service layer is the brain of the application. The handler is the mouth (talks to the outside world). The repository is the hands (talks to the database). The brain decides what to do (business rules), the mouth receives input and delivers output, and the hands execute the storage. The brain should not do the hands' job (SQL queries) or the mouth's job (HTTP parsing). The brain should be testable without the mouth or hands — inject mock dependencies and test the business logic in isolation.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The service layer in Go is enabled by two language features: interfaces and package visibility. Interfaces allow the service to define its dependencies abstractly (type UserRepository interface { Insert(ctx, User) error }). The concrete type satisfaction is implicit — any type with an Insert method with the right signature satisfies the interface. Package visibility (exported/unexported) enforces the layer boundary: the service package exports only business types and interfaces; implementation details are unexported. The `internal/` directory convention prevents external consumers from importing implementation packages directly. The service layer does not depend on any specific transport or storage — it is a pure Go package that could be used in any context.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/03-service-layer
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/03-service-layer
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Putting business logic in HTTP handlers — the handler parses the request, calls the database, formats the response, and handles errors. When the same logic is needed for a gRPC endpoint or a background job, the handler code cannot be reused. Business logic must live in the service layer, not the transport layer.
- Making the service layer a thin pass-through — the handler calls service.CreateUser which calls repo.InsertUser with no transformation, validation, or business logic. The service layer adds complexity without value. A service layer must justify its existence by containing business rules, not just forwarding calls.
- Letting services become god objects — internal/service/user.go has 3000 lines with 50 methods covering user CRUD, billing, notifications, and reporting. The service layer should be split by bounded context, not by entity. User billing is a different service than user profile management.

## In Production

The service layer pattern is universal in production Go services. Stripe's API handlers call the service layer for every operation. Kubernetes' API server has a service layer between the HTTP handler and the storage backend. Docker's engine API uses a service layer for image and container operations. The pattern is: handler → service → repository (or external API). The service layer is the only layer that contains business logic — handlers and repositories are mechanical.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-04`.
