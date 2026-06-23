# XSS

## Mission

Understand and apply XSS in the context of professional Go software engineering.

## Prerequisites

- core-10-14

## Mental Model

XSS is an injection attack where malicious scripts are injected into web pages viewed by other users. html/template acts as an automatic output sanitizer: it encodes special characters (<>"'&) based on the HTML context, ensuring user data is treated as text, not code.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

html/template parses a template into a tree of nodes. When executing, it applies context-sensitive escaping rules. In an HTML body context, it escapes <>&"'. In a JavaScript string context, it uses JavaScript string escaping. The package prevents most XSS, but template.HTML (raw) bypasses escaping entirely.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/15-xss
go test ./curriculum/modules/10-auth-security/lessons/15-xss
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Rendering user-generated HTML without escaping: template.HTML(userInput) creates a direct XSS vector.
- Sanitizing input but not output — input sanitization is fragile and bypassable.
- Using html/template for user-generated content without understanding context-sensitive escaping.
- Disabling CSP headers because they block legitimate scripts.

## In Production

XSS consistently ranks in the OWASP Top 10. Go's html/template is the standard defense for server-rendered HTML. For JSON APIs, the frontend framework is responsible for safe rendering.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-16`.
