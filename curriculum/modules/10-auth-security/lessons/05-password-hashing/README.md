# Password hashing

## Mission

Understand and apply Password hashing in the context of professional Go software engineering.

## Prerequisites

- core-10-04

## Mental Model

Password hashing is a one-way function with a work factor. The hash is deliberately slow: slow for the attacker (brute force), manageable for the server (login). Cost parameter tunes this.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

bcrypt runs the Blowfish key schedule with password and salt for 2^cost iterations. Each iteration depends on the previous one, making it inherently sequential and impossible to parallelize on GPU.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/05-password-hashing
go test ./curriculum/modules/10-auth-security/lessons/05-password-hashing
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using fast hash functions (SHA-256) for passwords — fast hashes can be brute-forced at billions of attempts/second on GPU.
- Implementing custom password hashing logic — use bcrypt, scrypt, or argon2 from the Go stdlib.
- Storing bcrypt hashes in a VARCHAR(60) field without testing with the slowest cost parameter.

## In Production

bcrypt is the industry standard for password hashing in Go. Large-scale Go services (Uber, Docker) use bcrypt for user authentication.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-06`.
