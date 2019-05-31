package internal

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// WorkingDirectory 当前的工作路径
var WorkingDirectory *string

func init() {
	wd := GetWorkingDirectory()
	WorkingDirectory = &wd
}

// GetWorkingDirectory 获取当前工作路径
func GetWorkingDirectory() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}
