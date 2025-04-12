package core

import (
	"io"
)

type Proc struct {
	Args   []string
	Cwd    string
	Env    map[string]string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
