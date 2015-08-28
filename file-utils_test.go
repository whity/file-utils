package fileUtils

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func TestFileExists(t *testing.T) {
	file, _ := filepath.Abs("./t/file_exists.txt")
	if !FileExists(file) {
		t.Errorf("FileExists(%q)", file)
	}
}

func TestDirExists(t *testing.T) {
	dir, _ := filepath.Abs("./t")
	if !DirExists(dir) {
		t.Errorf("DirExists(%q)", dir)
	}
}

func TestGetCurrentFile(t *testing.T) {
	_, expected, _, _ := runtime.Caller(0)
	got := GetCurrentFile()
	if got != expected {
		t.Errorf("GetCurrentFile, got: %q, want: %q", got, expected)
	}
}

func TestGetCurrentFileDir(t *testing.T) {
	_, expected, _, _ := runtime.Caller(0)
	expected = path.Dir(expected)
	got := GetCurrentFileDir()
	if got != expected {
		t.Errorf("GetCurrentFileDir, got: %q, want: %q", got, expected)
	}
}
