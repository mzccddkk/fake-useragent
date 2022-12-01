package useragent

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// currentAbsolutePath returns the absolute path of the
// currently executing file. compatible go build and go run.
func currentAbsolutePath() (string, error) {
	dir, err := currentAbsolutePathForBuild()
	if err != nil {
		return "", err
	}
	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}
	if strings.Contains(dir, tmpDir) {
		return currentAbsolutePathForRun(), nil
	}
	return dir, nil
}

// currentAbsolutePathForBuild returns the absolute path of
// the currently executing file. for go build.
func currentAbsolutePathForBuild() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(filepath.Dir(execPath))
}

// currentAbsolutePathForRun returns the absolute path of
// the currently executing file. for go run.
func currentAbsolutePathForRun() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
