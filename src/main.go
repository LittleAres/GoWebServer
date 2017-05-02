package main

import (
	"net/http"
	"views"
)

// server main method
func main() {
	// 编译 coffee 文件
	views.CoffeeCompile()
	// 初始化url容器
	views.CollectRouter()
	// 先把静态文件的处理扔到这里
	http.Handle("/javascript/", http.StripPrefix("/javascript/",
		http.FileServer(http.Dir("src/html/javascript"))))
	http.Handle("/css/", http.StripPrefix("/css/",
		http.FileServer(http.Dir("src/html/css"))))
	http.Handle("/templates/", http.StripPrefix("/templates/",
		http.FileServer(http.Dir("src/html/templates"))))
	http.Handle("/", views.RouterList)
	http.ListenAndServe(":8888", nil)
}