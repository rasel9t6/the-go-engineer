# Memory layout

## Mission

Understand and apply Memory layout in the context of professional Go software engineering.

## Prerequisites

- core-13-05

## Mental Model

Memory layout is like Tetris: each field is a block of a certain width (its size), and the compiler places them in a row (the struct). If a block cannot start at the current position because it would overlap a boundary (alignment), the compiler adds padding (empty space) before it. The row's total width is padded at the end to be a multiple of the widest block. Poor ordering wastes space (padding); optimal ordering fits all blocks with minimal gaps.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go uses the System V AMD64 ABI alignment rules on most platforms. Each type has a natural alignment equal to its size (e.g., int64 aligns to 8 bytes, int32 to 4, int16 to 2, bool/int8 to 1). The compiler processes struct fields in declaration order, maintaining a current offset. For each field, if offset % field.alignment != 0, it adds padding bytes until offset is aligned. After processing all fields, the compiler adds trailing padding to make the struct size a multiple of the maximum alignment among all fields. The resulting layout is visible via unsafe.Offsetof which returns the byte offset of each field from the struct's start. Go is a memory-safe language — it does not allow reading padding bytes, and the padding bytes are zero-initialized but their content is undefined after struct mutation. The compiler may reorder fields for optimization (Go 1.18+ attempts field reordering by size), but this does not change the safety invariants.

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/06-memory-layout
go test ./curriculum/modules/13-performance-memory-engineering/lessons/06-memory-layout
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Declaring struct fields in logical order instead of size order — field order determines struct size due to alignment padding. Sorting fields by size descending (largest first) minimizes padding. A struct with fields int64, int32, bool (8+4+1 = 13 bytes) can require 24 bytes if declared as bool, int32, int64 (1 + 3 padding + 4 + 4 padding + 8 = 24).
- Assuming sizeof(struct) = sum(sizeof(fields)) — alignment rules require each field to be at an address that is a multiple of its alignment (field size, up to 8 bytes on 64-bit). The compiler adds padding between fields to satisfy alignment, and at the end of the struct to ensure the next element in an array is aligned. The actual size includes invisible padding.
- Using bool fields in hot loops because 'bool is smaller' — bool is 1 byte but has alignment 1. Adjacent int32 fields require 3 bytes of padding. A bool in a struct accessed in a hot loop causes a cache line to hold only bool + padding + int32, wasting cache space. Using a bitmask (uint8) for multiple bools avoids the padding.

## In Production

Memory layout optimization is critical in high-performance Go: the standard library's sync.Map, net/http.Server, and runtime structures are carefully ordered to minimize size and avoid false sharing. DataDog's Go trace-agent reordered hotspot structs to fit within a 64-byte cache line. In Kubernetes, optimizing the Pod struct layout saved 8 bytes per pod — on a cluster with 10,000 pods, that is 80KB of saved heap (minor) but more importantly, better cache utilization for the scheduler's hot loop.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-13-07`.
