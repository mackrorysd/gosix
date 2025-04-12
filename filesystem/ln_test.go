package filesystem

import (
	"os"
	"path"
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestHardlink(t *testing.T) {
	proc, _, stderr := core.TestProc([]string{"-f", "target", "source"}, "")
	defer proc.CloseTest()

	proc.Args[1] = path.Join(proc.Cwd, proc.Args[1])
	proc.Args[2] = path.Join(proc.Cwd, proc.Args[2])
	file, err := os.Create(proc.Args[1])
	if err != nil {
		t.FailNow()
	}
	err = file.Close()
	if err != nil {
		t.FailNow()
	}

	y := Ln(proc)

	if y != 0 {
		t.Errorf("`ln` exited with non-zero code: %d, %s", y, stderr.String())
	}
}
