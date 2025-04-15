package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestFalse(t *testing.T) {
	proc, _, _ := core.TestProc()
	defer proc.CloseTest()

	y := False(proc)
	if y == 0 {
		t.Error("false must exit with zero status code!")
	}
}
