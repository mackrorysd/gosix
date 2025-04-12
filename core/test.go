package core

import (
	"bytes"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
)

func TestProc() (Proc, *bytes.Buffer, *bytes.Buffer) {
	wd := path.Join(os.TempDir(), strconv.FormatUint(rand.Uint64(), 36))
	err := os.Mkdir(wd, 0700)
	if err != nil {
		panic("Failed to create test directory: " + err.Error())
	}

	stdout := bytes.NewBuffer(make([]byte, 1024))
	stderr := bytes.NewBuffer(make([]byte, 1024))

	proc := Proc{
		Args:   []string{},
		Cwd:    wd,
		Env:    make(map[string]string),
		Stdin:  strings.NewReader(""),
		Stdout: stdout,
		Stderr: stderr,
	}
	return proc, stdout, stderr
}

func (proc *Proc) SetArgs(args ...string) {
	proc.Args = args
}

func (proc *Proc) SetInput(input string) {
	proc.Stdin = strings.NewReader(input)
}

func (proc *Proc) CloseTest() {
	os.RemoveAll(proc.Cwd)
}
