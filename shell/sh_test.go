package shell

import (
	"os"
	"strings"
	"testing"

	"github.com/mackrorysd/gosix/term"
	"github.com/mackrorysd/gosix/tests"
)

func TestEcho(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	proc := ctx.Proc(Sh)
	proc.SetInput("echo test\nexit\n")

	y := proc.Exec()
	if y != 0 {
		t.Errorf("shell exited with non-zero code: %d", y)
	}

	raw := proc.Out()
	text := term.StripEscapeCodes(raw)
	lines := strings.Split(text, "\n")
	output := strings.Split(lines[0], "> ")[1]
	if strings.Trim(output, " ") != "test" {
		t.Errorf("shell did not echo text back: '%s'", raw)
	}
}

func TestCommand(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	proc := ctx.Proc(Sh, "-c", "ln", "-f", "source", "target")

	file, err := os.Create(proc.ResolvePath("source"))
	if err != nil || file.Close() != nil {
		t.FailNow()
	}

	y := proc.Exec()

	if y != 0 {
		t.Errorf("ln exited with non-zero code: %d, %s", y, proc.Err())
	}
}
