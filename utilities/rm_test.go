package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/core"
	"github.com/mackrorysd/gosix/tests"
)

func TestRm(t *testing.T) {
	proc, _, _ := core.TestProc()
	defer proc.CloseTest()

	tests.InitFS(t, proc.Wd, tests.TestFS)

	proc.SetArgs("top_file")
	y := Rm(proc)
	if y != 0 {
		t.Error("rm of normal file should have succeeded")
	}

	proc.SetArgs("top_dir")
	y = Rm(proc)
	if y == 0 {
		t.Error("rm of a directory without -r should have failed")
		t.FailNow()
	}

	proc.SetArgs("-r", "top_dir")
	y = Rm(proc)
	if y != 0 {
		t.Error("rm of a directory with -r should have succeeded")
	}
}
