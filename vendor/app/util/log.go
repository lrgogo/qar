package util

import (
	"os"
	"log"
)

func InitLog() error {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return nil
}
