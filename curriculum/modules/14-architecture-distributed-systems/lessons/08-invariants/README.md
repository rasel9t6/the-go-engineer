# Invariants

## Mission

Understand and apply Invariants in the context of professional Go software engineering.

## Prerequisites

- core-14-07

## Mental Model

An invariant is a truth about your data that must always hold. It is not a goal — it is a constraint. 'Account balance >= 0' is an invariant. 'Accounts should generally have positive balances' is a goal. Invariants are enforced with checks at every write path, ideally at the database level (CHECK constraints) or at the service layer with a mutex or serialized access. Every read path should also check invariants defensively — stale data from migrations or manual edits can violate invariants silently.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Database isolation levels control how concurrent transactions interact. READ COMMITTED (default in PostgreSQL) shows only committed data but allows non-repeatable reads — two reads in the same transaction may see different values. REPEATABLE READ prevents non-repeatable reads but allows phantom reads. SERIALIZABLE prevents all concurrency anomalies but requires retry on serialization failure. SELECT FOR UPDATE locks the selected rows, preventing concurrent writes and concurrent SELECT FOR UPDATE from other transactions. The lock is held until the transaction commits or rolls back. In Go, use db.BeginTx with a context and sql.TxOptions{Isolation: sql.LevelSerializable} for the strongest isolation, but be prepared to retry on serialization errors.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/08-invariants
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/08-invariants
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming the database enforces all invariants — application-level invariants (e.g., 'order total equals sum of line items') are not expressible as database constraints and must be checked in code.
- Only checking invariants at write time, never at read time — a data migration, backfill, or manual DB edit can create invariant violations that are invisible until the data is read months later.
- Checking invariants in the wrong layer — validation in the handler passes, but the service layer has a different code path that bypasses validation, creating inconsistent state.
- Concurrent invariant violation — two goroutines read the same balance (both see $100), both approve a $90 withdrawal, and both write $10 — the invariant 'balance >= 0' is violated because the second write was based on stale data.
- Using panic for invariant violations in library code — an invariant check in a library panics the entire program instead of returning an error, making recovery impossible.

## In Production

Invariant enforcement is the foundation of data integrity in every production system. Financial systems enforce 'account balance >= 0' and 'sum of transactions = balance change'. E-commerce systems enforce 'inventory allocated <= inventory available' and 'order total = sum of line item totals'. Booking systems enforce 'seat is allocated to at most one booking'. Every violation of these invariants is a data corruption incident that requires manual reconciliation.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-09`.
