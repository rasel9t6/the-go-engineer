# Section 16: HTTP Clients & Mocking

## Beginner → Expert Mapping

| Topic | Level | Importance | Engineering Concept |
|-------|-------|------------|---------------------|
| REST API Fetch | Beginner | High | `net/http` client execution |
| Dependency Injection| Advanced | **Critical** | Interfaces as seams |
| Testify Mock | Advanced | High | Auto-generation mock contracts |

## Engineering Depth
You should never use `http.Get` in production, as it does not enforce a timeout, leaving your goroutine vulnerable to hanging indefinitely on a bad network request. Always initialize a custom `&http.Client{Timeout: time.Second * 10}`. To test APIs without external internet dependency, abstract the client behind an interface, achieving perfect Dependency Inversion.

## References
1. **[Go Blog]** [The empty interface (advanced routing/testing)](https://go.dev/doc/effective_go#interfaces)

---

## 🏗 Exercise: API Data Fetcher (`6-testify-mock`)

*Complete the mock-based test suite in the `6-testify-mock` sub-folder.*
