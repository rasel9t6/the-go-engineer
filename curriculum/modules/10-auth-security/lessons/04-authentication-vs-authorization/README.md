# Authentication vs authorization

## Mission

Understand and apply Authentication vs authorization in the context of professional Go software engineering.

## Prerequisites

- core-10-03

## Mental Model

Authentication answers 'who are you?' Authorization answers 'are you allowed to do that?' First is about identity proof; second is about permissions. Separate middleware: if auth fails -> 401, if authz fails -> 403.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Auth middleware validates credentials and stores identity in context. Authz middleware reads identity and compares against resource ownership or roles. Go's mux composes these as separate wrappers.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/04-authentication-vs-authorization
go test ./curriculum/modules/10-auth-security/lessons/04-authentication-vs-authorization
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing authentication (who you are) with authorization (what you can do).
- Implementing authorization as an afterthought in handlers instead of a separate middleware layer.
- Using the same middleware for both auth and authz — mixing them makes either hard to test independently.

## In Production

Production Go services separate auth and authz into distinct middleware. Auth middleware (JWT verification) runs first. Authz middleware (RBAC/ABAC) runs second.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-05`.
