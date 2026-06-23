# JWT implementation

## Mission

Understand and apply JWT implementation in the context of professional Go software engineering.

## Prerequisites

- core-10-06

## Mental Model

A JWT is a self-contained credential: the server signs a payload, and any service with the public key can verify it without contacting the issuing server. It is like a passport — it contains identity information, is tamper-evident (signature), and expires.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A JWT is three base64url-encoded segments separated by dots: header.payload.signature. The header contains the algorithm and type. The payload contains claims. The signature is the cryptographic output of signing header + payload with the private key. Verification recomputes the signature with the public key and compares.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/07-jwt-implementation
go test ./curriculum/modules/10-auth-security/lessons/07-jwt-implementation
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Not validating the algorithm header — accepting 'alg: none' tokens bypasses verification entirely.
- Using symmetric signing (HS256) for multi-service architectures — the shared secret must be distributed to every service, increasing exposure.
- Not setting expiration (exp claim) — tokens that never expire cannot be revoked and remain valid indefinitely.

## In Production

JWTs are used in: OAuth2 access tokens, OIDC ID tokens, service-to-service auth, and API authentication. Major Go services use golang-jwt/jwt for JWT handling.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-08`.
