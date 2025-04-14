package core

import (
	"bytes"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
)

// TestProc returns a Proc instance initialized for command functions to run
// in-memory as simple function calls for testing, with a temporary directory
// for testing. Other Set* functions can be used for further modification. Byte
// buffers representing stdout and stderr are returned so they can be checked
// directly by tests.
func TestProc() (proc Proc, stdout *bytes.Buffer, stderr *bytes.Buffer) {
	wd := path.Join(os.TempDir(), strconv.FormatUint(rand.Uint64(), 36))
	err := os.Mkdir(wd, 0700)
	if err != nil {
		panic("Failed to create test directory: " + err.Error())
	}

	stdout = bytes.NewBuffer(make([]byte, BufferSize))
	stderr = bytes.NewBuffer(make([]byte, BufferSize))

	proc = Proc{
		Args:   []string{},
		Wd:     wd,
		Env:    make(map[string]string),
		Stdin:  strings.NewReader(""),
		Stdout: stdout,
		Stderr: stderr,
	}
	return proc, stdout, stderr
}

// SetArgs is used to conveniently set the process args variadically.
func (proc *Proc) SetArgs(args ...string) {
	proc.Args = args
}

// SetInput takes the provided string and provides it to the process as it's
// stdin.
func (proc *Proc) SetInput(input string) {
	proc.Stdin = strings.NewReader(input)
}

// CloseTest should have a deferred call immediately after calling TestProc. It
// will delete the temporary directory and do any other required clean up.
func (proc *Proc) CloseTest() {
	os.RemoveAll(proc.Wd)
}
