package shell

import (
	"bufio"
	"os/exec"

	"github.com/mackrorysd/gosix/core"
)

func Sh(proc core.Proc) int {
	// Minimal functionality for Dockerfile
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

	// If there is input, just echo it for now
	reader := bufio.NewReader(proc.Stdin)
	proc.Stdout.Write([]byte("$ "))
	text, _ := reader.ReadString('\n')
	proc.Stdout.Write([]byte(text + "\n"))
	return 0
}
