# Sessions and cookies

## Mission

Understand and apply Sessions and cookies in the context of professional Go software engineering.

## Prerequisites

- core-10-05

## Mental Model

A session is a temporary association between a user and server-side state. The session ID is a key to a lockbox on the server: the client has the key, but the lockbox contents never leave the server.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Sessions in Go are typically cookie-based: the server sets a cookie with the session ID on login, the browser sends it on every request, and the server looks up session data from a backing store (memory, Redis, DB). The session ID must be cryptographically random to prevent guessing.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/06-sessions-and-cookies
go test ./curriculum/modules/10-auth-security/lessons/06-sessions-and-cookies
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Storing sessions only in memory with no persistence — all sessions are lost on restart.
- Using predictable session IDs — session tokens must be cryptographically random.
- Not setting HttpOnly and Secure flags on session cookies — exposes session ID to XSS and man-in-the-middle.

## In Production

Sessions are used in web applications, admin panels, and any service where users authenticate via browser. Go services use Redis-backed sessions for horizontal scaling.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-07`.
