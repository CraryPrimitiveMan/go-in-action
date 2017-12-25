package main

import (
	"fmt"
	"github.com/craryprimitiveman/go-in-action/ch3/config"
	"os"
	"os/signal"
	"syscall"
	"net/http"
	"log"
)

func main() {
	//fmt.Println(config.Config().DB.Server)
	//fmt.Println(config.Config().Owner.Name)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			config.ReloadConfig()
			log.Println("Reloaded config")
		}
	}()

	http.HandleFunc("/", hello) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", config.Config().Owner.Name) // 写入到w的是输出到客户端的
}
