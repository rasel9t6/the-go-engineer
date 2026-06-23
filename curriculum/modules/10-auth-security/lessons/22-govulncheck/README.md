# govulncheck

## Mission

Understand and apply govulncheck in the context of professional Go software engineering.

## Prerequisites

- core-10-21

## Mental Model

govulncheck is a security radar for your Go project. It continuously scans the dependency horizon for known threats and alerts you before they become problems. It filters out noise: if a vulnerability exists in a module but your code never calls the vulnerable function, govulncheck does not report it.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

govulncheck downloads the Go vulnerability database (OSV format) from vuln.go.dev. It builds the import graph of the project, matches module versions against known vulnerabilities, and then performs call graph analysis to determine if the vulnerable function is actually reachable from the project's code. Only reachable vulnerabilities are reported.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/22-govulncheck
go test ./curriculum/modules/10-auth-security/lessons/22-govulncheck
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Running govulncheck only once during initial setup instead of integrating it into CI/CD.
- Ignoring govulncheck findings because they are 'only in test dependencies' — test code imports production packages as transitive deps.
- Not updating the vuln database regularly — stale vulnerability data misses newly published CVEs.
- Fixing the wrong package: updating the direct dependency when the CVE is in a transitive dependency.

## In Production

govulncheck is the Go team's official vulnerability scanning tool. It is integrated into the Go VS Code extension (gopls), GitHub Actions (golangci-lint includes it), and CI pipelines for major Go projects including Kubernetes and Docker.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-23`.
