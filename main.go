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
	"github.com/mackrorysd/gosix/filesystem"
	"github.com/mackrorysd/gosix/shell"
)

// _main should include all logic that can be abstracted away from making system
// calls or anything else that can't be included in a unit test.
func _main(proc core.Proc) int {
	var f func(core.Proc) int
	switch filepath.Base(proc.Args[0]) {
	case "ln":
		f = filesystem.Ln
	case "sh":
		f = shell.Sh
	default:
		proc.Stdout.Write([]byte(fmt.Sprintf("Unrecognized command: %s", proc.Args[0])))
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
