// Package tests implements helper functions for many common tasks in testing
// other packages
package tests

import (
	"os"
	"testing"
)

const TestString = "The quick brown fox jumps over the lazy dog.\n"

func CreateFile(t *testing.T, path string, contents string) {
	file, err := os.Create(path)
	if err != nil {
		t.FailNow()
	}
	_, err = file.WriteString(contents)
	if err != nil {
		t.FailNow()
	}
	if err = file.Close(); err != nil {
		t.FailNow()
	}
}

func DeleteFile(t *testing.T, path string) {
	err := os.Remove(path)
	if err != nil {
		t.FailNow()
	}
}
