# Release Guide

This document describes how to plan, prepare, and publish releases of The Go Engineer.

## Overview

The Go Engineer follows semantic versioning for stable releases:

- **Major**: significant curriculum restructuring
- **Minor**: new lessons, features, or substantial improvements without breaking the locked architecture
- **Patch**: bug fixes, documentation corrections, CI improvements, and dependency updates

Current stable line:

```text
v2.1.x
```

## Branch Roles

- `main`: post-v2.1 implementation line
- `release/v1`: stable v1 maintenance line
- `release/v2`: stable v2.1.x maintenance line

Topic branches stay short-lived and branch from the line they should ship to.

## Commit Convention

Use the bracketed commit format:

```text
[TYPE] short imperative description
```

Allowed types:

```text
[FEAT] [FIX] [DOCS] [TEST] [CHORE] [REFACTOR] [SECURITY] [RELEASE]
```

Examples:

```text
[RELEASE] prepare v2.1.1 metadata
[FIX] correct DB.3 tenant-scoping example
[DOCS] update release guide for v2.1.x maintenance
```

## Release Lines

### v2.1.x maintenance

Use `release/v2`.

Typical work:

- bug fixes
- documentation corrections
- low-risk lesson corrections
- CI and validator fixes
- dependency and security updates

### Post-v2.1 implementation

Use `main`.

Typical work:

- Opslane flagship implementation
- deeper engineering content
- validator improvements
- new proof surfaces inside existing sections

Do not use `main` as a replacement for stable-line maintenance when the fix should ship to v2.1.x users.

## Release Process

### Step 1: Plan the release

1. Identify issues and PRs to include.
2. Confirm target branch: `release/v1`, `release/v2`, or `main`.
3. Update [CHANGELOG.md](./CHANGELOG.md).
4. Update version references in docs if needed.

### Step 2: Create release-prep branch

For a v2.1.x release:

```bash
git switch release/v2
git pull origin release/v2
git switch -c release/v2.1.x-prep
```

### Step 3: Update release metadata

Update:

- `CHANGELOG.md`
- `README.md` if the featured status changes
- `ROADMAP.md` if maintenance or implementation priorities change

Commit with:

```bash
git commit -m "[RELEASE] prepare v2.1.x metadata"
```

### Step 4: Verify locally

Run:

```bash
go build ./...
go vet ./...
unformatted=$(gofmt -l .); test -z "$unformatted" || (echo "$unformatted" && exit 1)
go mod tidy
git diff --exit-code -- go.mod go.sum
go test ./...
go test -race ./...
go test -coverprofile=coverage.out ./...
go run ./scripts/validate_curriculum.go
```

For benchmark-related releases:

```bash
go test -bench=. -benchmem -count=1 ./08-quality-test/01-quality-and-performance/testing/benchmarks/
```

### Step 5: Open release PR

Open a PR into the target long-lived branch.

PR title:

```text
[RELEASE] prepare v2.1.x
```

PR body:

```markdown
## Release

Version: v2.1.x
Target branch: release/v2

## Summary

-

## Included changes

-

## Validation

- [ ] go build ./...
- [ ] go vet ./...
- [ ] gofmt check
- [ ] go mod tidy check
- [ ] go test ./...
- [ ] go test -race ./...
- [ ] go run ./scripts/validate_curriculum.go

## Risk

-
```

### Step 6: Merge release PR

Use **Squash and Merge** after approval and green CI.

### Step 7: Tag the release

```bash
git switch release/v2
git pull origin release/v2
git tag v2.1.x
git push origin v2.1.x
```

### Step 8: Create GitHub release

Create a GitHub release with:

- tag: `v2.1.x`
- title: `The Go Engineer v2.1.x`
- description copied from `CHANGELOG.md`

## Backports

If a fix belongs in multiple supported lines:

1. Merge once into the correct source branch.
2. Use `git cherry-pick -x` to propagate it.
3. Open a follow-up PR if branch protection requires it.

```bash
git switch <target-branch>
git pull origin <target-branch>
git cherry-pick -x <merged-commit-sha>
git push origin HEAD
```

## Security Updates

For security issues:

1. Use `[SECURITY]` in the issue and commit title.
2. Apply the smallest safe fix.
3. Add tests where possible.
4. Release a patch version.
5. Document the change in `CHANGELOG.md`.

## Rollback Procedure

If a release needs to be withdrawn before broad use:

```bash
git tag -d v2.1.x
git push origin --delete v2.1.x
```

If already merged and published, prefer a revert followed by a patch release.

## Release Checklist

- [ ] All included issues are closed or linked.
- [ ] All included PRs are merged.
- [ ] `CHANGELOG.md` is updated.
- [ ] `ROADMAP.md` is accurate.
- [ ] CI passes.
- [ ] No release-line mismatch.
- [ ] No architecture drift.
- [ ] Any backport need is labeled and tracked.
- [ ] GitHub release notes are published.
