package util

import (
	"os"
	"log"
	"io"
)

func InitLog() error {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	writers := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writers)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return nil
}
