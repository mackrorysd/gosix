package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/core"
)

func TestTrue(t *testing.T) {
	proc, _, _ := core.TestProc()
	defer proc.CloseTest()

	y := True(proc)
	if y != 0 {
		t.Error("true must exit with zero status code!")
	}
}
