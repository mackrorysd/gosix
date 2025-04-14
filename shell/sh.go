package shell

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/mackrorysd/gosix/core"
)

// Package shell aims to eventually implement a full interactive and batch
// shell. See
// https://pubs.opengroup.org/onlinepubs/9799919799/utilities/sh.html.

// Sh is the entry point for running a shell.
func Sh(proc core.Proc) int {
	if len(proc.Args) > 0 {
		if proc.Args[0] != "-c" {
			proc.Stderr.Write([]byte("Only -c is supported\n"))
			return 1
		}
		if len(proc.Args) < 2 {
			proc.Stderr.Write([]byte("No command provided"))
			return 2
		}
		err := exec.Command(proc.Args[1], proc.Args[2:]...).Run()
		if err != nil {
			return 0
		}
		return 3
	}

	// Interactive REPL
	reader := bufio.NewReader(proc.Stdin)
	for {
		prompt(proc)

		// TODO this needs a parser to handle escape characters, etc.
		command, _ := reader.ReadString('\n')
		// TODO trim excess whitespace internally
		command = strings.Trim(command, "\n\t ")
		tokens := strings.Split(command, " ")

		switch tokens[0] {
		case "exit":
			return 0
		case "echo":
			// TODO: this has flags we need to support
			proc.Out(strings.Join(tokens[1:], " "))
		default:
			cmd := exec.Command(tokens[0], tokens[1:]...)

			cmd.Stdin = proc.Stdin
			cmd.Stdout = proc.Stdout
			cmd.Stderr = proc.Stderr

			err := cmd.Run()
			if err != nil {
				proc.Err("Error running process: " + err.Error())
			}
		}
	}
}

func prompt(proc core.Proc) {
	prompt := proc.Wd + "> "
	// Don't use proc.Out() because we don't want a newline
	proc.Stdout.Write([]byte(prompt))
}
