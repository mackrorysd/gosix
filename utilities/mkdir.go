package utilities

import (
	fs "io/fs"
	"os"

	"github.com/mackrorysd/gosix/core"
)

// Mkdir is a utility for making directories. See [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/mkdir.html
func Mkdir(proc core.Proc) int {
	dirs := []string{}
	intermediate := false

	for i := 0; i < len(proc.Args); i++ {
		arg := proc.Args[i]
		if arg == "-m" {
			proc.Err("-m option is not implemented")
			return core.ExitInvalidArgs
		}
		if arg == "-p" {
			intermediate = true
			continue
		}
		// Technically once we encounter a dir this could stop parsing flags
		dirs = append(dirs, arg)
	}

	for _, dir := range dirs {
		path := proc.ResolvePath(dir)
		var err error
		if intermediate {
			// There may be nuances to how the mode is handled in this scenario
			err = os.MkdirAll(proc.ResolvePath(path), fs.ModeDir|0700)
		} else {
			err = os.Mkdir(proc.ResolvePath(path), fs.ModeDir|0700)
		}
		if err != nil {
			proc.Err("Failed to create directory: " + err.Error())
			return core.ExitFileError
		}
	}

	return 0
}
