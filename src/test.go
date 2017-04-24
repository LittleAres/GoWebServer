package main

import (
	"path/filepath"
	"os"
	"fmt"
	//"flag"
	"strings"
	)
var Urls = []string{}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		abs_path := strings.Split(string(path), "/")
		file_name := abs_path[len(abs_path)-1]
		Urls = append(Urls, strings.Split(file_name, ".")[0])
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main(){
	getFilelist("./src/htmls/templates/")
	fmt.Print(Urls)
}