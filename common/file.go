package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// GetFiles 获取指定目录下的所有文件,包含子目录下的文件
func GetFiles(dirPth string, name string) (fileName string) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return ""
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if !fi.IsDir() {
			// 过滤指定格式
			ok := strings.HasPrefix(fi.Name(), name)
			if ok {
				return filepath.Join(dirPth, PthSep, fi.Name())
				//files = append(files, filepath.Join(dirPth,PthSep,fi.Name()))
			}
		}
	}
	return ""
}
