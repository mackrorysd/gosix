package core

import (
	"bytes"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
)

func TestProc(args []string, input string) (Proc, *bytes.Buffer, *bytes.Buffer) {
	// TODO consider the builder pattern for this
	wd := path.Join(os.TempDir(), strconv.FormatUint(rand.Uint64(), 36))
	os.Mkdir(wd, 0700)
	stdout := bytes.NewBuffer(make([]byte, 1024))
	stderr := bytes.NewBuffer(make([]byte, 1024))
	proc := Proc{
		Args:   args,
		Cwd:    wd,
		Env:    make(map[string]string),
		Stdin:  strings.NewReader(input),
		Stdout: stdout,
		Stderr: stderr,
	}
	return proc, stdout, stderr
}

func (proc *Proc) CloseTest() {
	os.RemoveAll(proc.Cwd)
}
