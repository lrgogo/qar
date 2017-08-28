package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func Connect()  {
	var err error
	db, err = sql.Open("mysql",
	"root:123456@tcp(127.0.0.1:3306)/qar")
	if err != nil {
		log.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("db connect success")
}

func Close()  {
	db.Close()
}