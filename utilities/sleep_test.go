package utilities

import (
	"testing"
	"time"

	"github.com/mackrorysd/gosix/tests"
)

func TestSleep(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	before := time.Now()
	y := ctx.Proc(Sleep, "1").Exec()
	if y != 0 {
		t.Error("sleep 1 should have succeeded")
	}
	after := time.Now()

	delta := after.Sub(before)
	if delta < time.Second || delta > 2*time.Second {
		t.Errorf("sleep 1 allowed %f seconds to pass", after.Sub(before).Seconds())
	}

	y = ctx.Proc(Sleep).Exec()
	if y == 0 {
		t.Error("sleep should have required a parameter")
	}

	y = ctx.Proc(Sleep, "1", "2").Exec()
	if y == 0 {
		t.Error("sleep should have required a single parameter")
	}

	y = ctx.Proc(Sleep, "-1").Exec()
	if y == 0 {
		t.Error("sleep should have required a positive parameter")
	}

	y = ctx.Proc(Sleep, "a").Exec()
	if y == 0 {
		t.Error("sleep should have required a numerical parameter")
	}
}
