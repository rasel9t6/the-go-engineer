# Linux processes

## Mission

Understand and apply Linux processes in the context of professional Go software engineering.

## Prerequisites

- core-14-17

## Mental Model

A process is an instance of a running program with its own address space, file descriptors, and process ID. In Go, exec.Command creates a child process that is isolated from the parent — changes to the child do not affect the parent. The parent-child relationship is asymmetric: the parent can wait for the child (Wait), send it signals (Signal), or kill it (Kill), but the child cannot affect the parent. Go's goroutines are lightweight threads within the same process, while processes are OS-level heavyweight isolation units.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

exec.Command creates the Cmd struct but does not fork. Calling Start() invokes os.StartProcess, which calls syscall.ForkExec on Unix. ForkExec clones the current process (fork), replaces the child's memory with the target binary (execve), and returns the child's PID. The child inherits the parent's file descriptors unless cmd.ExtraFiles or cmd.Stdout/Stderr/Stdin are set — in which case Go creates pipes (socketpair on Linux) between parent and child. cmd.Wait() calls syscall.Wait4(pid, &wstatus, 0, &rusage) which blocks until the child changes state (exits, is signaled, or is stopped). The wstatus encodes the exit code and signal information. Go's ProcessState parses wstatus to provide ExitCode(), Success(), and String().

## Run Instructions

```bash
go run ./curriculum/modules/15-docker-cicd-deployment/lessons/01-linux-processes
go test ./curriculum/modules/15-docker-cicd-deployment/lessons/01-linux-processes
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using exec.Command without checking the process exit code — a command that fails (exit code 1) still returns no error from Run/Wait if the process ran. The error is only returned when the process exits non-zero with an ExitError that contains stderr. Always check the error from cmd.Wait() and the exit code.
- Zombie processes from not calling cmd.Wait() — creating a process with cmd.Start() and forgetting to call cmd.Wait() leaves the child process as a zombie (defunct) until the parent exits. The OS keeps the process table entry until Wait is called to collect the exit status.
- Not handling SIGCHLD in long-running Go services — when a child process exits, the kernel sends SIGCHLD to the parent. If the parent is a Go program that does not handle SIGCHLD, the child becomes a zombie. Go's os/exec handles this internally for managed processes, but processes created via syscall.ForkExec directly can leak zombies.
- Assuming os/exec.FindProcess finds a running process — FindProcess on Unix returns a Process struct even if the PID does not exist. It only checks that the PID is valid syntactically (PID > 0). The process may have already exited. Always use signal or /proc to verify the process is alive.

## In Production

Go's build system uses process spawning extensively: go build invokes the Go compiler as a subprocess, go test compiles and runs test binaries as subprocesses, and go generate runs user-specified commands as subprocesses. Production Go services use exec.Command for: running backup scripts, invoking git for CI/CD workflows, calling ffmpeg for video processing pipelines, running database migrations as subprocesses, and executing shell commands in admin APIs.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-15-02`.
