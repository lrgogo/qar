package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB   //全局变量不能在函数中使用:=赋值

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
	return nil
}

func Close()  {
	db.Close()
}