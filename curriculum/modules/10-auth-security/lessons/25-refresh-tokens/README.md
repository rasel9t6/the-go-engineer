# Refresh tokens

## Mission

Implement a secure refresh token rotation flow in Go.

## Prerequisites

- core-10-24

## Mental Model

Access tokens are hotel room keys that expire at checkout. Refresh tokens are the front desk: you show your ID to get a new room key. If someone steals your room key, it stops working at checkout. If someone steals your ID (refresh token), rotating the lock renders theirs useless.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The OAuth 2.0 refresh token grant type (RFC 6749 section 6) defines the protocol. The rotation pattern is specified in RFC 6819 (OAuth 2.0 Threat Model) and OAuth 2.0 Security BCP (draft). In Go, database/sql's Tx provides the atomicity needed for safe rotation. The crypto/rand package provides cryptographically secure token generation — math/rand must never be used for tokens.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/25-refresh-tokens
go test ./curriculum/modules/10-auth-security/lessons/25-refresh-tokens
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Storing refresh tokens in the same database as access tokens without isolation.
- Not rotating refresh tokens on each use, allowing unlimited reuse of stolen tokens.
- Using opaque bearer tokens without a revocation check, making logout impossible.

## In Production

Every production authentication system uses refresh tokens. GitHub issues access tokens (1hr expiry) and refresh tokens (6mo rotation). Auth0, Keycloak, and Firebase Auth all implement the OAuth2 refresh token grant type. Go microservices behind an API gateway use refresh tokens to maintain session state without forcing regular re-login.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
