# Section 15: Time & Scheduling

## Beginner â†’ Expert Mapping

| Topic | Level | Importance | Engineering Concept |
|-------|-------|------------|---------------------|
| Timers | Beginner | High | Delayed execution |
| Tickers | Intermediate| Medium | Scheduled recurrence |
| Context | Advanced | **Critical** | Request-scoped deadlines |

## Engineering Depth
A common memory leak in Go is abandoning `time.Ticker` instances. The ticker uses a background goroutine to push to its channel. If you do not call `ticker.Stop()`, that goroutine lives forever. Contexts (`context.WithTimeout`) are the Google standard mechanism for passing scoped deadlines down massive call stacks without relying on global tickers.

## References
1. **[Go Docs]** [Package time](https://pkg.go.dev/time)

---

## ðŸ— Exercise: Console Reminder (`7-reminder`)

Build a countdown reminder that uses `time.NewTicker` and `time.AfterFunc`. Try it yourself first!

```bash
go run ./11-concurrency/time-and-scheduling/7-reminder/_starter 5 "Break time!"  # Try the exercise
go run ./11-concurrency/time-and-scheduling/7-reminder 5 "Break time!"           # See the solution
```


## Learning Path

| ID | Lesson | Concept | Requires |
| --- | --- | --- | --- |
| TM.1 | [time basics](./1-time) | time.Time Â· Duration Â· Add Â· Sub Â· wall vs monotonic clock | ðŸŸ¢ entry |
| TM.2 | [formatting](./2-formatting) | Reference Time "2006-01-02 15:04:05" Â· Parse Â· RFC3339 | TM.1 |
| TM.3 | [timers &amp; tickers](./3-timer-and-ticker) | time.NewTimer Â· time.NewTicker Â· &lt;-C Â· ticker.Stop() leak | TM.1, TM.2 |
| TM.4 | [random numbers](./4-random) | rand/v2 Â· IntN Â· Shuffle Â· Perm Â· seeded PCG | TM.1 |
| TM.5 | [scheduler](./5-schedule) | Actor model Â· ScheduleOnce Â· ScheduleInterval Â· StopAll drain | TM.3 |
| TM.6 | [timezones](./6-timezone) | time.LoadLocation Â· IANA database Â· In() Â· always store UTC | TM.1, TM.2 |
