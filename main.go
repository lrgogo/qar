package main

import (
	"app/db"
	"log"
	"time"
	"net/http"
	"app/controller"
)

func main() {
	db.Connect()
	defer db.Close()

	controller.Load()

	http.Handle("/", http.FileServer(http.Dir("static/html")))

	log.Println("qar server run at ", time.Now())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
