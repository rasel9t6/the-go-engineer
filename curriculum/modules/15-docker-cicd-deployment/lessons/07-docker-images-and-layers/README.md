# Docker images and layers

## Mission

Understand and apply Docker images and layers in the context of professional Go software engineering.

## Prerequisites

- core-15-06

## Mental Model

When a process reads a file, Docker's union filesystem (overlayfs) walks the stack top-to-bottom and returns the first version it finds. Deleting a file in a later overlay does not erase it from earlier overlays — it just places a 'hidden' marker on top. This mental model explains why a 800MB build stage can yield an 8MB final image, why deleting intermediate files mid-build does nothing to shrink the image, and why reordering instructions changes cache behavior.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Docker images are stored under /var/lib/docker/overlay2 (Linux) as a set of read-only snapshots, each corresponding to one Dockerfile instruction. Each layer is a tarball of filesystem diffs. The union mount — overlayfs — combines all layers into a merged view via a lowerdir (image layers) and upperdir (container writable layer). When a container writes a file, the copy-on-write mechanism copies the file from a lower layer to the upper layer. docker history decompiles an image into its instruction list with per-layer sizes built from the diff metadata. The --squash flag merges all layers into one, trading cache granularity for a flatter, smaller image at the cost of losing the ability to share intermediate layers across images.

## Run Instructions

```bash
go run ./curriculum/modules/15-docker-cicd-deployment/lessons/07-docker-images-and-layers
go test ./curriculum/modules/15-docker-cicd-deployment/lessons/07-docker-images-and-layers
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Putting RUN go build and COPY source in the same layer, defeating layer caching and forcing a full rebuild on every source change.
- Using alpine as the final base image when a scratch image would suffice, ignoring that Go static binaries need nothing from the OS.
- Not ordering Dockerfile instructions from least-frequently-changing to most-frequently-changing, causing unnecessary cache invalidation on every build.
- Copying the entire build context (including .git, node_modules, local binaries) into the image, inflating both build time and final image size.
- Believing distroless images are secure by default — they still need non-root users and proper capabilities.

## In Production

Every major Go project that ships as a Docker image uses multi-stage builds with scratch or distroless. Docker CLI, Kubernetes (kubelet, kube-apiserver, etcd), Prometheus, Grafana, Caddy, Traefik, and HashiCorp tools (Consul, Vault, Nomad, Terraform) all follow this pattern. The Go Dockerfile is so universally clean that Docker's official documentation uses it as the canonical multi-stage example.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-15-08`.
