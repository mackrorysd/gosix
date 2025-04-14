package core

import (
	"io"
	"path"
)

// Package core defines types and functionality expected to be common across
// other packages.

// ExitInvalidArgs is a standard exit status code for when CLI arguments are
// not valid, as determined while parsing.
const ExitInvalidArgs = 1

// ExitFileError is a standard exit status code for when specified paths are
// unusable because they do not exist or do not meet the requirements.
const ExitFileError = 2

// BufferSize is a standard number of bytes to use for buffers
const BufferSize = 4096

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

// Out is a convenience for writing strings to Stdout. It enforces a convention
// that text output end with a new-line character.
func (proc *Proc) Out(txt string) {
	if txt[len(txt)-1] != '\n' {
		txt += "\n"
	}
	proc.Stdout.Write([]byte(txt))
}

// Err is a convenience for writing strings to Stderr. It enforces a convention
// that text output end with a new-line character.
func (proc *Proc) Err(txt string) {
	if txt[len(txt)-1] != '\n' {
		txt += "\n"
	}
	proc.Stderr.Write([]byte(txt))
}
