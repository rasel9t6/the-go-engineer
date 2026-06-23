# File permissions

## Mission

Understand and apply File permissions in the context of professional Go software engineering.

## Prerequisites

- core-15-03

## Mental Model

Each ring has a three-bit mask (rwx). Octal notation encodes all three rings in one number: owner-group-other. The sticky bit on /tmp means only the owner of a file can delete it. Setuid means run this with the file owner's privileges regardless of who executes it. Go's FileMode is the language's direct port of this Unix model — when you see os.FileMode(0644), read it as owner writes and reads, everyone else only reads.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's os.FileMode is a uint32 where the lower 12 bits mirror the Unix permission bitmap: bits 0-2 are other rwx, bits 3-5 are group rwx, bits 6-8 are owner rwx, bit 9 is sticky (S_ISVTX), bit 10 is setgid (S_ISGID), bit 11 is setuid (S_ISUID). The os.Stat syscall on Linux fills a syscall.Stat_t struct whose Mode field is uint32 — Go casts it to os.FileMode. On Windows, os.FileMode still exists but the permission bits (owner/group/other) are largely ignored; Go maps the Windows file attributes to a simplified mode where regular files get 0666 and directories get 0777. The os.Chmod on Windows translates to SetFileAttributes. Docker layers each add their own metadata overlay: when you RUN chmod in a Dockerfile, the overlay filesystem's metadata layer records the new mode without rewriting the underlying file data. When os.Chmod is called inside a running container, it operates on the overlay's copy-on-write layer. Understanding that umask is a kernel-level process attribute (set per-process with syscall.Umask) explains why two Go programs in the same container can see different effective permissions — umask is inherited across fork/exec.

## Run Instructions

```bash
go run ./curriculum/modules/15-docker-cicd-deployment/lessons/04-file-permissions
go test ./curriculum/modules/15-docker-cicd-deployment/lessons/04-file-permissions
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using chmod 0777 as a default to make things work, which bypasses all permission controls and creates security vulnerabilities in production.
- Confusing rwx (read/write/execute) octal notation — writing chmod 644 instead of chmod 0644 for a script file, forgetting the execute bit, or using 0777 for directories thinking it mirrors files.
- Running containers as root without setting USER in the Dockerfile, causing all created files and logs to be owned by root and inaccessible to non-root CI/CD or host processes.
- Assuming os.WriteFile defaults to a secure permission — it uses 0666 (before umask) by default unless you explicitly pass os.FileMode(0644), which means secrets or config files can be world-readable.
- Forgetting that umask subtracts from the mode passed to os.OpenFile or os.WriteFile, then wondering why the resulting file has different permissions than the octal value they passed.

## In Production

Go production services use os.FileMode extensively: HTTP servers restrict TLS key files to 0600, CI/CD pipelines enforce permission audits in pull requests, Kubernetes init containers chown volumes to non-root UIDs, logging libraries open log files with 0644 (or 0600 for audit logs), and os.Stat checks gate access to config directories in multi-tenant API servers. The Go linter gosec (G301/G302/G306) flags missing or overly-permissive FileMode arguments in os.WriteFile, os.OpenFile, and os.Chmod calls — preventing 0777 or 0644 on secrets before they reach production.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-15-05`.
