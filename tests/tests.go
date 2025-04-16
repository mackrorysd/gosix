// Package tests implements helper functions for many common tasks in testing
// other packages
package tests

import (
	"bytes"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/mackrorysd/gosix/core"
)

// A TestContext manages a temporary directory for filesystem operations in
// tests, and provides helper functions for common tasks.
type TestContext struct {
	T   *testing.T
	Dir string
}

// NewTestContext initializes a random temporary directory. A call should be
// followed immediately by a deferred call to TestContext.Close().
func NewTestContext(t *testing.T) *TestContext {
	dir := path.Join(os.TempDir(), strconv.FormatUint(rand.Uint64(), 36))
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		t.Errorf("Failed to create test directory: %s", err.Error())
		t.FailNow()
	}

	return &TestContext{
		T:   t,
		Dir: dir,
	}
}

// Close should have a deferred call immediately after calling TestProc. It will
// delete the temporary directory.
func (c *TestContext) Close() {
	os.RemoveAll(c.Dir)
}

const TestString = "The quick brown fox jumps over the lazy dog.\n"

var TestFS = map[string]interface{}{
	"top_file": TestString,
	"top_dir": map[string]interface{}{
		"middle_file": TestString,
		"bottom_dir": map[string]interface{}{
			"bottom_file": TestString,
		},
	},
	"empty_dir":  map[string]interface{}{},
	"empty_file": "",
}

func (c *TestContext) InitFS(dir map[string]interface{}) {
	c.initFS(c.Dir, dir)
}

func (c *TestContext) initFS(parent string, dir map[string]interface{}) {
	for k, v := range dir {
		if content, ok := v.(string); ok {
			c.CreateFile(k, content)
		}
		if subdir, ok := v.(map[string]interface{}); ok {
			c.CreateDir(k)
			c.initFS(k, subdir)
		}
	}
}

func (c *TestContext) CreateDir(dir string) {
	err := os.Mkdir(path.Join(c.Dir, dir), 0700)
	if err != nil {
		c.T.Error(err.Error())
		c.T.FailNow()
	}
}

func (c *TestContext) CreateFile(dir string, contents string) {
	file, err := os.Create(path.Join(c.Dir, dir))
	if err != nil {
		c.T.Error(err.Error())
		c.T.FailNow()
	}
	_, err = file.WriteString(contents)
	if err != nil {
		c.T.Error(err.Error())
		c.T.FailNow()
	}
	if err = file.Close(); err != nil {
		c.T.Error(err.Error())
		c.T.FailNow()
	}
}

func (c *TestContext) DeleteFile(file string) {
	err := os.Remove(path.Join(c.Dir, file))
	if err != nil {
		c.T.Error(err.Error())
		c.T.FailNow()
	}
}

// A TestProc is an extension of Proc with helper functions to make invocation
// from Go code more idiomatic.
type TestProc struct {
	core.Proc
	Fn     func(proc core.Proc) int
	OutBuf *bytes.Buffer
	ErrBuf *bytes.Buffer
}

// Proc returns a Proc instance initialized for command functions to run
// in-memory as simple function calls for testing, with a temporary directory
// for testing. Other Set* functions can be used for further modification. Byte
// buffers representing stdout and stderr are returned so they can be checked
// directly by tests.
func (c *TestContext) Proc(fn func(proc core.Proc) int, args ...string) *TestProc {
	stdout := bytes.NewBuffer(make([]byte, core.BufferSize))
	stderr := bytes.NewBuffer(make([]byte, core.BufferSize))

	return &TestProc{
		Fn:     fn,
		OutBuf: stdout,
		ErrBuf: stderr,
		Proc: core.Proc{
			Args:   args,
			Wd:     c.Dir,
			Env:    make(map[string]string),
			Stdin:  strings.NewReader(""),
			Stdout: stdout,
			Stderr: stderr,
		},
	}
}

// SetInput takes the provided string and provides it to the process as it's
// stdin.
func (proc *TestProc) SetInput(input string) {
	proc.Stdin = strings.NewReader(input)
}

func (proc *TestProc) Exec() int {
	return proc.Fn(proc.Proc)
}

func (proc *TestProc) Out() string {
	return strings.Trim(proc.OutBuf.String(), "\x00")
}

func (proc *TestProc) Err() string {
	return strings.Trim(proc.ErrBuf.String(), "\x00")
}
