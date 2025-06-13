package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init() {
	Logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
}

func Info(msg string) {
	Logger.Println(msg)
}
