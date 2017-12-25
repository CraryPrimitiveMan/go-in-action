package main

import (
	"net/http"
	"log"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!") // 写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", hello) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
