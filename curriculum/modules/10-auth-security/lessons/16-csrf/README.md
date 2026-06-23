# CSRF

## Mission

Understand and apply CSRF in the context of professional Go software engineering.

## Prerequisites

- core-10-15

## Mental Model

CSRF is a confused deputy attack: the browser is tricked into making an authenticated request on behalf of the attacker. CSRF tokens act as a secret handshake: the server embeds a unique token in each form, and only requests carrying this token are accepted.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The double-submit cookie pattern: the server generates a random token, sets it as a cookie, and requires the same token in a custom request header. Since the attacker's page cannot read the cookie (same-origin policy) or the custom header (CORS), they cannot forge a valid request.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/16-csrf
go test ./curriculum/modules/10-auth-security/lessons/16-csrf
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Building APIs without CSRF protection, assuming JSON requests are not vulnerable — CORS misconfigurations enable cross-origin JSON requests.
- Using the wrong token pattern: storing CSRF tokens in cookies without the SameSite attribute.
- Implementing CSRF tokens manually instead of using a proven library.
- Disabling CSRF protection for authenticated endpoints.

## In Production

CSRF protection is mandatory for all browser-based applications that use cookie authentication. Major Go frameworks (Chi, Gin, Echo) include CSRF middleware. Modern browsers also support SameSite cookies as a defense-in-depth measure.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-17`.
