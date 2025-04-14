package utilities

import (
	"io"
	"os"

	"github.com/mackrorysd/gosix/core"
)

// Cat is a utility for reading files and printing them to stdout. See
// [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/cat.html
func Cat(proc core.Proc) int {
	files := []string{}

	for _, arg := range proc.Args {
		if arg[0] == '-' {
			if len(arg) == 1 {
				// It is implementation-defined if this means stdin, or a file named -
				// See https://pubs.opengroup.org/onlinepubs/9799919799/basedefs/V1_chap12.html#tag_12_02
				files = append(files, arg)
			}
			for i := 1; i < len(arg); i++ {
				switch arg[i] {
				case 'u':
					// Unbuffered input is not well-supported in Go
					proc.Err("-u option is not implemented")
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

	buffer := make([]byte, core.BufferSize)
	for _, path := range files {
		file, err := os.Open(proc.ResolvePath(path))
		if err != nil {
			proc.Err("Cannot open file " + path + ": " + err.Error())
			return core.ExitFileError
		}

		var n int
		for n, err = file.Read(buffer); n > 0 && err == nil; n, err = file.Read(buffer) {
			proc.Stdout.Write(buffer[0:n])
		}
		if err != io.EOF {
			proc.Err("Error reading file " + path + ": " + err.Error())
			file.Close()
			if err != nil {
				proc.Err("Error closing file " + path + ": " + err.Error())
			}
			return core.ExitFileError
		}
		err = file.Close()
		if err != nil {
			proc.Err("Error closing file " + path + ": " + err.Error())
			return core.ExitFileError
		}
	}

	return 0
}
