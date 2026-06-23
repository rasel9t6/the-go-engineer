# API keys

## Mission

Understand and apply API keys in the context of professional Go software engineering.

## Prerequisites

- core-10-09

## Mental Model

An API key is a shared secret between the client and the server. It identifies the client and is used for authentication, not authorization.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

An API key is a cryptographically random string, typically 32–64 bytes encoded as hex or base64url. The server stores only a hash of the key, so a database breach does not expose active keys. Key validation involves hashing the presented key and looking up the hash.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/10-api-keys
go test ./curriculum/modules/10-auth-security/lessons/10-api-keys
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using API keys as the sole authentication mechanism for human users — API keys lack granular revocation and rotation.
- Logging API keys in request logs or error messages — keys in logs are accessible to anyone with log access.
- Storing API keys in plaintext in the database — a database breach exposes all keys.
- Generating predictable API keys — sequential or time-based keys are guessable.

## In Production

API keys are used by Stripe, Twilio, GitHub, and most SaaS platforms. A Go API gateway typically validates API keys as the first middleware in the chain before routing to backend services.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-11`.
