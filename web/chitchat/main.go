package main

import (
	. "chichat/config"
	. "chichat/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	startWebServer("8083")
}

// 通过指定端口启动 Web 服务器
func startWebServer(port string) {

	r := NewRouter()

	// 处理静态资源文件
	// 处理静态资源文件
	assets := http.FileServer(http.Dir(ViperConfig.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 应用路由器到 HTTP 服务器

	r.HandleFunc("/test", test)

	log.Println("Starting HTTP service at " + ViperConfig.App.Address)
	err := http.ListenAndServe(ViperConfig.App.Address, r) // 启动协程监听请求

	fmt.Println(err)
	if err != nil {
		log.Println("An error occured starting HTTP listener at" + ViperConfig.App.Address)
		log.Println("Error: " + err.Error())
	}
}

func test(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("nimp22"))
}
