package utilities

import (
	"os"

	"github.com/mackrorysd/gosix/core"
)

// Rm is a utility for removing directory entries. See [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/rm.html
func Rm(proc core.Proc) int {
	//force := false
	//interactive := false
	recursive := false
	//verbose := false

	paths := []string{}

	for _, arg := range proc.Args {
		if arg[0] == '-' {
			// NOTE: if len(arg) == 1, that has special meaning for other utilities
			// See https://pubs.opengroup.org/onlinepubs/9799919799/basedefs/V1_chap12.html#tag_12_02
			for i := 1; i < len(arg); i++ {
				switch arg[i] {
				case 'd':
					recursive = true
				case 'f':
					//force = true
					proc.Err("-f option is not implemented")
					return core.ExitInvalidArgs
				case 'i':
					//interactive = true
					proc.Err("-i option is not implemented")
					return core.ExitInvalidArgs
				case 'R':
					recursive = true
				case 'r':
					recursive = true
				case 'v':
					//verbose = true
					proc.Err("-v option is not implemented")
					return core.ExitInvalidArgs
				default:
					proc.Err("Unrecognized flag: " + string([]byte{arg[i]}))
					return core.ExitInvalidArgs
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}

	for _, path := range paths {
		var err error
		path = proc.ResolvePath(path)
		if recursive {
			err = os.RemoveAll(path)
		} else {
			var stat os.FileInfo
			stat, err = os.Stat(path)
			if err != nil {
				proc.Err("Error stating path: " + err.Error())
				return core.ExitFileError
			}
			if stat.IsDir() {
				proc.Err("Cannot remove directory: " + path)
				return core.ExitInvalidArgs
			}
			err = os.Remove(path)
		}
		if err != nil {
			proc.Err("Error removing path: " + err.Error())
			return core.ExitFileError
		}
	}
	return core.ExitSuccess
}
