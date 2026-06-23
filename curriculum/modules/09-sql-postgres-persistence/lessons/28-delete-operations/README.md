# DELETE operations

## Mission

Understand and apply DELETE operations using sqlc-generated code in Go.

## Prerequisites

- core-09-27

## Mental Model

DELETE is a destructive operation that removes rows permanently (hard delete) or marks them as inactive (soft delete). The WHERE clause is the safety mechanism — without it, every row in the table is removed. Soft-delete adds a recovery window by marking the row instead of removing it.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

PostgreSQL MVCC handles DELETE by marking the row as dead rather than immediately reclaiming the space. Dead rows are later removed by VACUUM. Each DELETE writes a WAL entry recording the deletion. ON DELETE CASCADE is handled by the database automatically by deleting child rows in the same transaction.

## Run Instructions

```bash
go run ./curriculum/modules/09-sql-postgres-persistence/lessons/28-delete-operations
go test ./curriculum/modules/09-sql-postgres-persistence/lessons/28-delete-operations
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Running DELETE without a WHERE clause in an automated migration or script — this truncates the entire table without the ability to roll back individual rows.
- Cascading deletes that remove more data than the caller expected because foreign-key CASCADE propagates to related tables silently.
- Hard-deleting rows that are referenced by foreign keys in other tables — this causes constraint violations or orphaned records.

## In Production

Most production systems use soft-delete for user-facing resources (accounts, documents, orders) to provide recovery and audit trails. Hard DELETE is reserved for cleaning up temporary data, test fixtures, or PII after the legal retention period expires.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
