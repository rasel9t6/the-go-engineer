# Architecture decision records

## Mission

Understand and apply Architecture decision records in the context of professional Go software engineering.

## Prerequisites

- core-14-16

## Mental Model

ADRs are commit messages for architecture decisions. A commit message explains WHY a code change was made. An ADR explains WHY an architecture decision was made. Without ADRs, the team has to ask 'why did we choose PostgreSQL?' and the answer is 'I don't know, it was before my time.' With ADRs, the answer is in the repo: read ADR-001. ADRs preserve context that would otherwise be lost when team members leave or memories fade.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

ADRs are plain Markdown files — no special tooling required. The standard format (Michael Nygard's template) has 5 sections: Title (ADR-XXX: Decision title), Status (Proposed, Accepted, Deprecated, Superseded), Context (why this decision is needed, what problem it solves, what constraints exist), Decision (the chosen option, with justification), Consequences (what becomes easier and harder as a result). ADRs are stored in a docs/adr/ directory in the repository. The filename convention is NNNN-title-with-hyphens.md. The first ADR (ADR-001) explains the ADR format itself. Tools like adr-tools and adr-log automate ADR management, but a text editor and git are sufficient.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/17-architecture-decision-records
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/17-architecture-decision-records
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Writing ADRs as design documents (10-page Word docs that no one reads) — ADRs must be short: title, status, context, decision, consequences. If it takes more than 5 minutes to write, the team will not write them. A good ADR is 5-10 sentences.
- Writing ADRs after the decision is implemented — by the time the code is written, the decision is already made and the context is lost. ADRs must be written BEFORE or DURING the decision process, capturing the alternatives considered and the reasoning.
- Not updating ADR status when the decision changes — the team decides to use PostgreSQL, writes ADR-001. Six months later they switch to CockroachDB, but ADR-001 still says 'status: accepted' and no one knows why they switched. ADRs must be living documents: superseded ADRs get status 'deprecated' and reference the new ADR.

## In Production

ADRs are standard practice in every mature engineering organization. ThoughtWorks pioneered the format and uses it in all client engagements. The Kubernetes project uses ADRs (called KEPs — Kubernetes Enhancement Proposals). The Go project itself uses proposal documents that follow the ADR format. ADRs are stored in the repository at docs/adr/ or docs/decisions/. The format is standardized: title, status, context, decision, consequences, compliance, notes.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
