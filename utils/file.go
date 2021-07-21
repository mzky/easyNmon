package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"log"

	"github.com/bingoohuang/statiq/fs"
)

//获取指定目录下的所有文件,包含子目录下的文件
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

func InitFile(sfs *fs.StatiqFS, tplFileName, fileName string) error {
	if _, err := os.Stat(fileName); err == nil {
		log.Printf("%s already exists, ignored!\n", fileName)
		return nil
	} else if os.IsNotExist(err) {
		// continue
	} else {
		return err
	}

	conf := sfs.Files[tplFileName].Data
	if err := ioutil.WriteFile(fileName, conf, 0755); err != nil {
		return err
	}

	log.Printf(fileName + " created!")
	return nil
}
