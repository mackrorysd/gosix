package utilities

import (
	"github.com/mackrorysd/gosix/core"
)

// False is a utility that always returns a non-zero exit code, useful in shell
// programming. See [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/false.html
func False(proc core.Proc) int {
	return 1
}
