package views

import (
	"strings"
	"path/filepath"
	"os"
	"os/exec"
	"path"
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
			files = append(files, strings.TrimSuffix(filenameWithSuffix, fileSuffix))
		}
		return nil
	})
	return files, err
}

func GetTemplates()(files []string){
	// 获取当前路径
	cmd := exec.Command("pwd")
	res, _ := cmd.Output()
	go_path := strings.TrimSpace(string(res))
	files, _ = WalkDir(go_path,".html")
	return files
}
