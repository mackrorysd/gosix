package utilities

import (
	"github.com/mackrorysd/gosix/core"
)

// True is a utility that always returns a zero exit code, useful in shell
// programming. See [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/true.html
func True(proc core.Proc) int {
	return 0
}
