# HC.4 Terminal Confidence

## Mission

Build confidence with the terminal as the text-based environment that launches programs, passes arguments, and handles output streams.

## Prerequisites

- `HC.3` memory basics

## Mental Model

The terminal is not "where scary commands live." It is a professional interface for the operating system.

Think of it as a **Conversation with the OS**. You send a request (a command), and the OS provides a response (output).

## Visual Model

```mermaid
graph LR
    User["You type a command"] --> Shell["Shell parses tokens"]
    Shell --> PATH["Finds binary in PATH"]
    PATH --> Process["OS launches Process"]
    Process --> stdout["Standard Output (Default)"]
    Process --> stderr["Standard Error (Issues)"]
    stdout --> Screen["Terminal Screen"]
    stderr --> Screen
```

## Machine View

When you press Enter in the shell:
1. The shell parses the command into a program name and arguments.
2. It looks up the program's location using the `PATH` environment variable.
3. The OS creates a new process for that program.
4. The process is connected to three default **File Descriptors** (indices in the process's open file table):
   - **0 (stdin)**: Input from the keyboard.
   - **1 (stdout)**: Normal output.
   - **2 (stderr)**: Error messages.

> [!NOTE]
> The "Everything is a file" philosophy in Unix-like systems makes these descriptors incredibly powerful for I/O, as you will see in [FS.1 File Basics](../../05-packages-io/02-io-and-cli/filesystem/1-files/README.md).

## Run Instructions

```bash
go run ./00-how-computers-work/4-terminal-confidence
```

## Code Walkthrough

- **stdout**: Using `fmt.Println` writes to the standard output by default.
- **stderr**: Using `fmt.Fprintln(os.Stderr, ...)` writes to the standard error stream. This separation allows you to capture logs in a file while still seeing errors on the screen.

## Try It

1. Run the lesson normally and read both lines.
2. Run `go run ./00-how-computers-work/4-terminal-confidence > output.txt`. Only the error message stays on your screen.
3. Run `go run ./00-how-computers-work/4-terminal-confidence 2> errors.txt`. Now the error is trapped in the file.

## In Production

Every modern backend is managed through the terminal or terminal-like APIs. Knowing how to filter logs (`grep`), find processes (`ps`), and check exit codes (`$?`) is essential for any backend engineer.

## Thinking Questions

1. Why does redirecting stdout not automatically redirect stderr?
2. Why do shells care about numeric exit codes instead of reading the words in the program output?
3. What would break if the shell could not find programs through the `PATH`?

## Next Step

Next: `HC.5` -> [`00-how-computers-work/5-os-processes`](../5-os-processes/README.md)
