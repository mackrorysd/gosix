package utilities

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestTeeNewFile(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	proc := ctx.Proc(Tee, "new.txt")
	proc.SetInput(tests.TestString)
	y := proc.Exec()
	if y != 0 {
		t.Errorf("tee exited with non-zero code: %d", y)
	}
	if proc.Out() != tests.TestString {
		t.Errorf("tee did not print correct text: '%s'", proc.Out())
	}
	bytes, err := os.ReadFile(filepath.Join(proc.Wd, "new.txt"))
	if err != nil {
		t.Error("Failed to open new.txt")
		t.FailNow()
	}
	if string(bytes) != tests.TestString {
		t.Errorf("tee did not write correct text fo file: '%s'", string(bytes))
	}
}
