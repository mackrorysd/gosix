package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestMkdir(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	y := ctx.Proc(Mkdir, "dir1/dir2").Exec()
	if y == 0 {
		t.Error("mkdir inside a non-existent directory should have failed")
	}

	y = ctx.Proc(Mkdir, "dir1").Exec()
	if y != 0 {
		t.Error("mkdir of a single directory should have succeeded")
		t.FailNow()
	}

	y = ctx.Proc(Mkdir, "dir1/dir2").Exec()
	if y != 0 {
		t.Error("mkdir of a child directory should have succeeded")
	}

	y = ctx.Proc(Mkdir, "-p", "dir3/dir4").Exec()
	if y != 0 {
		t.Error("mkdir -p nested directories should have succeeded")
		t.FailNow()
	}

	y = ctx.Proc(Mkdir, "dir3/dir4/dir5").Exec()
	if y != 0 {
		t.Error("mkdir inside of nested directories should have succeeded")
	}
}
