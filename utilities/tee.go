package utilities

import (
	"io"
	"os"

	"github.com/mackrorysd/gosix/core"
)

// Tee writes its input to stdout as well as specified files. See
// [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/tee.html
func Tee(proc core.Proc) int {
	a := false
	paths := []string{}

	for _, arg := range proc.Args {
		if arg[0] == '-' {
			for i := 1; i < len(arg); i++ {
				switch arg[i] {
				case 'a':
					a = true
				case 'i':
					proc.Err("-i option is not implemented")
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

	files := []*os.File{}
	for _, path := range paths {
		flags := os.O_WRONLY | os.O_CREATE
		if a {
			flags |= os.O_APPEND
		}
		file, err := os.OpenFile(proc.ResolvePath(path), flags, 0644)
		if err != nil {
			proc.Err("Error opening file " + path + ": " + err.Error())
			return core.ExitFileError
		}
		files = append(files, file)
	}

	b := make([]byte, 1) // tee is not supposed to buffer
	for {
		n, err := proc.Stdin.Read(b)
		if err != nil && err != io.EOF {
			proc.Err("Error reading from stdin: " + err.Error())
			break
		}
		if err == io.EOF || n == 0 {
			break
		}
		_, err = proc.Stdout.Write(b)
		if err != nil {
			proc.Err("Error writing to stdout: " + err.Error())
			break
		}
		for _, file := range files {
			_, err := file.Write(b)
			if err != nil {
				proc.Err("Error writing to file: " + err.Error())
				break
			}
		}
	}

	for _, file := range files {
		err := file.Close()
		if err != nil {
			proc.Err("Error closing file: " + err.Error())
			return core.ExitFileError
		}
	}

	return core.ExitSuccess
}
