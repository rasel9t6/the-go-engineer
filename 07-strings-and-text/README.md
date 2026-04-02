# Section 7: Strings & Text Processing

## Beginner → Expert Mapping

| Topic | Level | Importance | Engineering Concept |
|-------|-------|------------|---------------------|
| String Internals | Beginner | High | Immutable byte slices, UTF-8 strings |
| Regexp | Intermediate | Medium | Regular expressions engine |
| Text Templates | Intermediate | High | Dynamic layout generation |

## Engineering Depth
A `string` in Go is a read-only slice of bytes, typically representing UTF-8 encoded text. Since strings are immutable, concatenation using `+=` behaves exceptionally poorly inside loops, operating at $O(N^2)$ due to recursive reallocation. For dynamic text construction, **always use `strings.Builder`** which allocates $O(N)$ memory dynamically just like an integer slice. 

## References
1. **[Go Blog]** [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)

---

## 🏗 Exercise: Log File Parsing System (`6-log-parser`)

### Step-by-Step Instructions & Hints
1. **Load data:** Read lines from a simulated file or multi-line string.
   - *Hint:* Use `strings.Split()` to tokenize the file.
2. **Apply regex:** Compile a `regexp.MustCompile()` pattern out of the loop.
   - *Hint:* Compiling a regex inside a loop is devastating to performance. Compile it at the package level or before the loop!
3. **Build the output:** Extract matched substrings.
   - *Hint:* Use `strings.Builder` to combine the matching items into a single summary output instead of `+=`.
