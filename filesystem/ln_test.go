package filesystem

import (
	"os"
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestHardlink(t *testing.T) {
	proc, _, stderr := core.TestProc()
	defer proc.CloseTest()

	proc.SetArgs("-f", "source", "target")

	file, err := os.Create(proc.ResolvePath("source"))
	if err != nil || file.Close() != nil {
		t.FailNow()
	}

	y := Ln(proc)

	if y != 0 {
		t.Errorf("`ln` exited with non-zero code: %d, %s", y, stderr.String())
	}
}
