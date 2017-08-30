package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func Connect() error {
	var err error
	db, err = sql.Open("mysql",
	"root:123456@tcp(127.0.0.1:3306)/qar")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	log.Println("mysql connect success")
	return nil
}

func Close()  {
	db.Close()
}