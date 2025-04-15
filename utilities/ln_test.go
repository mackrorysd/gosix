package utilities

import (
	"os"
	"testing"

	"github.com/mackrorysd/gosix/core"
	"github.com/mackrorysd/gosix/tests"
)

func testLinkFile(t *testing.T, symlink bool, force bool) {
	proc, _, stderr := core.TestProc()
	defer proc.CloseTest()

	args := []string{}
	if symlink {
		args = append(args, "-s")
	}
	if force {
		args = append(args, "-f")
	}
	args = append(args, "source", "target")
	proc.SetArgs(args...)

	tests.CreateFile(t, proc.ResolvePath("source"), tests.TestString)

	if force {
		tests.CreateFile(t, proc.ResolvePath("target"), "")
	} else {
		tests.CreateFile(t, proc.ResolvePath("target"), "")
		y := Ln(proc)
		if y == 0 {
			t.Errorf("ln did not fail when target exists and link was not forced")
			t.FailNow()
		}
		tests.DeleteFile(t, proc.ResolvePath("target"))
	}

	y := Ln(proc)

	content, err := os.ReadFile(proc.ResolvePath("target"))
	if err != nil {
		t.FailNow()
	}
	if string(content) != tests.TestString {
		t.Errorf("Read unexpected file contents from link: %s", content)
	}

	if y != 0 {
		t.Errorf("ln exited with non-zero code: %d, %s", y, stderr.String())
	}
}

func TestHardlink(t *testing.T) {
	testLinkFile(t, false, false)
}

func TestForcedHardlink(t *testing.T) {
	testLinkFile(t, false, true)
}

func TestSymlink(t *testing.T) {
	testLinkFile(t, true, false)
}

func TestForcedSymlink(t *testing.T) {
	testLinkFile(t, true, true)
}

// TODO: test relative file paths
