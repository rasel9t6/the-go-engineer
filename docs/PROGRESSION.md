# The Go Engineer - Visual Progression

## Overall Progress Model

```
┌──────────────────────────────────────────────────────────────────────────────┐
│                                                                              │
│   100% ──────────────────────────────────────────────────────────────────  │
│                                                                              │
│       ╔══════════════════════════════════════════════════════════════════╗  │
│       ║  PHASE 5: FLAGSHIP PROJECT (100%)                                 ║  │
│       ║  "I build complete production systems"                           ║  │
│       ║  Engineering Context: 100%                                        ║  │
│       ╚══════════════════════════════════════════════════════════════════╝  │
│                                    │                                         │
│       ╔══════════════════════════════════════════════════════════════════╗  │
│       ║  PHASE 4: SYSTEM ENGINEERING (85% → 100%)                        ║  │
│       ║  "I think like a senior engineer"                                 ║  │
│       ║  Engineering Context: 85-95%                                     ║  │
│       ║                                                                  ║  │
│       ║  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐  ║  │
│       ║  │HTTP     │ │Database │ │Production│ │Config   │ │Docker & │  ║  │
│       ║  │Servers  │ │Engineering│ │Engineering│ │& Env    │ │Deploy  │  ║  │
│       ║  │         │ │         │ │         │ │         │ │        │  ║  │
│       ║  └─────────┘ └─────────┘ └─────────┘ └─────────┘ └─────────┘  ║  │
│       ╚══════════════════════════════════════════════════════════════════╝  │
│                                    │                                         │
│       ╔══════════════════════════════════════════════════════════════════╗  │
│       ║  PHASE 3: PRODUCTION ENGINEERING (60% → 85%)                     ║  │
│       ║  "I write code that doesn't break at scale"                      ║  │
│       ║  Engineering Context: 70-80%                                     ║  │
│       ║                                                                  ║  │
│       ║  ┌──────────────┐ ┌──────────────┐ ┌──────────────┐ ┌──────────┐║  │
│       ║  │Concurrency   │ │Context &     │ │Advanced     │ │Testing   │║  │
│       ║  │Fundamentals  │ │Cancellation  │ │Concurrency  │ │Fundamen- │║  │
│       ║  │              │ │              │ │Patterns     │ │tals     │║  │
│       ║  └──────────────┘ └──────────────┘ └──────────────┘ └──────────┘║  │
│       ║                                                                  ║  │
│       ║  ┌───────────────────────────────────────────────────────────┐   ║  │
│       ║  │Section 12: Performance & Profiling                        │   ║  │
│       ║  └───────────────────────────────────────────────────────────┘   ║  │
│       ╚══════════════════════════════════════════════════════════════════╝  │
│                                    │                                         │
│       ╔══════════════════════════════════════════════════════════════════╗  │
│       ║  PHASE 2: ENGINEERING FOUNDATION (30% → 60%)                    ║  │
│       ║  "I write code that works reliably"                              ║  │
│       ║  Engineering Context: 40-55%                                   ║  │
│       ║                                                                  ║  │
│       ║  ┌──────────┐ ┌──────────────────┐ ┌──────────────────────┐    ║  │
│       ║  │Data      │ │Functions &       │ │Types &               │    ║  │
│       ║  │Structures│ │Errors ⭐         │ │Interfaces            │    ║  │
│       ║  │          │ │(transformed)     │ │                      │    ║  │
│       ║  └──────────┘ └──────────────────┘ └──────────────────────┘    ║  │
│       ║                                                                  ║  │
│       ║  ┌───────────────────────────────────────────────────────────┐    ║  │
│       ║  │Section 7: Packages & Modules                             │    ║  │
│       ║  └───────────────────────────────────────────────────────────┘    ║  │
│       ╚══════════════════════════════════════════════════════════════════╝  │
│                                    │                                         │
│       ╔══════════════════════════════════════════════════════════════════╗  │
│       ║  PHASE 1: SYNTAX FOUNDATION (0% → 30%)                         ║  │
│       ║  "I can write Go code"                                          ║  │
│       ║  Engineering Context: 10-30%                                   ║  │
│       ║                                                                  ║  │
│       ║  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────────────┐   ║  │
│       ║  │How       │ │Core     │ │Language │ │Control               │   ║  │
│       ║  │Computers │ │Foundations│ │Basics  │ │Flow                  │   ║  │
│       ║  │Work      │ │         │ │         │ │                      │   ║  │
│       ║  └──────────┘ └──────────┘ └──────────┘ └──────────────────────┘   ║  │
│       ╚══════════════════════════════════════════════════════════════════╝  │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘

    0%                              30%      40%          60%      70%         100%
    └───────────────────────────────┴────────┴────────────┴─────────┴──────────┘
    Syntax                                 Engineering                         Senior
    Foundation                            Foundation                          Engineering
```

---

## Learning Path: Zero to Senior

### Path A: Complete Beginner

```
START
  │
  ▼
[0] How Computers Work ──→ [1] Core Foundations ──→ [2] Language Basics ──→ [3] Control Flow
  (0%)                     (5%)                    (10%)                    (15%)
  │
  ▼
[4] Data Structures ──→ [5] Functions & Errors ──→ [6] Types & Interfaces ──→ [7] Packages
  (20%)                    (30%)                    (40%)                     (50%)
  │
  ▼
[8] Concurrency ──→ [9] Context ──→ [10] Advanced Patterns ──→ [11] Testing ──→ [12] Performance
  (55%)          (60%)          (65%)                    (70%)           (75%)
  │
  ▼
[13] HTTP ──→ [14] Database ──→ [15] Production ──→ [16] Config ──→ [17] Docker ──→ FLAGSHIP
  (80%)         (82%)           (85%)           (88%)           (90%)          (100%)
```

---

## Engineering Context Growth

```
Engineering Context %

100% ┤
     │                                              ████ FLAGSHIP ████
 90% ┤                                         ████
     │                                    ████
 80% ┤                               ████
     │                          █████
 70% ┤                     █████
     │                █████
 60% ┤            █████                  ●●●● PHASE 3
     │       █████                      ●●●●
 50% ┤  ●●●●                       ●●●●
     │  ●●●●                    ●●●●
 40% ┤  ●●●●              ●●●●
     │  ●●●●         ●●●●
 30% ┤  ●●●●    ●●●●
     │  ●●●●●●●●●●●
 20% ┤  ●●●●●●●●●●●●  ●●●● PHASE 1
     │                 ●●●●
 10% ┤
     │
  0% └──────────────────────────────────────────────────────────────
         Section 0-3     Section 4-7     Section 8-12    Section 13-17

     ● = Syntax + Basic Why
     █ = Full Engineering Context
```

---

## Key Milestones

| Milestone             | Section  | Engineering Checkpoint                | % Complete |
| --------------------- | -------- | ------------------------------------- | ---------- |
| First Program         | 1        | Run and modify hello world            | 5%         |
| Type Converter        | LB.4     | Unit test with multiple cases         | 10%        |
| FizzBuzz              | CF.5     | Handle all cases cleanly              | 15%        |
| Contact Directory     | DS.6     | Use slices + maps + pointers together | 20%        |
| Order Processing      | FE.7 ⭐  | Handle security, scale, overflow      | 30%        |
| Custom Error Types    | TI.6     | Implement error interface             | 40%        |
| Concurrent Downloader | GC.6     | Handle race conditions                | 55%        |
| Health Checker        | CP.5     | Debug concurrent failures             | 65%        |
| Benchmark Suite       | TE.4     | Profile and optimize                  | 70%        |
| REST API              | HS.7     | Handle timeouts, middleware           | 80%        |
| Database CRUD         | DB.6     | Handle transactions                   | 82%        |
| Graceful Shutdown     | GS.3     | Deploy with zero downtime             | 85%        |
| Dockerized Service    | DOCKER   | Deploy to container                   | 90%        |
| **Flagship Complete** | FLAGSHIP | Full production system                | 100%       |

---

## Core Engineering Principles Per Phase

### Phase 1: Foundation

- **Core Idea**: Code is instructions for a computer
- **Key Insight**: Variables store data in memory
- **Engineering Mind**: "What does the computer actually do?"

### Phase 2: Reliability

- **Core Idea**: Functions are boundaries, errors are values
- **Key Insight**: Fail fast, validate early, handle errors explicitly
- **Engineering Mind**: "What happens when things go wrong?"

### Phase 3: Scalability

- **Core Idea**: Concurrency is about coordination
- **Key Insight**: Race conditions, deadlocks, goroutine leaks
- **Engineering Mind**: "What breaks when 1000 users use this?"

### Phase 4: Production

- **Core Idea**: Systems are observed, not guessed
- **Key Insight**: Logs, metrics, tracing, graceful shutdown
- **Engineering Mind**: "How do I debug at 3 AM?"

### Phase 5: Flagship

- **Core Idea**: Everything comes together
- **Key Insight**: Real-world complexity, trade-offs, decisions
- **Engineering Mind**: "What would a senior engineer do?"

---

## The Promise

```
By completing this curriculum, you will be able to:

  ✓ Write Go code from scratch
  ✓ Structure code for maintainability
  ✓ Handle errors properly
  ✓ Write concurrent code safely
  ✓ Test code comprehensively
  ✓ Profile and optimize performance
  ✓ Build production HTTP services
  ✓ Work with databases reliably
  ✓ Deploy and operate systems
  ✓ Think like a senior engineer
```

---

_This is a living document. The curriculum will evolve as sections are transformed._
