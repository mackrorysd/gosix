package filesystem

import (
	"fmt"
	"os"

	"github.com/mackrorysd/gosix/core"
)

// Package filesystem implements commands related to managing files, etc.

// Ln is a command function for creating hard links or symlinks. See
// https://pubs.opengroup.org/onlinepubs/9799919799/utilities/ln.html.
func Ln(proc core.Proc) int {
	// For now, only supporting "single file" invocation
	// and the -f flag
	if len(proc.Args) == 0 || proc.Args[0] != "-f" {
		proc.Stderr.Write([]byte("Only forced linking is supported\n"))
		return 1
	}

	oldname := proc.ResolvePath(proc.Args[1])
	newname := proc.ResolvePath(proc.Args[2])

	err := os.Remove(newname)
	if err != nil && !os.IsNotExist(err) {
		proc.Stdout.Write([]byte(err.Error()))
		return 2
	}

	err = os.Link(oldname, newname)
	if err != nil {
		proc.Stderr.Write([]byte(fmt.Sprintf("Error creating link: %s", err.Error())))
		return 3
	}
	return 0
}
