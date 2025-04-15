// Package tests implements helper functions for many common tasks in testing
// other packages
package tests

import (
	"os"
	"path"
	"testing"
)

const TestString = "The quick brown fox jumps over the lazy dog.\n"

var TestFS = map[string]interface{}{
	"top_file": TestString,
	"top_dir": map[string]interface{}{
		"middle_file": TestString,
		"bottom_dir": map[string]interface{}{
			"bottom_file": TestString,
		},
	},
	"empty_dir":  map[string]interface{}{},
	"empty_file": "",
}

func InitFS(t *testing.T, parent string, dir map[string]interface{}) {
	for k, v := range dir {
		child := path.Join(parent, k)
		if content, ok := v.(string); ok {
			CreateFile(t, child, content)
		}
		if subdir, ok := v.(map[string]interface{}); ok {
			CreateDir(t, child)
			InitFS(t, child, subdir)
		}
	}
}

func CreateDir(t *testing.T, path string) {
	err := os.Mkdir(path, 0700)
	if err != nil {
		t.FailNow()
	}
}

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
