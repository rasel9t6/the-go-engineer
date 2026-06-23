# Transactions

## Mission

Understand and apply Transactions in the context of professional Go software engineering.

## Prerequisites

- core-09-16

## Mental Model

A database transaction groups multiple operations into a single atomic unit that either commits entirely or rolls back entirely. Transactions provide isolation so concurrent readers see a consistent snapshot. Go's database/sql exposes BeginTx, Commit, and Rollback for explicit transaction control.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, tx = db.BeginTx interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that group operations into database transactions is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/09-sql-postgres-persistence/lessons/17-transactions
go test ./curriculum/modules/09-sql-postgres-persistence/lessons/17-transactions
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming group operations into database transactions does something it does not — read the package docs instead of guessing behavior.
- Hardcoding Transactions values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying Transactions code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use Transactions daily in production services, CLI tools, API servers, and data pipelines. Understanding Transactions is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-09-18`.
