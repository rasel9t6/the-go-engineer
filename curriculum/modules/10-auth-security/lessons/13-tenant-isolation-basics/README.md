# Tenant isolation basics

## Mission

Understand and apply Tenant isolation basics in the context of professional Go software engineering.

## Prerequisites

- core-10-12

## Mental Model

Tenant isolation is the guarantee that one tenant cannot access another tenant's data. The database is the locked filing cabinet inside that room.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Shared-table isolation uses a tenant_id column on every table with a database index. Row-level security (RLS) in PostgreSQL enforces tenant filtering at the database level, preventing application bugs from leaking data. Schema-per-tenant isolation creates separate database schemas for each tenant.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/13-tenant-isolation-basics
go test ./curriculum/modules/10-auth-security/lessons/13-tenant-isolation-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using SELECT * from shared tables without a tenant WHERE clause — one missing filter leaks all tenants' data.
- Implementing tenant isolation in application code without database-level enforcement.
- Storing tenant_id as a string without validation — path traversal-like attacks on tenant context.
- Sharing connection pools across tenants without context propagation.

## In Production

Every SaaS platform implements tenant isolation. GitHub (organization-scoped data), Slack (workspace isolation), and AWS (account isolation) all use tenant-scoped access patterns.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-14`.
