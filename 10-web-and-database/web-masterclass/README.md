# Section 13: Web Development Masterclass

## Learning Objectives

This is the **capstone section** â€” it ties together every concept from earlier sections into a real web application.

By the end, you will understand:
- HTTP routing with `net/http` and Go 1.22+ method-based patterns
- Dependency injection via the `application` struct
- HTML template rendering with caching and layouts
- Middleware pattern (`func(http.Handler) http.Handler`)
- Session management and flash messages
- Authentication with bcrypt password hashing
- Form processing and server-side validation
- Repository pattern for database abstraction
- Pagination computation and API response structure
- Nested comment systems

## Contents

| Directory | Topic | Lectures | Level |
|-----------|-------|----------|-------|
| `1-routing/` | HTTP routing, handlers, path params | 1-2 | Intermediate |
| `2-dependency-injection/` | App struct, slog, handler methods | 3-4 | Intermediate |
| `3-templates/` | html/template, caching, layouts, CSS | 5-7 | Intermediate |
| `4-middleware/` | Security headers, logging, panic recovery | 9 | Advanced |
| `5-sessions/` | Cookie sessions, flash messages | 10 | Advanced |
| `6-auth/` | bcrypt, login, registration, auth middleware | 11, 17-19 | Advanced |
| `7-forms/` | Form validation library, error handling | 12-16 | Advanced |
| `8-posts-crud/` | Repository pattern, SQLite, CRUD operations | 23-30 | Advanced |
| `9-pagination/` | Metadata computation, dynamic links | 25, 34, 38 | Advanced |
| `10-comments/` | Nested comments, thread-safe store | 28, 36-37 | Advanced |
| `11-websockets/` | **Capstone** â€” WebSockets and live updates | All | Advanced |

## How to Run

Each sub-directory is a standalone, runnable example:

```bash
# Run individual topics
go run ./10-web-and-database/web-masterclass/1-routing
go run ./10-web-and-database/web-masterclass/4-middleware
go run ./10-web-and-database/web-masterclass/8-posts-crud

# Run the full capstone app
go run ./10-web-and-database/web-masterclass/11-websockets

# Test endpoints with curl
curl http://localhost:8080/health
curl -X POST -d '{"email":"test@go.dev","password":"password123"}' http://localhost:8080/register
curl -X POST -d '{"email":"test@go.dev","password":"password123"}' http://localhost:8080/login
```

## Architecture

```
Request â†’ secureHeaders â†’ logRequest â†’ recoverPanic â†’ Router â†’ Handler
                                                          â†“
                                                   Repository Interface
                                                          â†“
                                                    SQLite Database
```

## Recommended Study Order

1. Start with `1-routing/` to understand HTTP basics
2. Move to `2-dependency-injection/` to learn the app struct pattern
3. Study `4-middleware/` â€” the most important pattern for production apps
4. Explore `6-auth/` for security fundamentals
5. Finish with `11-websockets/` to see everything wired together

## References

- [Let's Go by Alex Edwards](https://lets-go.alexedwards.net/)
- [Go 1.22 HTTP Routing Enhancements](https://go.dev/blog/routing-enhancements)
- [Writing Web Applications (Official)](https://go.dev/doc/articles/wiki/)


## Learning Path

| ID | Lesson | Concept | Requires |
| --- | --- | --- | --- |
| WM.1 | [routing](./1-routing) | http.NewServeMux Â· method patterns Â· {param} Â· HandleFunc | ðŸŸ¢ entry |
| WM.2 | [dependency injection](./2-dependency-injection) | application struct Â· constructor Â· handler methods on struct | WM.1 |
| WM.3 | [templates](./3-templates) | html/template Â· cache once Â· layout + partials + page | WM.1, WM.2 |
| WM.4 | [middleware](./4-middleware) | func(http.Handler) http.Handler Â· chain Â· panic recovery | WM.1, WM.2, WM.3 |
| WM.5 | [sessions](./5-sessions) | Cookie-based sessions Â· flash messages Â· session store | WM.1, WM.2, WM.3, WM.4 |
| WM.6 | [authentication](./6-auth) | bcrypt Â· requireAuth middleware Â· context.WithValue user ID | WM.1, WM.2, WM.4, WM.5 |
| WM.7 | [forms](./7-forms) | Form struct Â· Required Â· MinLength Â· MatchesField Â· Valid() | WM.1, WM.2, WM.4 |
| WM.8 | [posts CRUD](./8-posts-crud) | PostRepository interface Â· ExecContext Â· pagination offset | WM.1, WM.2, WM.4, WM.6, WM.7 |
| WM.9 | [pagination](./9-pagination) | computeMetadata Â· LIMIT/OFFSET Â· HasNext/HasPrev Â· links | WM.8 |
| WM.10 | [comments](./10-comments) | Adjacency list Â· 2-pass tree build Â· sync.RWMutex store | WM.8, WM.9 |
