package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestFalse(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	y := ctx.Proc(False).Exec()
	if y == 0 {
		t.Error("false must exit with non-zero status code!")
	}
}
