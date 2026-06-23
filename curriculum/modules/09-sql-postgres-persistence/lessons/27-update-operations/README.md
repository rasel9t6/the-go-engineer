# UPDATE operations

## Mission

Understand and apply UPDATE operations using sqlc-generated code in Go.

## Prerequisites

- core-09-26

## Mental Model

An UPDATE is a targeted mutation that should affect only the rows explicitly identified by the WHERE clause. Unlike a full replacement (PUT), a partial update (PATCH) changes only the fields the client sends while preserving all other column values.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

PostgreSQL MVCC does not overwrite rows in place. An UPDATE creates a new version of the row and marks the old version as dead, later reclaimed by VACUUM. This means UPDATE generates write-ahead log (WAL) entries proportional to the size of the updated columns.

## Run Instructions

```bash
go run ./curriculum/modules/09-sql-postgres-persistence/lessons/27-update-operations
go test ./curriculum/modules/09-sql-postgres-persistence/lessons/27-update-operations
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Performing a SELECT before every UPDATE to fetch the current row, then sending the unchanged fields back — this doubles the database round trip and is the most common ORM-style anti-pattern.
- Writing UPDATE statements that set every column regardless of what actually changed — this causes unnecessary write locks and invalidates cache entries for unaffected columns.
- Forgetting to add a WHERE clause to an UPDATE — without it every row in the table is overwritten.

## In Production

UPDATE is the second-most common SQL operation after SELECT in production systems. Every API that modifies resources — user profiles, order status, document content — relies on UPDATE with appropriate WHERE clauses and concurrency control.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-09-28`.
