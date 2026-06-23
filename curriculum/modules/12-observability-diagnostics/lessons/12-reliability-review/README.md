# Reliability review

## Mission

Understand and apply Reliability review in the context of professional Go software engineering.

## Prerequisites

- core-12-11

## Mental Model

A reliability review is like an airplane crash investigation. The goal is not to punish the pilot (blame) but to understand why the crash happened and prevent it from happening again. The investigation looks at: the pilot's actions (the immediate trigger), the plane's design (the system), the airline's training program (the process), and the regulatory framework (the culture). Fixing the pilot (firing the engineer) does not fix the design, training, or culture. A good review identifies fixes at ALL levels: code, deployment, testing, monitoring, and process.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A reliability review is a structured document, not a technical system. The structure follows Google's SRE post-incident review format: (1) Incident ID and severity, (2) Summary (1 paragraph), (3) Impact (users affected, duration, revenue loss), (4) Timeline (chronological events with timestamps), (5) Root cause analysis (5 Whys technique), (6) Systemic gaps (process, testing, monitoring, deployment), (7) Action items (what, who, when), (8) Blamelessness statement (explicitly stating no individual is blamed). The review is stored in a shared drive or wiki. Action items are tracked in the team's project management system with recurring follow-ups. The review process is iterative: after 3 months, the team reviews whether the action items prevented similar incidents.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/12-reliability-review
go test ./curriculum/modules/12-observability-diagnostics/lessons/12-reliability-review
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Treating the post-incident review as a blame exercise — writing 'engineer X made a mistake' as the root cause. The goal is to improve the system, not assign blame. A better root cause: 'the deployment pipeline did not have a canary stage to detect the regression before production.' Blame-free reviews encourage engineers to report incidents honestly.
- Fixing the symptom without addressing the systemic gap — after an incident, patching the specific bug (fixing one missing validation) without addressing why the bug reached production (no code review for that path, no test for that edge case). The incident will repeat with a different bug. Fix: identify the systemic gap (testing, review, monitoring) and close it.
- Not documenting action items with owners and deadlines — a review produces a list of fixes but no one is assigned to implement them, and there is no follow-up. The fixes are forgotten until the next incident. Every action item must have: owner, deadline, and a verification step.

## In Production

Post-incident reviews are standard practice in every reliable software organization. Google's SRE team reviews every SEV-1/SEV-2 incident. GitHub publishes post-incident reviews publicly. Cloudflare's incident review process is documented in their blog. In Go production services, the review pattern is applied to: deployment failures, data loss incidents, security breaches, and capacity-related outages. The review document becomes part of the team's knowledge base and is referenced in future incident training.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
