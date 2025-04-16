package shell

import (
	"bufio"
	"os/exec"
	filepath "path/filepath"
	"strings"

	"github.com/mackrorysd/gosix/core"
	"github.com/mackrorysd/gosix/term"
)

// Package shell aims to eventually implement a full interactive and batch
// shell. See [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/sh.html.

// Sh is the entry point for running a shell.
func Sh(proc core.Proc) int {
	if len(proc.Args) > 0 {
		if proc.Args[0] != "-c" {
			proc.Err("Only -c is supported")
			return 1
		}
		if len(proc.Args) < 2 {
			proc.Err("No command provided")
			return 2
		}
		y, _ := run(&proc, strings.Join(proc.Args[1:], " "))
		return y
	}

	// Interactive REPL
	var y int
	reader := bufio.NewReader(proc.Stdin)
	for exit := false; exit == false; {
		prompt(proc)

		// TODO this needs a parser to handle escape characters, etc.
		command, _ := reader.ReadString('\n')
		y, exit = run(&proc, command)
	}
	return y
}

func run(proc *core.Proc, command string) (int, bool) {
	// TODO trim excess whitespace internally
	command = strings.Trim(command, "\n\t ")
	tokens := strings.Split(command, " ")

	switch tokens[0] {
	case "cd":
		// TODO: this has other flags
		// See https://pubs.opengroup.org/onlinepubs/9799919799/utilities/cd.html
		if len(tokens) > 2 {
			proc.Err("cd only takes one parameter; ignoring all by the first")
		}
		// TODO: ensure it exists
		proc.Wd = proc.ResolvePath(tokens[1])
		return 0, false
	case "exit":
		return 0, true
	case "echo":
		// See https://pubs.opengroup.org/onlinepubs/9799919799/utilities/echo.html
		// The Shell Command Language spec has requirements around built-ins
		if len(tokens) == 1 {
			proc.Out("")
		} else {
			proc.Out(strings.Join(tokens[1:], " "))
		}
		return 0, false
	default:
		// TODO do variable substitution on tokens, or at parsing time

		exe, err := exec.LookPath(tokens[0])
		if err != nil {
			proc.Err("Error looking in path: " + err.Error())
			return 1, false
		}
		exe, err = filepath.EvalSymlinks(exe)
		if err != nil {
			proc.Err("Error resolving executable: " + err.Error())
			return 1, false
		}
		cmd := exec.Command(exe, tokens[1:]...)

		// This is probably unsafe, but this allows the symlink names to be
		// passed in instead of the resolved, absolute path.
		cmd.Args[0] = tokens[0]

		cmd.Dir = proc.Wd
		cmd.Stdin = proc.Stdin
		cmd.Stdout = proc.Stdout
		cmd.Stderr = proc.Stderr

		err = cmd.Run()
		if err != nil {
			proc.Err("Error running process: " + err.Error())
			return 1, false
		}
		return cmd.ProcessState.ExitCode(), false
	}
}

func prompt(proc core.Proc) {
	prompt := term.SGR(term.CyanForeground) +
		proc.Wd + "> " +
		term.SGR(term.Reset)

	// Don't use proc.Out() because we don't want a newline
	proc.Stdout.Write([]byte(prompt))
}
