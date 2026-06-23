# Hexagonal architecture

## Mission

Understand and apply Hexagonal architecture in the context of professional Go software engineering.

## Prerequisites

- core-14-05

## Mental Model

Hexagonal architecture is a power outlet. The domain is the device (lamp, phone charger). The port is the outlet shape (Type A, Type C). The adapter is the plug that connects the device to the outlet. The device does not know what power source is behind the outlet (grid, generator, battery). The outlet does not know what device is plugged in. The adapter translates between the device's needs and the source's capabilities. In Go: the domain defines the outlet (interface). The adapter (Postgres repo, HTTP handler) implements the outlet shape. main.go is the electrician that wires them together.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's implicit interface satisfaction is what makes hexagonal architecture possible without annotations or code generation. An adapter type satisfies a domain port interface simply by having methods with the right names and signatures. The compiler checks this at the point of assignment: var repo domain.OrderRepository = postgres.NewOrderRepo(db) — if postgres.OrderRepo does not have the right methods, the assignment fails to compile. This is a compile-time check with zero runtime cost. The adapter package imports the domain package (to access the interface and domain types), but the domain package never imports the adapter package. This creates the desired dependency direction: infrastructure depends on domain, not the other way around.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/06-hexagonal-architecture
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/06-hexagonal-architecture
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Creating ports for every possible adapter — defining 20 interfaces (UserReader, UserWriter, UserNotifier, UserValidator) before any adapter exists. The ports should emerge from the adapters the application actually needs, not predicted upfront. Start with one interface per adapter type (UserRepository, EmailSender, PaymentGateway) and add more as needed.
- Letting the domain depend on framework types — a domain service that imports gin.Context, *sql.DB, or prometheus.Counter. The domain should be pure Go with zero framework imports. Framework dependencies belong in the adapter layer, not the domain.
- Making hexagonal architecture an all-or-nothing choice — applying hexagonal architecture to every package, including utility functions and configuration reading. Hexagonal architecture is for the domain core: business logic, entities, and use cases. Configuration, CLI parsing, and logging setup are infrastructure and do not need ports and adapters.

## In Production

Hexagonal architecture is the standard pattern for Go services that need to outlive their initial infrastructure choices. Uber's Go monorepo uses hexagonal architecture for critical payment and pricing services. Monzo's Go backend uses ports and adapters for banking core services. Kubernetes' controller-runtime library follows the hexagonal pattern: controllers define ports (reconcile interfaces), and adapters implement them for different infrastructure backends.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-07`.
