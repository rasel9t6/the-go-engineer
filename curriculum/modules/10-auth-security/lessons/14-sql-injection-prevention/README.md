# SQL injection prevention

## Mission

Understand and apply SQL injection prevention in the context of professional Go software engineering.

## Prerequisites

- core-10-13

## Mental Model

SQL injection is a code injection attack where malicious SQL is smuggled inside a data value. Parameterized queries separate SQL code from data: the query template is parsed first (definition phase), then the parameters are substituted as data only, never as executable SQL.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

When a parameterized query is sent, the database receives the SQL template and the parameters as separate messages. The database parses the SQL template, builds a query plan, and then binds the parameters as data values. Since parameters are bound after parsing, they cannot alter the query structure.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/14-sql-injection-prevention
go test ./curriculum/modules/10-auth-security/lessons/14-sql-injection-prevention
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Concatenating user input into SQL strings: db.Query(fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", input)).
- Using string interpolation for IN clauses: strings.Join(userIDs, ",") instead of parameterized query markers.
- Believing that input validation (e.g., stripping quotes) prevents SQL injection — it does not.
- Only sanitizing GET parameters but not POST/PUT body fields or headers.

## In Production

Every Go database application uses parameterized queries. sqlx, pgx, GORM, and ent all internally use parameterized queries. The Go standard library database/sql enforces the parameterized pattern.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-15`.
