# Multi-tenancy architecture

## Mission

Understand and apply Multi-tenancy architecture in the context of professional Go software engineering.

## Prerequisites

- core-14-12

## Mental Model

Multi-tenancy is a hotel. Each tenant is a guest in their own room. The hotel has shared spaces (lobby, elevator — shared infrastructure) but each room is private. In software: each tenant has their own data (their room), but they share the application instance (the hotel building). The key challenge is making sure Guest A never accidentally enters Guest B's room — every door (database query, cache lookup, API call) must be scoped to the correct tenant.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's context.Context is the mechanism for propagating tenant ID through the call chain. The tenant middleware sets the value: ctx = context.WithValue(ctx, tenantKey, tenantID). Handlers, services, and repositories extract it: ctx.Value(tenantKey).(string). The tenant registry is typically an in-memory map loaded from a database or configuration file at startup, with periodic refresh. The per-tenant connection pool is a map[string]*sql.DB where each pool is created with sql.Open with tenant-specific connection string. The pools are cached for the lifetime of the process, with health checks to reconnect on failure.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/13-multi-tenancy-architecture
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/13-multi-tenancy-architecture
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using a single database with tenant_id on every table and assuming it is sufficient — a missing WHERE tenant_id = ? clause in a query leaks data across tenants. This is the most common multi-tenancy bug and is invisible until a tenant complains. Every query must be scoped to the tenant, and no default query should return data without tenant filtering.
- Sharing connection pools across tenants — a single *sql.DB shared by all tenants means one tenant's slow query consumes all connections, starving other tenants. Each tenant should have a separate connection pool (or at least separate pool settings with max connections per tenant).
- Not isolating tenant configuration — all tenants share the same feature flags, rate limits, and pricing tiers. When tenant A needs a higher rate limit, it must be added to the code. Fix: store tenant configuration in a per-tenant settings record and load it at request time.

## In Production

Multi-tenancy is the standard architecture for SaaS platforms. GitHub uses organization-scoped data access (each org is a tenant). Slack uses workspace-scoped data (each workspace is a tenant). Every B2B SaaS application implements multi-tenancy. The isolation level varies: low-trust tenants (competitors sharing the same platform) need separate databases; high-trust tenants (departments of the same company) can share a database with tenant_id scoping.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-14`.
