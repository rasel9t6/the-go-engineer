# Domain modeling

## Mission

Understand and apply Domain modeling in the context of professional Go software engineering.

## Prerequisites

- core-14-06

## Mental Model

Domain modeling is making illegal states unrepresentable. If an Order cannot have a Shipped status without a ShipDate, then the Order type should not have a Status field and a ShipDate field independently. Instead, the Order should have a union-like structure: type Order struct { Pending *PendingOrder; Shipped *ShippedOrder }. Or simpler: provide methods that enforce the invariant: order.Ship(now) sets both Status and ShipDate atomically. The goal is that no valid Go program can create an invalid domain state.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's type system supports domain modeling through: (1) type definitions (type OrderID string) that create distinct types with no implicit conversions; (2) method declarations on any named type; (3) unexported fields that prevent external mutation; (4) constructor functions (NewOrder) that validate invariants at creation time; (5) error return values that make invalid operations explicit. Go does not have algebraic data types (Rust's enum, Haskell's sum types), so domain modelers use the 'validate on construction' pattern: accept all inputs in a constructor, validate business rules, return error if invalid, and return a valid entity if all rules pass. Once constructed, the entity's methods maintain the invariant.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/07-domain-modeling
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/07-domain-modeling
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Making domain types anemic — a User struct with only fields and getters/setters. All business logic lives in services. This is not domain modeling — it is data structures with no behavior. Domain types should encapsulate business rules: user.Cancel() should enforce that the user cannot be canceled after 30 days.
- Using primitive obsession — representing OrderStatus as a string, Money as a float64, UserID as an int. The compiler cannot catch type errors: passing an OrderID where a UserID is expected compiles fine. Fix: use value types: type OrderID string, type Money struct { amount Decimal; currency Currency }.
- Modeling the database schema instead of the domain — the User struct has CreatedAt, UpdatedAt, DeletedAt timestamps because the DB table has them, but the domain does not need them. The User domain type should only have fields that the business logic needs. Database concerns (timestamps, version numbers) belong in the repository mapping.

## In Production

Domain modeling in Go is used by projects where business rules are complex and change frequently. Kubernetes models its API resources as domain types with validation methods (Validate(), Mutate()) on the types themselves. Docker's engine models images, containers, and volumes with domain methods for lifecycle operations. Monzo's banking core models accounts, transactions, and payments with strong domain types that prevent invalid financial states. The pattern: when the business logic is complex, put it on the types.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-08`.
