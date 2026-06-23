# Repository pattern deep dive

## Mission

Understand and apply Repository pattern deep dive in the context of professional Go software engineering.

## Prerequisites

- core-14-03

## Mental Model

The repository is a translation layer between two worlds: the domain world (User, Order, Product) and the database world (rows, columns, SQL, connections). The service lives in the domain world and speaks only domain language. The repository lives in both worlds: it receives domain types, translates them to SQL, executes queries, and translates the results back to domain types. The service should never see a *sql.Row, never write a WHERE clause, and never manage a connection pool. If the database changes from Postgres to MySQL, only the repository changes — the service is isolated.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's database/sql package provides a generic SQL interface. sql.DB is a connection pool (not a single connection) — it maintains a pool of connections, manages keepalives, and retries on transient failures. The repository uses sql.DB methods: QueryContext (returns *sql.Rows), QueryRowContext (returns *sql.Row), ExecContext (returns sql.Result). Each method takes a context for cancellation and timeout. The repository's Scan method calls row.Scan, which uses reflection to map DB column types to Go types. The repository also maps DB errors to domain errors: sql.ErrNoRows → NotFoundError, pq.Error with code 23505 → ConflictError (duplicate key). This error mapping ensures the service layer never imports database-specific packages.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/04-repository-pattern-deep-dive
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/04-repository-pattern-deep-dive
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Defining repository interfaces in the implementation package — the postgres package exports a UserRepository interface and the service imports it. This forces the service to depend on the storage package for a type definition, creating a dependency in the wrong direction. The interface must be defined in the package that NEEDS it (the service), not the package that IMPLEMENTS it (the repository).
- Making repository methods that mirror SQL exactly — FindByID, FindByName, FindByEmail, FindByStatus creates one method per query. When a new query is needed (FindByRoleAndStatus), a new method must be added. Fix: use a query specification pattern — FindUsers(ctx, spec UserQuerySpec) where spec contains the filter criteria. This reduces the interface surface and supports any combination of filters.
- Leaking database types through the repository interface — GetUser returns *sql.Row or the repository accepts *sql.DB in its constructor. The service layer must understand database types to use the repository. Fix: the repository interface should accept and return only domain types. The *sql.DB connection is injected at the concrete implementation level.

## In Production

The repository pattern is the standard data access pattern in Go production services. Kubernetes uses repository-like interfaces for etcd storage (pkg/registry/). CockroachDB uses a repository abstraction for its SQL layer. Standard Go ORMs like sqlx, pgx, and GORM all support the repository pattern implicitly — the ORM query builder is the repository internals, and the application wraps it in a domain-typed interface. In production Go, every database access goes through a repository layer.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-05`.
