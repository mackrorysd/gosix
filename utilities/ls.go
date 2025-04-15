package utilities

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/mackrorysd/gosix/core"
)

// Ls is a utility for listing files and directories. See
// [IEEE Std 1003.1-2024].
//
// [IEEE Std 1003.1-2024]: https://pubs.opengroup.org/onlinepubs/9799919799/utilities/ls.html
func Ls(proc core.Proc) int {
	// Right now this assumes -l and the current working directory

	entries, err := os.ReadDir(proc.ResolvePath(proc.Wd)) // Is this double resolution?
	if err != nil {
		proc.Err("Error reading directory: " + err.Error())
		return core.ExitFileError
	}
	sizeWidth := 1
	linksWidth := 1
	ownerWidth := 1
	groupWidth := 1
	for _, entry := range entries {
		links := 1 // TODO
		linksWidth = max(linksWidth, len(strconv.Itoa(links)))

		info, err := entry.Info()
		if err == nil {
			sizeWidth = max(sizeWidth, len(strconv.FormatInt(info.Size(), 10)))
		}

		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			linksWidth = max(linksWidth, len(strconv.FormatUint(uint64(stat.Nlink), 10)))
			ownerWidth = max(ownerWidth, len(strconv.FormatUint(uint64(stat.Uid), 10)))
			groupWidth = max(groupWidth, len(strconv.FormatUint(uint64(stat.Gid), 10)))
		}
	}
	longFormat := fmt.Sprintf("%%s %%%ds %%%ds %%%ds %%%dd %%s %%s", linksWidth, ownerWidth, groupWidth, sizeWidth)
	for _, entry := range entries {
		name := entry.Name()

		info, err := entry.Info()
		if err != nil {
			proc.Err("Error reading file " + name + ": " + err.Error())
		}

		mode := entry.Type().String()

		links := "?"
		owner := "?"
		group := "?"
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			links = strconv.FormatUint(uint64(stat.Nlink), 10)
			owner = strconv.FormatUint(uint64(stat.Uid), 10)
			group = strconv.FormatUint(uint64(stat.Gid), 10)
		}

		size := info.Size()
		mod := formatDate(info.ModTime())

		if info.Mode()&os.ModeSymlink > 0 {
			// Note: this resolves things relative to the current directory
			// This is different from how a shell would resolve things
			resolved, err := filepath.EvalSymlinks(name)
			if err != nil {
				proc.Err("Error resolving symlink: " + err.Error())
			}
			name += " -> " + resolved
		}

		line := fmt.Sprintf(longFormat, mode, links, owner, group, size, mod, name)
		proc.Out(line)
	}
	return core.ExitSuccess
}

const (
	// Format for files newer than 6 months
	newFormat = "Jan 02 15:04"

	// Format for files older than 6 months
	oldFormat = "Jan 02  2006"

	// Duration of a typical day
	day = 24 * time.Hour

	// Duration of a typical months
	month = 30 * day

	// Duration of the user of newFormat
	sixMonths = 6 * month
)

func formatDate(t time.Time) string {
	if time.Since(t) > sixMonths {
		return t.Format(oldFormat)
	}
	return t.Format(newFormat)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
