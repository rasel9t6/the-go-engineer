# JWT risks

## Mission

Understand and apply JWT risks in the context of professional Go software engineering.

## Prerequisites

- core-10-07

## Mental Model

A JWT is a tamper-evident credential container. The server signs a JSON payload, producing a compact string any service can verify with the public key.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A JWT is three base64url-encoded segments separated by dots: header.payload.signature. The header specifies the algorithm and token type. The payload contains registered claims (iss, sub, exp) and custom claims. The signature is the cryptographic output of signing header+payload with the private key. Verification recomputes the signature using the public key and compares.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/08-jwt-risks
go test ./curriculum/modules/10-auth-security/lessons/08-jwt-risks
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
- Using HS256 in multi-service architectures — every service must hold the shared secret, increasing exposure.
- Failing to validate the exp claim — tokens without expiration remain valid indefinitely if leaked.
- Storing sensitive data in JWT claims — the payload is base64-encoded, not encrypted.

## In Production

JWTs are the foundation of OAuth2 access tokens, OIDC ID tokens, and service-to-service authentication. Major Go projects use golang-jwt/jwt for token handling in production.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-09`.
