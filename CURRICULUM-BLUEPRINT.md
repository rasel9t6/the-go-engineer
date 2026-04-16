# The Go Engineer - Complete Curriculum Blueprint (v2.0)

## Vision

A progressive engineering learning system that takes learners from **zero** (never written code) to **Google-level Go engineer**.

The key principle: **syntax enables, engineering differentiates**.

---

## Key Insights Integrated

### 1. Strong Learning Loop

- **Feedback Loop**: Each section has immediate code execution → output → understanding
- **Pressure**: Thinking sections, failure scenarios, time-boxed challenges
- **Validation**: Tests pass/fail, benchmarks measure, code review criteria

### 2. Truly Beginner-Friendly

- **Soft Introduction Pattern**: Use loops/functions as tools before teaching them deeply
- **Mental Model Building**: Start with "what is memory?" before teaching pointers
- **Real-World Context**: Connect to API payloads, DB models, not just syntax

### 3. Engineering Depth Layers

- **Layer 1**: Syntax - What is this? (Beginner)
- Layer 2: Usage - How do I use it? (Intermediate)
- Layer 3: Trade-offs - When to use what? (Advanced)
- Layer 4: Production - What breaks at scale? (Expert)

### 4. Production Notes

Every lesson includes "⚠️ In Production" section covering:

- What can go wrong
- Common bugs and fixes
- What to monitor

### 5. Thinking Sections

Every major concept includes "🤔 Thinking Questions":

- Design thinking
- Constraint reasoning
- Failure engineering
- Debugging

### 6. Failure-Based Learning

Each section has "Failure Scenarios":

- Bug diagnosis exercises
- Performance regression fixes
- Failure injection scenarios
- Scale testing challenges

### 7. Error Handling Framework

Three-tier error system:

- **UserError**: Input validation → return to caller
- **SystemError**: Infrastructure → wrap + propagate
- **FatalError**: Bugs → log + exit

---

## Progress Progression Model

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  PHASE 1: SYNTAX FOUNDATION (0% → 30%)                                      │
│  "I can write Go code"                                                       │
│  ─────────────────────────────────────────────                               │
│  70% Syntax + 30% Basic Why                                                  │
│  Goal: Learn to write working Go code                                        │
│  Pattern: Soft introduction (use before teach deeply)                       │
└─────────────────────────────────────────────────────────────────────────────┘
                                    ↓
┌─────────────────────────────────────────────────────────────────────────────┐
│  PHASE 2: ENGINEERING FOUNDATION (30% → 60%)                               │
│  "I write code that works reliably"                                         │
│  ─────────────────────────────────────────────                               │
│  50% Syntax + 50% Engineering Why                                           │
│  Goal: Write code with proper error handling, validation, structure        │
│  Pattern: Error framework, validation, thinking questions                   │
└─────────────────────────────────────────────────────────────────────────────┘
                                    ↓
┌─────────────────────────────────────────────────────────────────────────────┐
│  PHASE 3: PRODUCTION ENGINEERING (60% → 85%)                               │
│  "I write code that doesn't break at scale"                                 │
│  ─────────────────────────────────────────────                               │
│  30% Syntax + 70% Engineering Why                                           │
│  Goal: Write concurrent, testable, performant code                          │
│  Pattern: Failure injection, production notes, benchmarks                  │
└─────────────────────────────────────────────────────────────────────────────┘
                                    ↓
┌─────────────────────────────────────────────────────────────────────────────┐
│  PHASE 4: SYSTEM ENGINEERING (85% → 100%)                                   │
│  "I think like a senior engineer"                                           │
│  ─────────────────────────────────────────────                               │
│  10% Syntax + 90% Engineering Why                                           │
│  Goal: Build production systems with debugging, resilience, architecture    │
│  Pattern: Flagship project, production failure scenarios                    │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Phase 1: Syntax Foundation (0% → 30%)

**Target**: Complete beginners  
**Focus**: Learn Go syntax, write working programs  
**Engineering Context**: 30% (Production notes, basic why)
**Key Pattern**: Soft introduction - use tools before teaching them

### Section 0: How Computers Work (NEW)

- What is a program?
- How does code become execution?
- Memory basics (what is a variable?)
- Terminal confidence

### Section 1: Core Foundations

- GT.1 Installation
- GT.2 Hello World
- GT.3 How Go Works
- GT.4 Development Environment
- **Milestone**: First program runs

### Section 2: Language Basics

- LB.1 Variables and Constants
- LB.2 Basic Types
- LB.3 Operators
- LB.4 Type Conversions
- **Milestone**: Type converter with tests
- **Thinking**: Why does type matter?

### Section 3: Control Flow ⭐ (UPDATED)

- CF.1 If Statements
- CF.2 Loops (for) - **Soft intro: use before teach**
- CF.3 Switch Statements
- CF.4 Break/Continue
- CF.5 Loop Patterns
- **Milestone**: FizzBuzz variant
- **Thinking**: When to use each control structure?

---

## Phase 2: Engineering Foundation (30% → 60%)

**Target**: Can write basic Go  
**Focus**: Error handling, validation, structure  
**Engineering Context**: 50%
**Key Patterns**: Error framework, Thinking sections, Production notes

### Section 4: Data Structures ⭐ (UPDATED)

- DS.1 Arrays (explain memory layout)
- DS.2 Slices (explain header, append, reallocation)
- DS.3 Maps (explain internals, hash, collisions)
- DS.4 Pointers (explain memory, escape)
- DS.5 Slice Sharing (explain backing array)
- DS.6 Contact Directory
- **Milestone**: Contact directory with validation
- **Production Notes**: Memory leaks, performance
- **Thinking**: When to use array vs slice vs map?
- **Failure Scenario**: Map nil panic, slice growth

### Section 5: Functions and Errors ⭐ (TRANSFORMED)

- FE.1 Functions as Boundaries
- FE.2 Parameters and Returns
- FE.3 Multiple Return Values (error philosophy)
- FE.4 Errors as Values
- FE.5 Validation (security, DoS)
- FE.6 Orchestration
- FE.7 Order Processing System (Engineering Capstone)
- **Milestone**: Order system with security, overflow, error codes
- **Error Framework**: UserError, SystemError, FatalError
- **Thinking**: Why fail fast? What happens at scale?

### Section 6: Types and Interfaces

- TI.1 Structs
- TI.2 Methods
- TI.3 Value vs Pointer Receivers
- TI.4 Interface Basics
- TI.5 Interface Composition
- TI.6 Error Interface Implementation
- **Milestone**: Custom error types
- **Thinking**: What makes good interface?

### Section 7: Packages and Modules

- PM.1 Package Basics
- PM.2 Import and Export
- PM.3 Module Management
- PM.4 Package Design
- PM.5 Circular Dependencies
- PM.6 Vendor and Proxy
- **Thinking**: Why package boundaries matter?

---

## Phase 3: Production Engineering (60% → 85%)

**Target**: Can write reliable code  
**Focus**: Concurrency, testing, performance  
**Engineering Context**: 70%
**Key Patterns**: Failure injection, performance testing, production hardening

### Section 8: Concurrency Fundamentals ⭐

- GC.1 Goroutine Basics
- GC.2 WaitGroups
- GC.3 Channels
- GC.4 Buffered Channels
- GC.5 Channel Closing
- GC.6 Concurrent Downloader
- GC.7 Select
- GC.8 Race Conditions
- GC.9 Select Deep Dive
- GC.10 Sync Primitives
- **Milestone**: Concurrent downloader with limits
- **Thinking**: What breaks at 1000 goroutines?
- **Failure Scenario**: Goroutine leak, race condition, deadlock
- **Production Notes**: Memory exhaustion, backpressure

### Section 9: Context and Cancellation

- CT.1 Context Background
- CT.2 With Cancel
- CT.3 With Timeout
- CT.4 With Value
- CT.5 Timeout-Aware Client
- **Failure Scenario**: Context cancellation propagation

### Section 10: Advanced Concurrency

- CP.1 Errgroup
- CP.2 Fan-out, Fan-in
- CP.3 Pipeline
- CP.4 Worker Pools
- CP.5 Health Checker
- **Thinking**: Backpressure, circuit breaker

### Section 11: Testing ⭐

- TE.1 Unit Testing
- TE.2 Table-Driven Tests
- TE.3 Coverage
- TE.4 Benchmarking
- TE.5 HTTP Handler Testing
- TE.6 Mocking
- **Thinking**: What makes code testable?

### Section 12: Performance and Profiling

- PR.1 pprof Basics
- PR.2 CPU Profiling
- PR.3 Memory Profiling
- PR.4 Benchmark-Driven Development
- **Failure Scenario**: Performance regression

---

## Phase 4: System Engineering (85% → 100%)

**Target**: Ready for production  
**Focus**: Backend, operations, architecture  
**Engineering Context**: 90%
**Key Pattern**: Flagship project, real-world complexity

### Section 13: HTTP Servers ⭐

- HS.1 net/http Basics
- HS.2 Routing
- HS.3 Middleware
- HS.4 Request Handling
- HS.5 Response Writing
- HS.6 Error Handling
- HS.7 Server Timeouts
- **Production Notes**: Slowloris, connection exhaustion
- **Failure Scenario**: Timeout issues

### Section 14: Database Engineering

- DB.1 Database Basics
- DB.2 Connection Pools
- DB.3 Queries and Rows
- DB.4 Transactions
- DB.5 Prepared Statements
- DB.6 Error Handling
- **Thinking**: What happens when pool exhausted?

### Section 15: Production Engineering

- SL.1 Structured Logging
- SL.2 Log Levels
- SL.3 Request Tracing
- GS.1 Graceful Shutdown
- **Failure Scenario**: Shutdown under load

### Section 16: Configuration

- CFG.1 Environment Variables
- CFG.2 Configuration Files
- CFG.3 Flag Parsing
- CFG.4 12-Factor App

### Section 17: Docker and Deployment

- DOCKER.1 Basics
- DOCKER.2 Multi-stage Builds
- DOCKER.3 Docker Compose
- DEPLOY.1 CI/CD

---

## Phase 5: Flagship Project (100%)

### GoScale SaaS Backend ⭐⭐⭐

A production-grade multi-tenant SaaS backend:

**Modules**:

1. Foundation & Configuration
2. Database & Models
3. Authentication System
4. HTTP API Layer
5. Order Processing (with concurrency)
6. Payment Pipeline (mock)
7. Event System & Worker Pools
8. Caching Layer
9. Observability
10. Graceful Shutdown

**Engineering Challenges**:

- 10,000 requests/second stress test
- Memory leak detection
- Race condition debugging
- Cascading failure handling
- Graceful shutdown under load

**Thinking Questions**:

- Architecture decisions
- Trade-off analysis
- Production scenarios

**Assessment**:

- Code review (25%)
- Failure injection (25%)
- Scale testing (25%)
- Architecture discussion (25%)

---

## Key Templates

### 1. Production Notes Template

Every lesson includes:

- What can cause issues
- Common bugs and fixes
- What to monitor

### 2. Thinking Questions Template

- Design thinking
- Constraint reasoning
- Failure engineering
- Debugging

### 3. Failure Learning Template

- Bug diagnosis exercises
- Performance regression fixes
- Failure injection scenarios

### 4. Error Framework

- UserError (input validation)
- SystemError (infrastructure)
- FatalError (bugs)

---

## Summary Table

| Phase         | Sections | Engineering Context | Key Patterns                        |
| ------------- | -------- | ------------------- | ----------------------------------- |
| 1: Syntax     | 0-3      | 30%                 | Soft introduction                   |
| 2: Reliable   | 4-7      | 50%                 | Error framework, Thinking           |
| 3: Scalable   | 8-12     | 70%                 | Failure injection, Production notes |
| 4: Production | 13-17    | 90%                 | Flagship project                    |
| 5: Flagship   | GoScale  | 100%                | Full integration                    |

---

## The Promise

A learner who completes this curriculum can:

- ✅ Write Go code from scratch
- ✅ Structure code for maintainability
- ✅ Handle errors with proper framework
- ✅ Write concurrent code safely
- ✅ Test code comprehensively
- ✅ Profile and optimize
- ✅ Build production HTTP services
- ✅ Work with databases reliably
- ✅ Deploy and operate systems
- ✅ Debug production issues
- ✅ Make architectural trade-offs
- ✅ Think like a senior engineer

---

_This blueprint integrates all insights: learning loops, beginner-friendly soft-introduction, engineering depth layers, production notes, thinking sections, failure-based learning, and the flagship project._
