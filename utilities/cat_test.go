package utilities

import (
	"strings"
	"testing"

	"github.com/mackrorysd/gosix/core"
	"github.com/mackrorysd/gosix/tests"
)

func TestCat(t *testing.T) {
	proc, stdout, _ := core.TestProc()
	defer proc.CloseTest()

	tests.CreateFile(t, proc.ResolvePath("one.txt"), tests.TestString)
	tests.CreateFile(t, proc.ResolvePath("two.txt"), tests.TestString)

	proc.SetArgs("one.txt")

	y := Cat(proc)
	if y != 0 {
		t.Errorf("cat exited with non-zero code: %d", y)
	}

	output := stdout.String()
	if strings.Trim(output, "\x00") != tests.TestString {
		t.Errorf("cat did not print correct text: '%s'", []byte(output))
	}

	proc, stdout, _ = core.TestProc()
	defer proc.CloseTest()

	tests.CreateFile(t, proc.ResolvePath("one.txt"), tests.TestString)
	tests.CreateFile(t, proc.ResolvePath("two.txt"), tests.TestString)

	proc.SetArgs("one.txt", "two.txt")

	y = Cat(proc)
	if y != 0 {
		t.Errorf("cat exited with non-zero code: %d", y)
	}

	output = stdout.String()
	if strings.Trim(output, "\x00") != (tests.TestString + tests.TestString) {
		t.Errorf("cat did not print correct text: '%s'", []byte(output))
	}
}
