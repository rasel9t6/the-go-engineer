# CORS

## Mission

Understand and apply CORS in the context of professional Go software engineering.

## Prerequisites

- core-10-16

## Mental Model

CORS is a browser-enforced access control policy for cross-origin requests. It is not a server-side security measure — a malicious client (curl, Postman) ignores CORS entirely. CORS protects the browser's same-origin policy from being bypassed by JavaScript.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

CORS headers are: Access-Control-Allow-Origin (which origins can read the response), Access-Control-Allow-Methods (allowed HTTP methods), Access-Control-Allow-Headers (allowed custom headers), and Access-Control-Allow-Credentials (whether cookies/auth headers are allowed). The browser enforces these headers; the server just sets them.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/17-cors
go test ./curriculum/modules/10-auth-security/lessons/17-cors
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Setting Access-Control-Allow-Origin: * with credentials: true — this is invalid and causes the browser to reject the response.
- Reflecting the Origin header in Access-Control-Allow-Origin without validation — allows any site to read the response.
- Not handling preflight OPTIONS requests — real requests fail silently in the browser.
- Using CORS as a security mechanism — CORS is a browser policy, not server-side security.

## In Production

Every public API that supports browser clients configures CORS. Cloud providers (AWS, Google Cloud), social APIs (Twitter, GitHub), and payment gateways (Stripe) all use CORS to control cross-origin access.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-18`.
