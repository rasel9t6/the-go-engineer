# SQL injection prevention

## Mission

Understand and apply SQL injection prevention in the context of professional Go software engineering.

## Prerequisites

- core-10-13

## Mental Model

SQL injection occurs when untrusted user input is concatenated directly into SQL query strings, allowing an attacker to execute arbitrary SQL commands against the database.

## Visual Model

```text
input -> query construction -> database execution -> safe/unsafe result
```

## Machine View

Parameterized queries separate SQL logic from data values. The database server receives the query template and the parameters separately, preventing malicious input from being interpreted as SQL commands.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/14-sql-injection-prevention
go test ./curriculum/modules/10-auth-security/lessons/14-sql-injection-prevention
```

## Try It

Write a query that uses parameterized query markers instead of string interpolation.

## In Production

Production systems must never use string interpolation for SQL queries that contain user-supplied values.

## Thinking Questions

- Why can't input validation alone prevent SQL injection?
- What happens to the query plan when parameters are used?

## Next Step

XSS
