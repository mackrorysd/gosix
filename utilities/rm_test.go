package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestRm(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	ctx.InitFS(tests.TestFS)

	y := ctx.Proc(Rm, "top_file").Exec()
	if y != 0 {
		t.Error("rm of normal file should have succeeded")
	}

	proc := ctx.Proc(Rm, "top_dir")
	y = proc.Exec()
	if y == 0 {
		t.Error(proc.Err())
		t.Error("rm of a directory without -r should have failed")
		t.FailNow()
	}

	y = ctx.Proc(Rm, "-r", "top_dir").Exec()
	if y != 0 {
		t.Error("rm of a directory with -r should have succeeded")
	}
}
