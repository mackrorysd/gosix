package filesystem

import (
	"fmt"
	"os"

	"github.com/mackrorysd/gosix/core"
)

func Ln(proc core.Proc) int {
	// For now, only supporting "single file" invocation
	// and the -f flag
	if len(proc.Args) == 0 || proc.Args[0] != "-f" {
		proc.Stderr.Write([]byte("Only forced linking is supported\n"))
		return 1
	}
	oldname := proc.Args[1]
	newname := proc.Args[2]
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
