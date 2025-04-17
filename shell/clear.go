package shell

import (
	"github.com/mackrorysd/gosix/core"
	"github.com/mackrorysd/gosix/term"
)

func Clear(proc core.Proc) int {
	output := term.CleanScreen + term.CursorHome
	proc.Stdout.Write([]byte(output))
	return core.ExitSuccess
}
