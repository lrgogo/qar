package main

import (
	"app/db"
	"log"
	"net/http"
	"app/controller"
	"os"
)

func main() {
	//日志
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//数据库
	err = db.Connect()
	defer db.Close()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("mysql connect success")

	//接口路由
	controller.Load()

	//网页路由
	http.Handle("/", http.FileServer(http.Dir("static/html")))

	//服务器
	log.Println("server running", addr())
	err = http.ListenAndServe(addr(), nil)
	if err != nil {
		log.Println(err)
	}
}

func addr() string {
	return ":8080"
}
