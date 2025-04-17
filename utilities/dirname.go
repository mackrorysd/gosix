package utilities

import (
	"strings"

	"github.com/mackrorysd/gosix/core"
)

// Dirname is a utility for getting the directory a path is in. See
// [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/dirname.html
func Dirname(proc core.Proc) int {
	var path string

	if len(proc.Args) != 1 {
		proc.Err("Must provide 1 argument")
		return core.ExitInvalidArgs
	}
	path = proc.Args[0]

	if path == "" {
		proc.Out(".")
		return core.ExitSuccess
	}

	// Remove trailing slashes that aren't also leading slashes
	path = string(path[0]) + strings.TrimRight(path[1:], "/")

	if !strings.Contains(path, "/") {
		proc.Out(".")
		return core.ExitSuccess
	}

	if path == "/" {
		proc.Out(path)
		return core.ExitSuccess
	}

	// This does not take escaped slashes into account
	tokens := strings.Split(path, "/")
	path = strings.Join(tokens[0:len(tokens)-1], "/")

	proc.Out(path)
	return core.ExitSuccess
}
