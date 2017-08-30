package main

import (
	"app/db"
	"log"
	"net/http"
	"app/controller"
	"app/util"
)

func main() {
	//日志
	err := util.InitLog()
	if err != nil {
		log.Println(err)
		return
	}

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
