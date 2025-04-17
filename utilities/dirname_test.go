package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestDirnameArgs(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	proc := ctx.Proc(Dirname)
	if proc.Exec() == 0 {
		t.Error("dirname should have returned non-zero with no arguments")
	}
	proc = ctx.Proc(Dirname, "extra", "args")
	if proc.Exec() == 0 {
		t.Error("basename should have returned non-zero with > 1 arguments")
	}
}

func TestDirname(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	cases := map[string]string{
		"":                ".",
		"file.txt":        ".",
		"dir/file.txt":    "dir",
		"/dir/file.txt":   "/dir",
		"parent/child/":   "parent",
		"/parent/child/":  "/parent",
		"///":             "/",
		"//":              "/",
		"/":               "/",
		"/parent/../peer": "/parent/..",
	}
	for k, v := range cases {
		proc := ctx.Proc(Dirname, k)
		if proc.Exec() != 0 {
			t.Errorf("`dirname %s` returned non-zero exit code", k)
		}
		if proc.Out() != v+"\n" {
			t.Errorf("`dirname %s` should have been `%s` but got `%s`", k, v, proc.Out())
		}
	}
}
