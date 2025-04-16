package utilities

import (
	"regexp"
	"testing"
	"time"

	"github.com/mackrorysd/gosix/tests"
)

func TestLs(t *testing.T) {
	ctx := tests.NewTestContext(t)
	defer ctx.Close()

	ctx.InitFS(tests.TestFS)
	y := ctx.Proc(Ls).Exec()
	if y != 0 {
		t.Errorf("ls exited with non-zero code: %d", y)
	}
}

func TestDateFormat(t *testing.T) {
	oldDate := time.Unix(0, 0).UTC()
	newDate := time.Now()

	oldFormat := formatDate(oldDate)
	if oldFormat != "Jan 01  1970" {
		t.Errorf("Old date in unexpected format: %s", oldFormat)
	}

	newFormat := formatDate(newDate)
	matched, err := regexp.Match(`[A-Z][a-z][a-z] \d\d \d\d:\d\d`, []byte(newFormat))
	if err != nil {
		t.Errorf("Error matching date format: %s", err.Error())
		t.FailNow()
	}
	if !matched {
		t.Errorf("New date in unexpected format: %s", newFormat)
	}

	if len(oldFormat) != len(newFormat) {
		t.Errorf("Date formats must be a consistent length")
	}
}
