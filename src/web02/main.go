package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprint(w, "通过自己创建的处理器处理请求")
	fmt.Fprint(w, "通过自己配置服务器信息处理请求")
}

func main() {

	myHandler := MyHandler{}

	// http.Handle("/myhandler", &MyHandler)
	// 创建server结构
	server := http.Server{
		Addr:        ":4396",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}

	//创建路由
	//http.ListenAndServe(":4396", nil)
	server.ListenAndServe()
}
