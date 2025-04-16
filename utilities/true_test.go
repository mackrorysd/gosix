package utilities

import (
	"testing"

	"github.com/mackrorysd/gosix/tests"
)

func TestTrue(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	y := ctx.Proc(True).Exec()
	if y != 0 {
		t.Error("true must exit with zero status code!")
	}

}
