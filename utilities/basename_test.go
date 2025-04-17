package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestBasenameArgs(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	proc := ctx.Proc(Basename)
	if proc.Exec() == 0 {
		t.Error("basename should have returned non-zero with no arguments")
	}
	proc = ctx.Proc(Basename, "too", "many", "args")
	if proc.Exec() == 0 {
		t.Error("basename should have returned non-zero with > 2 arguments")
	}
}

func TestBasenameSuffix(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	proc := ctx.Proc(Basename, "file.txt", ".txt")
	if proc.Exec() != 0 {
		t.Error("basename should have returned non-zero with > 2 arguments")
	}
	if proc.Out() != "file\n" {
		t.Error("basename was incorrect for 'file.txt' with suffix '.txt': " + proc.Out())
	}

	proc = ctx.Proc(Basename, "/dir/file.txt", ".txt")
	if proc.Exec() != 0 {
		t.Error("basename should have returned non-zero with > 2 arguments")
	}
	if proc.Out() != "file\n" {
		t.Error("basename was incorrect for '/dir/file.txt' with suffix '.txt': " + proc.Out())
	}
}

func TestBasename(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	cases := map[string]string{
		"file.txt":       "file.txt",
		"dir/file.txt":   "file.txt",
		"/dir/file.txt":  "file.txt",
		"parent/child/":  "child",
		"/parent/child/": "child",
		"///":            "",
		"//":             "",
		"/":              "",
	}
	for k, v := range cases {
		proc := ctx.Proc(Basename, k)
		if proc.Exec() != 0 {
			t.Errorf("`basename %s` returned non-zero exit code", k)
		}
		v = v + "\n"
		if proc.Out() != v {
			t.Errorf("`basename %s` should have been `%s` but got `%s`", k, v, proc.Out())
		}
	}
}
