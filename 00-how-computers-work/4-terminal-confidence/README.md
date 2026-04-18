# HC.4 Terminal Confidence

## Mission
Build foundational confidence in navigating and using the terminal, an essential environment for production engineering.

## Prerequisites
- `HC.3` Memory Basics: Stack and Heap

## Mental Model
The terminal is just a text interface to the operating system. Instead of moving a mouse to click a folder, you type `cd`. Instead of double-clicking a program to run it, you type its name.

## Visual Model
```mermaid
graph LR
    A[You Type] --> B[Shell parses command]
    B --> C[Finds Program in $PATH]
    C --> D[Runs Program]
    D --> E[Prints stdout/stderr]
```

## Machine View
The terminal runs a shell (bash, zsh, powershell). The shell executes commands. Every program returns an exit code (0 for success, non-zero for failure). The shell also maintains environment variables (like `$PATH`) that programs can read.

## Run Instructions
```bash
go run ./00-how-computers-work/4-terminal-confidence
```

## Code Walkthrough
In `main.go`, we write to both standard output (`stdout`) and standard error (`stderr`), demonstrating how the terminal handles different streams.

## Try It
1. Run `go run main.go > output.txt` and see what gets printed vs what goes into the file.
2. Check the exit code of a command using `echo $?`.

## ⚠️ In Production
**You will SSH into production servers.** When things go wrong at 2am, there is no GUI. You will have a terminal, log files, and commands. Start using the terminal for everything now.

## 🤔 Thinking Questions
1. You run `go` and get `command not found`. What's the most likely cause? How would you diagnose it?
2. Why do you think exit code 0 means success and non-zero means failure?
3. You have a log file with 2 million lines. How would you find all lines containing "PANIC" without opening the file in an editor?

## Next Step
[HC.5 OS Processes](../5-os-processes)
