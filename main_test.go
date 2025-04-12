package main

import (
	"os"
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestCommand(t *testing.T) {
	proc, _, stderr := core.TestProc()
	defer proc.CloseTest()

	proc.SetArgs("ln", "-f", "source", "target")

	file, err := os.Create(proc.ResolvePath("source"))
	if err != nil || file.Close() != nil {
		t.FailNow()
	}

	y := _main(proc)

	if y != 0 {
		t.Errorf("`ln` exited with non-zero code: %d, %s", y, stderr.String())
	}
}
