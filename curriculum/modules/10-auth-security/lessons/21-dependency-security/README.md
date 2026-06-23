# Dependency security

## Mission

Understand and apply Dependency security in the context of professional Go software engineering.

## Prerequisites

- core-10-20

## Mental Model

Dependencies are like building materials in construction. You need to verify they come from a trusted source and are not counterfeit. go.sum is your material manifest: it lists every piece and its fingerprint. govulncheck is the building inspector that checks for known defective materials.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

When go mod download fetches a module, it computes a cryptographic hash of the module's content and records it in go.sum. On subsequent downloads, Go verifies the hash matches. The checksum database (sum.golang.org) provides a global, transparent log of module hashes, making it detectable if a module's content changes.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/21-dependency-security
go test ./curriculum/modules/10-auth-security/lessons/21-dependency-security
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Running go mod vendor without reviewing the vendored code — compromised dependencies are invisible in go.sum alone.
- Using the @latest version of a dependency without pinning — a malicious update gets deployed automatically.
- Adding direct dependencies for functionality that Go's standard library already provides.
- Not running go mod verify after a security incident to confirm go.sum matches the module content.

## In Production

Software supply chain attacks are on the rise. The 2024 XZ Utils backdoor (CVE-2024-3094) and the 2018 event-stream incident highlighted the criticality. Go's module system with go.sum, the module proxy, and govulncheck provide defense-in-depth.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-22`.
