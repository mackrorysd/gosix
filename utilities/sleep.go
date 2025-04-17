package utilities

import (
	"strconv"
	"time"

	"github.com/mackrorysd/gosix/core"
)

// Sleep is a utility for pausing a specified number of seconds. See
// [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/sleep.html
func Sleep(proc core.Proc) int {
	if len(proc.Args) != 1 {
		proc.Err("sleep only accepts a number of seconds as a parameter")
		return core.ExitInvalidArgs
	}
	seconds, err := strconv.Atoi(proc.Args[0])
	if err != nil {
		proc.Err("Failed to parse seconds: " + err.Error())
		return core.ExitInvalidArgs
	}
	if seconds < 0 {
		proc.Err("Seconds must be non-negative")
		return core.ExitInvalidArgs
	}

	time.Sleep(time.Duration(seconds) * time.Second)
	return 0
}
