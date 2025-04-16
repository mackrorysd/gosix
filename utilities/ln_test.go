package utilities

import (
	"os"
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func testLinkFile(t *testing.T, symlink bool, force bool) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	args := []string{}
	if symlink {
		args = append(args, "-s")
	}
	if force {
		args = append(args, "-f")
	}
	args = append(args, "source", "target")
	proc := ctx.Proc(Ln, args...)

	ctx.CreateFile("source", tests.TestString)

	if force {
		ctx.CreateFile("target", "")
	} else {
		ctx.CreateFile("target", "")
		y := proc.Exec()
		if y == 0 {
			t.Errorf("ln did not fail when target exists and link was not forced")
			t.FailNow()
		}
		proc = ctx.Proc(Ln, args...)
		ctx.DeleteFile("target")
	}

	y := proc.Exec()

	content, err := os.ReadFile(proc.ResolvePath("target"))
	if err != nil {
		t.FailNow()
	}
	if string(content) != tests.TestString {
		t.Errorf("Read unexpected file contents from link: %s", content)
	}

	if y != 0 {
		t.Errorf("ln exited with non-zero code: %d, %s", y, proc.Err())
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
