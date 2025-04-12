package core

import (
	"io"
	"path"
)

// Package core defines types and functionality expected to be common across
// other packages.

// A Proc provides all the context necessary for the interface between the OS
// and a process.
type Proc struct {
	Args   []string
	Wd     string
	Env    map[string]string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// ResolvePath ensures a path is an absolute, canonicalized path relative to the
// process's working directory. Paths that require no modification are simply
// returned.
func (proc *Proc) ResolvePath(p string) string {
	if path.IsAbs(p) {
		return p
	}
	return path.Join(proc.Wd, p)
}
