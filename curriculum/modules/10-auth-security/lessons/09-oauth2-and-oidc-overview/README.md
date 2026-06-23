# OAuth2 and OIDC overview

## Mission

Understand and apply OAuth2 and OIDC overview in the context of professional Go software engineering.

## Prerequisites

- core-10-08

## Mental Model

OAuth2 is a valet key system for APIs. The user gives the client a limited-access token (valet key) instead of their password. OIDC adds an identity layer: the client also receives a signed ID card (id_token) proving who the user is.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

OAuth2 defines four roles: resource owner, client, authorization server, and resource server. The authorization code flow uses a one-time code that the client exchanges for tokens, preventing the token from being exposed in the redirect URL. OIDC extends OAuth2 with the id_token (JWT) and UserInfo endpoint.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/09-oauth2-and-oidc-overview
go test ./curriculum/modules/10-auth-security/lessons/09-oauth2-and-oidc-overview
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Implementing OAuth2 from scratch instead of using a proven library — the specification has many edge cases.
- Confusing OAuth2 (authorization delegation) with OIDC (authentication) — OAuth2 alone does not verify user identity.
- Not validating the id_token signature and claims in OIDC — accepting unverified tokens allows impersonation.
- Storing client secrets in client-side code — public clients cannot keep secrets.

## In Production

OAuth2 and OIDC are the foundation of modern identity federation. Google, GitHub, Microsoft, and Auth0 all use OAuth2/OIDC. Go services integrate with these providers via golang.org/x/oauth2.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-10`.
