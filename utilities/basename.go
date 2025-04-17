package utilities

import (
	"strings"

	"github.com/mackrorysd/gosix/core"
)

// Basename is a utility for stripping parent directories and suffixes from paths.
// See [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/basename.html
func Basename(proc core.Proc) int {
	var path string
	var suffix string

	switch len(proc.Args) {
	case 1:
		path = proc.Args[0]
	case 2:
		path = proc.Args[0]
		suffix = proc.Args[1]
	default:
		proc.Err("Invalid number of arguments")
		return core.ExitInvalidArgs
	}

	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}
	path = strings.TrimRight(path, "/")
	if path == "" {
		// Any string that is all slashes or empty should finish here
		proc.Out("")
		return core.ExitSuccess
	}

	if strings.Contains(path, "/") {
		tokens := strings.Split(path, "/")
		path = tokens[len(tokens)-1]
	}

	if suffix != "" && suffix != path {
		path = strings.TrimSuffix(path, suffix)
	}
	proc.Out(path)
	return core.ExitSuccess
}
