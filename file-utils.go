package fileUtils

import (
	"os"
	"path"
	"runtime"
)

// FileExists check if file exists
func FileExists(filename string) bool {
	// get file info
	finfo, err := getFileDir(filename)
	if err != nil {
		return false
	}

	// it's a dir not a file
	if finfo.IsDir() {
		return false
	}

	return true
}

// DirExists check if a directory exists
func DirExists(filename string) bool {
	// get file info
	finfo, err := getFileDir(filename)
	if err != nil {
		return false
	}

	// it's a file, not a directory
	if !finfo.IsDir() {
		return false
	}

	return true
}

// GetCurrentFile return the current file name
func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

// GetCurrentFileDir return the directory of the current file
func GetCurrentFileDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func getFileDir(path string) (os.FileInfo, error) {
	finfo, err := os.Stat(path)
	if err != nil {
		// no such file or dir
		return nil, err
	}

	return finfo, nil
}
