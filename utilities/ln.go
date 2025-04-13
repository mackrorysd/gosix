package utilities

import (
	"os"

	"github.com/mackrorysd/gosix/core"
)

// Ln is a utility for creating hard links or symlinks. See
// [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/ln.html
func Ln(proc core.Proc) int {
	force := false
	symlink := false
	files := []string{}

	for _, arg := range proc.Args {
		if arg[0] == '-' {
			// NOTE: if len(arg) == 1, that has special meaning for other utilities
			// See https://pubs.opengroup.org/onlinepubs/9799919799/basedefs/V1_chap12.html#tag_12_02
			for i := 1; i < len(arg); i++ {
				switch arg[i] {
				case 'f':
					force = true
				case 's':
					symlink = true
				case 'L':
					proc.Err("-L option is not implemented")
					return core.ExitInvalidArgs
				case 'P':
					proc.Err("-P option is not implemented")
					return core.ExitInvalidArgs
				default:
					proc.Err("Unrecognized flag: " + string([]byte{arg[i]}))
					return core.ExitInvalidArgs
				}
			}
		} else {
			files = append(files, arg)
		}
	}

	// TODO handle directories and more than 2 files
	if len(files) != 2 {
		proc.Err("Must provide 2 file paths")
		return core.ExitInvalidArgs
	}

	oldname := proc.ResolvePath(files[0])
	newname := proc.ResolvePath(files[1])

	if force {
		err := os.Remove(newname)
		if err != nil && !os.IsNotExist(err) {
			proc.Err(err.Error())
			return 2
		}
	}

	var err error
	if symlink {
		err = os.Symlink(oldname, newname)
	} else {
		err = os.Link(oldname, newname)
	}
	if err != nil {
		proc.Err("Error creating link: " + err.Error())
		return 3
	}
	return 0
}
