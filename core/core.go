package core

import (
	"io"
	"path"
)

type Proc struct {
	Args   []string
	Cwd    string
	Env    map[string]string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (proc *Proc) ResolvePath(p string) string {
	if path.IsAbs(p) {
		return p
	}
	return path.Join(proc.Cwd, p)
}
