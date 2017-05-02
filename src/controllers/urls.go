package controllers

import (
	"strings"
	"path/filepath"
	"os"
	"os/exec"
	"path"
	//"fmt"
)

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			filenameWithSuffix := path.Base(filename)
			fileSuffix := path.Ext(filename)
			parentPath := strings.TrimSuffix(filename, filenameWithSuffix)
			fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
			// 先留着 虽然没什么用
			if suffix == ".HTML" {
				files = append(files, fileName)
			}
			if suffix == ".COFFEE"{
				command := "coffee -co " + strings.Replace(parentPath,"coffee","javascript",1) + " " + filename
				cmd := exec.Command("/bin/bash", "-c", command)
				_, err := cmd.Output()
				_ = err
			}
		}
		return nil
	})
	return files, err
}
