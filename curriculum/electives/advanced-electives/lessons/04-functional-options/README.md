# Functional options

## Mission

Understand and apply Functional options in the context of professional Go software engineering.

## Prerequisites

- elective-03

## Mental Model

A functional option is a function that modifies a config struct. The constructor accepts an empty config struct, applies each option function to it, and uses the resulting config. Each option is a standalone function that sets exactly one field. Callers specify only the options they care about — unset fields use their zero values or defaults. This is equivalent to the builder pattern but uses functions instead of a builder struct.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The functional options pattern uses a type Option func(*Config). The constructor NewServer(opts ...Option) creates a default Config, then calls each option in sequence: for _, opt := range opts { opt(&cfg) }. Each option is a closure: func WithTimeout(d time.Duration) Option { return func(c *Config) { c.Timeout = d } }. The closure captures the parameter (d) and returns a function that applies it. Options are applied in order, so later options override earlier ones for the same field.

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

- Using positional parameters for optional config — adding a new parameter breaks all callers and makes the function signature unreadable.
- Mutating the option closure's captured variable after creating the option — the option captures a reference, and changes to the variable affect the option's behavior.
- Not providing defaults for unset options — callers must specify every option or accept zero values that may not be valid.
- Applying options that conflict — two options set the same field to different values, and the last one wins (silently), which may surprise callers who expect an error.
- Using functional options for required parameters — required parameters should be explicit constructor arguments, not optional options.

## In Production

Functional options are used in every Go library with configurable constructors: gRPC dial options, HTTP client options, database connection options, and logger configuration. Google Cloud Go client libraries use functional options for all service clients. Production Go services use functional options for server configuration, middleware setup, and client initialization.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-05`.
