package shell

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestEcho(t *testing.T) {
	proc, stdout, _ := core.TestProc([]string{}, "test")
	defer proc.CloseTest()

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
	proc, _, stderr := core.TestProc([]string{"-c", "ln", "-f", "target", "source"}, "")
	defer proc.CloseTest()

	proc.Args[3] = path.Join(proc.Cwd, proc.Args[3])
	proc.Args[4] = path.Join(proc.Cwd, proc.Args[4])
	file, err := os.Create(proc.Args[4])
	if err != nil {
		t.FailNow()
	}
	err = file.Close()
	if err != nil {
		t.FailNow()
	}

	y := Sh(proc)

	if y != 0 {
		t.Errorf("`ln` exited with non-zero code: %d, %s", y, stderr.String())
	}
}
