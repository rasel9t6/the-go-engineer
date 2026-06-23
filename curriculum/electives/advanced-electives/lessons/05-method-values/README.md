# Method values

## Mission

Understand and apply Method values in the context of professional Go software engineering.

## Prerequisites

- elective-04

## Mental Model

A method value is a function closure that has a receiver bound to it. When you write f := obj.Method, f is a function value that, when called, calls Method on obj. The receiver is captured at the point of assignment — if you change obj later, f still calls Method on the original obj. A method expression is the unbound function: T.Method returns a function whose first parameter is the receiver.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A method value m := obj.Method creates a closure that holds a pointer to obj. When called, it invokes Method on the captured pointer. This is equivalent to func(args...) { return obj.Method(args...) }. A method expression T.Method creates a function where the receiver is the first explicit parameter: func(receiver T, args...) returns. This allows passing methods as callbacks where the caller supplies the receiver, useful for sort.Interface and http.Handler patterns.

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

- Confusing method values (bound) with method expressions (unbound) — a method value captures the receiver, while a method expression requires the receiver as the first argument.
- Capturing a method value inside a loop — the method captures the loop variable by reference, and all callbacks operate on the last iteration's value.
- Using a method value when a method expression is needed — passing obj.Method to sort.Sort does not work because sort.Interface expects Len, Less, Swap as separate methods, not a single function.
- Assuming method values are free — each method value allocation creates a closure, adding GC pressure in hot paths.
- Using method values with value receivers on large structs — the entire struct is copied into the closure.

## In Production

Method values are used in every Go HTTP service (http.HandlerFunc), every sort implementation (sort.Sort), every event emitter (registering callbacks), and every test suite (registering test helpers). Production Go code uses method values for dependency injection (passing service methods to HTTP handlers) and middleware chaining.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-06`.
