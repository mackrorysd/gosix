package main

import (
	"os"
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestCommand(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	proc := ctx.Proc(_main, "ln", "-f", "source", "target")

	file, err := os.Create(proc.ResolvePath("source"))
	if err != nil || file.Close() != nil {
		t.FailNow()
	}

	y := proc.Exec()
	if y != 0 {
		t.Errorf("`ln` exited with non-zero code: %d, %s", y, proc.Err())
	}
}
