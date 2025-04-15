/*
Gosix implements many Unix-like utilities.

Each supported utility should have a symlink to the binary, and it will infer
which utility was invoked and execute the relevant functionality.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mackrorysd/gosix/core"
	"github.com/mackrorysd/gosix/shell"
	"github.com/mackrorysd/gosix/utilities"
)

// _main should include all logic that can be abstracted away from making system
// calls or anything else that can't be included in a unit test.
func _main(proc core.Proc) int {
	var f func(core.Proc) int
	switch filepath.Base(proc.Args[0]) {
	case "cat":
		f = utilities.Cat
	case "false":
		f = utilities.False
	case "ln":
		f = utilities.Ln
	case "ls":
		f = utilities.Ls
	case "mkdir":
		f = utilities.Mkdir
	case "rm":
		f = utilities.Rm
	case "sh":
		f = shell.Sh
	case "true":
		f = utilities.True
	default:
		proc.Out("Unrecognized command: " + proc.Args[0])
		return 1
	}
	proc.Args = proc.Args[1:]
	return (f(proc))
}

func main() {
	// TODO setup environment variables

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve working directory: %s\n", err.Error())
	}

	os.Exit(_main(core.Proc{
		Args:   os.Args,
		Wd:     wd,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}))
}
