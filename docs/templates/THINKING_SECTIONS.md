# Thinking Sections - Engineering Question Pattern

## Purpose

Every key concept should have a "Thinking Section" that forces learners to reason about design decisions, trade-offs, and implications. This is what creates senior engineers.

---

## Question Categories

### 1. Design Thinking Questions

```
❓ Why does Go prefer composition over inheritance?

❓ When should I use a struct vs a map vs a slice?

❓ What is the cost of abstraction?

❓ Where should boundaries exist in this system?

❓ Why is this design better than alternatives?
```

### 2. Constraint Reasoning Questions

```
❓ What happens when memory is limited?

❓ What happens when this runs 1000 times per second?

❓ What happens when network latency is high?

❓ What happens when database is slow?

❓ What is the time complexity of this operation?
```

### 3. Failure Engineering Questions

```
❓ What happens when this component fails?

❓ How does failure propagate?

❓ What is the blast radius?

❓ How do we recover?

❓ What are the retry semantics?
```

### 4. Debugging Questions

```
❓ How do you debug when this returns an error?

❓ What would you look for in logs?

❓ What metrics would you add?

❓ How do you reproduce this issue in test?
```

---

## Example: Functions (FE.1)

```markdown
## 🤔 Thinking Questions

### Design Thinking

1. **Why decompose?**
   - When should you break code into functions?
   - What makes a "good" function boundary?
   - What is single responsibility principle in practice?

2. **Function size trade-off**
   - Too small: excessive indirection
   - Too large: hard to test, understand
   - What's the right size?

### Constraint Reasoning

1. **Performance cost**
   - What's the cost of a function call?
   - When should you inline vs extract?
   - Does optimization matter here?

2. **Memory impact**
   - What gets copied when passing parameters?
   - When does this matter?

### Failure Engineering

1. **What happens on error?**
   - How should errors propagate?
   - Who handles them?

2. **What happens on panic?**
   - When is panic appropriate?
   - How does it affect the caller?

### Debugging

1. **How do you debug?**
   - What's the call stack?
   - How do you trace execution?
```

---

## Example: Concurrency

```markdown
## 🤔 Thinking Questions

### Design Thinking

1. **Concurrency vs Parallelism**
   - What's the difference?
   - When is each appropriate?

2. **When to use concurrency?**
   - What problems is concurrency good for?
   - What problems does concurrency NOT solve?

3. **Channel vs Mutex**
   - When to use channels?
   - When to use mutexes?
   - What are the trade-offs?

### Constraint Reasoning

1. **What happens with 1000 goroutines?**
   - How does Go scheduler handle this?
   - What's the memory cost?

2. **What happens at CPU limits?**
   - How does GOMAXPROCS affect this?
   - When should you increase it?

### Failure Engineering

1. **What happens on goroutine leak?**
   - How do they accumulate?
   - How do you detect them?

2. **What happens on race condition?**
   - How do races manifest?
   - How do you debug them?

3. **What happens on deadlock?**
   - How do deadlocks occur?
   - How do you debug?

### Failure Scenarios

1. **Scenario: 10k requests/second**
   - What breaks first?
   - How do you handle backpressure?

2. **Scenario: Worker pool exhaustion**
   - What happens when all workers busy?
   - How do you shed load?
```

---

## Example: HTTP Servers

```markdown
## 🤔 Thinking Questions

### Design Thinking

1. **RESTful Design**
   - What makes an API RESTful?
   - When to use REST vs RPC?

2. **Error Response Design**
   - What should error responses contain?
   - How to handle error codes consistently?

### Constraint Reasoning

1. **Timeouts**
   - What happens without ReadTimeout?
   - What happens without WriteTimeout?
   - Why is IdleTimeout important?

2. **Connection Management**
   - What happens with unlimited connections?
   - What's the cost per connection?

### Failure Engineering

1. **Slowloris Attack**
   - How does this work?
   - How do you prevent it?

2. **Connection Pool Exhaustion**
   - What happens when DB pool is full?
   - How do you handle gracefully?

### Production Questions

1. **What metrics matter?**
   - Request latency p50, p95, p99
   - Error rate
   - Active connections
   - Goroutine count

2. **What do you log?**
   - Request ID for tracing
   - Correlation IDs
   - Timing information
```

---

## Implementation Pattern

For each major section, add a file or section:

```
Thinking-Questions.md
```

Or add directly to README:

```markdown
## 🤔 Thinking Questions

### Design Thinking

[Questions]

### Constraint Reasoning

[Questions]

### Failure Engineering

[Questions]

### Debugging

[Questions]
```

---

## Grading Thinking Answers

These questions should be **discussed, not just answered**.

Expected answer quality:

- **Beginner**: Can identify the concept
- **Intermediate**: Can explain trade-offs
- **Senior**: Can defend design decisions with evidence
- **Expert**: Can propose alternatives and explain why

---

_Every major section should have Thinking Questions to develop engineering judgment._
