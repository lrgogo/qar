package main

import (
	"app/db"
	"log"
	"time"
	"net/http"
	"app/controller"
	"fmt"
)

func main() {
	err := db.Connect()
	defer db.Close()
	if err != nil {
		log.Println(err)
		return
	}

	controller.Load()

	http.Handle("/", http.FileServer(http.Dir("static/html")))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "server running", addr())
	err = http.ListenAndServe(addr(), nil)
	if err != nil {
		log.Println(err)
	}
}

func addr() string {
	return ":8080"
}
