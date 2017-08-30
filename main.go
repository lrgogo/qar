package main

import (
	"app/db"
	"log"
	"time"
	"net/http"
	"app/controller"
	"flag"
	"os"
)

func main() {
	//日志
	name := flag.String("log", "server.log", "log file name")
	flag.Parse()
	file, err := os.OpenFile(*name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
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

	//接口路由
	controller.Load()

	//网页路由
	http.Handle("/", http.FileServer(http.Dir("static/html")))

	//服务器
	log.Println(time.Now().Format("2006-01-02 15:04:05"), "server running", addr())
	err = http.ListenAndServe(addr(), nil)
	if err != nil {
		log.Println(err)
	}
}

func addr() string {
	return ":8080"
}
