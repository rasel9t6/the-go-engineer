# Protobuf

## Mission

Understand and apply Protobuf in the context of professional Go software engineering.

## Prerequisites

- elective-08

## Mental Model

Protobuf is a contract-first serialization format. You define the message structure in a .proto file (the contract), and the protoc compiler generates Go structs with marshal/unmarshal methods. Producers and consumers share the contract — they do not share code or runtime dependencies. Fields are numbered (field 1, field 2) instead of named — this enables backward-compatible field addition and removal.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Protobuf encoding uses a binary wire format with varint encoding for integers, length-delimited strings/bytes, and fixed-width encoding for floats/doubles. Each field is identified by its field number and wire type (varint, 64-bit, length-delimited, start-group, 32-bit). The Go protoc plugin (protoc-gen-go) generates Go structs with protobuf struct tags and Marshal/Unmarshal methods that use the google.golang.org/protobuf/proto package for encoding and decoding.

## Run Instructions

```bash
Read the lesson and complete the practice task.
No automated test is required for this lesson.
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Reusing field numbers after removing a field — the old field number is reserved and should not be reused (use reserved keyword in .proto).
- Changing field types after deployment — a int32 field cannot be changed to string; use a new field number and deprecate the old one.
- Not using proto3 syntax — proto2 supports required fields and custom defaults, which complicate backward compatibility.
- Storing protobuf binary data in a text-based format (JSON database) — protobuf binary is opaque and cannot be searched or indexed.
- Importing .proto files with relative paths that break when the project structure changes — use protoc -I with absolute or module-relative import paths.

## In Production

Protobuf is used in every gRPC service and many event-driven systems. Google uses protobuf for all internal RPC. Go microservices use protobuf for service-to-service communication, event schemas in message queues (Kafka, NATS), and configuration file formats. Protobuf is the standard schema format for Kubernetes API resources.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-10`.
