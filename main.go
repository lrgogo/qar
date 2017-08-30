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
	db.Connect()
	defer db.Close()

	controller.Load()

	http.Handle("/", http.FileServer(http.Dir("static/html")))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "server running", addr())
	log.Fatal(http.ListenAndServe(addr(), nil))
}

func addr() string {
	return ":8080"
}
