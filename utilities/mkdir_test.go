package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestMkdir(t *testing.T) {
	proc, _, _ := core.TestProc()
	defer proc.CloseTest()

	proc.SetArgs("dir1/dir2")
	y := Mkdir(proc)
	if y == 0 {
		t.Error("mkdir inside a non-existent directory should have failed")
	}

	proc.SetArgs("dir1")
	y = Mkdir(proc)
	if y != 0 {
		t.Error("mkdir of a single directory should have succeeded")
		t.FailNow()
	}

	proc.SetArgs("dir1/dir2")
	y = Mkdir(proc)
	if y != 0 {
		t.Error("mkdir of a child directory should have succeeded")
	}

	proc.SetArgs("-p", "dir3/dir4")
	y = Mkdir(proc)
	if y != 0 {
		t.Error("mkdir -p nested directories should have succeeded")
		t.FailNow()
	}

	proc.SetArgs("dir3/dir4/dir5")
	y = Mkdir(proc)
	if y != 0 {
		t.Error("mkdir inside of nested directories should have succeeded")
	}
}
