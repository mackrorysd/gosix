package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestCat(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	ctx.CreateFile("one.txt", tests.TestString)
	ctx.CreateFile("two.txt", tests.TestString)

	proc := ctx.Proc(Cat, "one.txt")

	y := proc.Exec()
	if y != 0 {
		t.Errorf("cat exited with non-zero code: %d", y)
	}

	if proc.Out() != tests.TestString {
		t.Errorf("cat did not print correct text: '%s'", proc.Out())
	}

	proc = ctx.Proc(Cat, "one.txt", "two.txt")
	y = proc.Exec()
	if y != 0 {
		t.Errorf("cat exited with non-zero code: %d", y)
	}

	if proc.Out() != (tests.TestString + tests.TestString) {
		t.Errorf("cat did not print correct text: '%s'", proc.Out())
	}
}
