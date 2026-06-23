# OWASP for Go APIs

## Mission

Understand and apply OWASP for Go APIs in the context of professional Go software engineering.

## Prerequisites

- core-10-22

## Mental Model

OWASP is the periodic table of web security risks. It categorizes the elements of insecurity: injection, broken auth, data exposure, etc. Each element has known properties, detection methods, and remediation. The OWASP API Security Top 10 is the subset most relevant to Go API developers.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

OWASP API Security Top 10 is developed by the OWASP community based on real-world incident data. Each category includes: description (what the risk is), example attack scenario, prevention measures, and references. The Top 10 is updated every 3–4 years as the threat landscape evolves.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/23-owasp-for-go-apis
go test ./curriculum/modules/10-auth-security/lessons/23-owasp-for-go-apis
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming OWASP is only relevant for traditional web applications, not JSON APIs — injection attacks work on any input channel.
- Implementing security controls in isolation without referencing a threat model.
- Focusing only on authentication while ignoring authorization, rate limiting, and input validation.
- Treating OWASP recommendations as a checklist instead of a risk-based framework.

## In Production

OWASP's Top 10 lists are the de facto standard for web security. Compliance frameworks (PCI DSS, SOC 2) reference OWASP. Major companies use OWASP as the foundation for their secure development lifecycle. Go projects with security requirements use gosec for OWASP-relevant static analysis.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-24`.
