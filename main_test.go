package main

import (
	"os"
	"path"
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestCommand(t *testing.T) {
	proc, _, stderr := core.TestProc([]string{"ln", "-f", "target", "link"}, "")
	defer proc.CloseTest()

	proc.Args[2] = path.Join(proc.Cwd, proc.Args[2])
	proc.Args[3] = path.Join(proc.Cwd, proc.Args[3])
	file, err := os.Create(proc.Args[2])
	if err != nil {
		t.FailNow()
	}
	err = file.Close()
	if err != nil {
		t.FailNow()
	}

	y := _main(proc)

	if y != 0 {
		t.Errorf("`ln` exited with non-zero code: %d, %s", y, stderr.String())
	}
}
