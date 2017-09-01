package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis"
)

var db *sql.DB   //全局变量不能在函数中使用:=赋值

func ConnectDB() error {
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

func CloseDB()  {
	db.Close()
}

var client *redis.Client

func InitRedis() error {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func CloseRedis()  {
	client.Close()
}