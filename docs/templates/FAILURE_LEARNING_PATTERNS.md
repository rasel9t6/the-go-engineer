# Failure-Based Learning - Pattern

## Purpose

Learning by breaking is more effective than learning by reading. This pattern creates exercises where learners must diagnose, fix, and prevent failures.

---

## Exercise Types

### Type 1: Bug Diagnosis

```
Given: Broken code that compiles but doesn't work
Task: Run it, observe failure, identify root cause, fix

Example:
- Race condition
- Nil pointer dereference
- Deadlock
- Memory leak
```

### Type 2: Performance Regression

```
Given: Working code with performance issue
Task: Profile, identify hotspot, optimize

Example:
- Unnecessary allocation
- O(n²) algorithm
- Blocking in hot path
```

### Type 3: Failure Injection

```
Given: Working code
Task: Introduce failure, observe cascading effects, fix

Example:
- Database connection failure
- Network timeout
- External service down
```

### Type 4: Scale Testing

```
Given: Code that works at small scale
Task: Test at scale, identify breaking point, fix

Example:
- 10k requests/second
- 1M items in memory
- 100 concurrent users
```

---

## Implementation Template

### File: `failure_scenarios.md`

```markdown
# Failure Scenarios

## Scenario 1: [Name]

### Setup

[Description of the broken code]

### Expected Behavior

[What should happen]

### Actual Behavior

[What actually happens]

### Hint

[Optional hint]

### Investigation Steps

1. [Step 1]
2. [Step 2]
3. [Step 3]

### Solution

[Code fix]

### Why This Happens

[Explanation]

### Prevention

[How to prevent in future]
```

---

## Example: Data Race

````markdown
# Failure Scenario 1: Data Race

### Setup

The following code should increment a counter 1000 times concurrently:

```go
var counter int

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            counter++  // DATA RACE!
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(counter)
}
```
````

### Expected Behavior

Should print 1000 (counter incremented 1000 times)

### Actual Behavior

Prints different number each time (e.g., 987, 1001, 942)

### Investigation Steps

1. Run with `-race` flag: `go run -race main.go`
2. Observe race detector output
3. Identify the unsafe read/write

### Solution

```go
var (
    counter int
    mu      sync.Mutex
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            mu.Lock()
            counter++
            mu.Unlock()
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(counter)
}
```

### Why This Happens

Multiple goroutines reading/writing `counter` simultaneously without synchronization. Operations are not atomic at the machine level.

### Prevention

- Always use `-race` in testing
- Use proper synchronization (mutex, atomic)
- Or use channels for shared state

````

---

## Example: Goroutine Leak

```markdown
# Failure Scenario 2: Goroutine Leak

### Setup
This code starts a worker that processes items from a channel:

```go
func worker(ch chan int) {
    for {
        // Process item
        // No way to stop!
    }
}

func main() {
    ch := make(chan int)
    go worker(ch)

    ch <- 1
    ch <- 2

    // Main exits, worker never stops
}
````

### Expected Behavior

Worker should be stoppable

### Actual Behavior

Worker runs forever, even after main exits (goroutine leak)

### Solution

```go
func worker(ctx context.Context, ch chan int) {
    for {
        select {
        case <-ctx.Done():
            return
        case item := <-ch:
            process(item)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    ch := make(chan int)
    go worker(ctx, ch)

    ch <- 1
    ch <- 2

    cancel()  // Signal worker to stop
    time.Sleep(time.Second)
}
```

### Why This Happens

No way to signal the goroutine to stop. It runs forever.

### Prevention

- Always provide a way to stop goroutines
- Use context cancellation
- Use channels to signal shutdown
- Monitor goroutine count in production

````

---

## Example: Memory Leak

```markdown
# Failure Scenario 3: Slice Memory Leak

### Setup
This code processes items in batches:

```go
func processItems() []Result {
    var results []Result

    for batch := range getBatches() {
        batchResults := processBatch(batch)
        results = append(results, batchResults...)
    }

    return results
}
````

### Expected Behavior

Should process all batches and return results

### Actual Behavior

Memory keeps growing, never released

### Solution

```go
func processItems() []Result {
    results := make([]Result, 0, 1000)  // Pre-allocate

    for batch := range getBatches() {
        batchResults := processBatch(batch)
        results = append(results, batchResults...)
    }

    return results
}

// OR process in chunks without accumulating
func processItems() []Result {
    var allResults []Result

    for batch := range getBatches() {
        batchResults := processBatch(batch)

        // Process immediately instead of accumulating
        for _, r := range batchResults {
            saveResult(r)
        }
    }

    return allResults
}
```

### Why This Happens

Slice keeps growing as batches are added. Even if processing finishes, the underlying array capacity remains high.

### Prevention

- Pre-allocate with reasonable capacity
- Process in chunks, don't accumulate all in memory
- Use streaming patterns
- Monitor memory in production

````

---

## Example: Connection Pool Exhaustion

```markdown
# Failure Scenario 4: DB Connection Pool Exhaustion

### Setup
This code queries the database in a loop:

```go
func getUsers() ([]User, error) {
    users := []User{}

    // BAD: Opens new connection each time!
    for id := range getUserIDs() {
        user, err := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}
````

### Expected Behavior

Should use connection pooling properly

### Actual Behavior

Connection pool exhausted, queries hang

### Solution

```go
func getUsers() ([]User, error) {
    // Use single query with IN clause
    ids := getUserIDs()
    if len(ids) == 0 {
        return nil, nil
    }

    placeholders := strings.Repeat("?,", len(ids)-1) + "?"
    query := fmt.Sprintf("SELECT * FROM users WHERE id IN (%s)", placeholders)

    rows, err := db.Query(query, ids)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        rows.Scan(&user.ID, &user.Name, &user.Email)
        users = append(users, user)
    }

    return users, rows.Err()
}
```

### Why This Happens

Each query opens new connection from pool. Pool gets exhausted.

### Prevention

- Batch queries
- Use connection pool properly
- Set appropriate pool size
- Monitor connection usage

```

---

## Production Failure Testing

In the flagship project, include mandatory failure scenarios:

1. **Traffic spike test** - 10x normal load
2. **Circuit breaker test** - External service down
3. **Memory stress test** - Run for extended period
4. **Concurrency test** - Race conditions
5. **Shutdown test** - Graceful shutdown under load

---

*Failure-based learning creates engineers who can debug, not just code.*
```
