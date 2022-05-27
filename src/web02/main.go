package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "通过自己创建的处理器处理请求")
}

func main() {

	MyHandler := MyHandler{}

	http.Handle("/myhandler", &MyHandler)

	//创建路由
	http.ListenAndServe(":4396", nil)
}
