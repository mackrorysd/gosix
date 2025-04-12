package shell

import (
	"os"
	"strings"
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestEcho(t *testing.T) {
	proc, stdout, _ := core.TestProc()
	defer proc.CloseTest()

	proc.SetInput("test")

	y := Sh(proc)

	if y != 0 {
		t.Errorf("shell exited with non-zero code: %d", y)
	}

	output := stdout.String()
	if strings.Trim(output, "\x00\n") != "$ test" {
		t.Errorf("shell did not echo text back: '%s'", []byte(output))
	}
}

func TestCommand(t *testing.T) {
	proc, _, stderr := core.TestProc()
	defer proc.CloseTest()

	proc.SetArgs("-c", "ln", "-f", "source", "target")

	file, err := os.Create(proc.ResolvePath("source"))
	if err != nil || file.Close() != nil {
		t.FailNow()
	}

	y := Sh(proc)

	if y != 0 {
		t.Errorf("`ln` exited with non-zero code: %d, %s", y, stderr.String())
	}
}
