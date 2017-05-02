package views

import (
	"strings"
	"controllers"
	"os/exec"
)

func CoffeeCompile()(files []string){
	// 获取当前路径
	cmd := exec.Command("pwd")
	res, _ := cmd.Output()
	go_path := strings.TrimSpace(string(res))
	files, _ = controllers.WalkDir(go_path,".coffee")
	return files
}
