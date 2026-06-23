# Docker basics

## Mission

Understand and apply Docker basics in the context of professional Go software engineering.

## Prerequisites

- core-15-05

## Mental Model

A container is just a Linux process isolated by namespaces (PID, network, mount, UTS, IPC) and constrained by cgroups. The Dockerfile is a recipe that assembles a read-only filesystem snapshot (image) — each instruction adds a layer (an overlayfs diff). When you run a container, Docker overlays a writable layer on top and executes ENTRYPOINT/CMD as PID 1. For Go developers, the mental model simplifies: the binary + its layer = the running process. Everything else is packaging detail.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Docker is a client-server architecture. The `docker` CLI talks to `dockerd` (the daemon) via a REST API (by default on `/var/run/docker.sock`). When you run `docker build`, the daemon receives the build context (tarred), unpack it, and parses the Dockerfile instruction-by-instruction. Each instruction becomes a layer: a commit of the filesystem diff using overlayfs (or aufs, devicemapper, etc.). Layers are content-addressed and cached. When you `docker run` an image, the daemon creates a new mount namespace, overlays the image layers + a writable layer, sets up networking (bridge by default), then forks the ENTRYPOINT/CMD as PID 1 in that namespace. The Go binary runs as PID 1 — it handles all signals directly (a key reason Go and Docker work well together, since Go's runtime handles SIGTERM, SIGINT cleanly).

## Run Instructions

```bash
go run ./curriculum/modules/15-docker-cicd-deployment/lessons/06-docker-basics
go test ./curriculum/modules/15-docker-cicd-deployment/lessons/06-docker-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using `ubuntu:latest` as the base image instead of `golang:alpine` or `scratch`, producing a 800 MB+ image for a 15 MB Go binary.
- Running `go build` inside the container without `CGO_ENABLED=0` and `-ldflags="-s -w"`, resulting in a binary linked against glibc that crashes on `scratch` or `alpine`.
- Copying the entire source tree with `COPY . .` instead of using `.dockerignore` or a multi-stage `COPY --from=build`, bloating layers with cached Go modules and `.git` history.
- Using `CMD ./myapp` instead of the exec form `CMD ["./myapp"]`, causing signals (SIGTERM) to be swallowed by the shell wrapper and preventing graceful shutdown.
- Forgetting `--chmod` on COPY or running as root, leaving the container with a security profile that violates the principle of least privilege.

## In Production

Every major Go project uses Docker for delivery. Docker Hub lists official images for Go itself, as well as Go-based tools like Traefik, Caddy, Hugo, and Syncthing. In production, Go services are deployed as Docker containers on ECS, GKE, AKS, or EKS — often built via CI (GitHub Actions, GitLab CI) and pushed to a registry, then pulled by the orchestrator. `docker-compose.yml` is used for local development to spin up the Go API alongside Postgres, Redis, and a message broker.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-15-07`.
