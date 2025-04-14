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

	proc.SetInput("echo test\nexit\n")

	y := Sh(proc)

	if y != 0 {
		t.Errorf("shell exited with non-zero code: %d", y)
	}

	raw := strings.Trim(stdout.String(), "\x00")
	lines := strings.Split(raw, "\n")
	output := strings.Split(lines[0], "> ")[1]
	if strings.Trim(output, " ") != "test" {
		t.Errorf("shell did not echo text back: %d '%s'", len(raw), []byte(raw))
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
