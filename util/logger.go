package util

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "[INFO] ", log.Ltime)
	WarningLogger = log.New(os.Stdout, "[WARNING] ", log.Ltime)
	ErrorLogger = log.New(os.Stdout, "[ERROR] ", log.Ltime)
}

func Info(v any) {
	InfoLogger.Println(v)
}

func Warn(v any) {
	WarningLogger.Println(v)
}

func Error(v any) {
	ErrorLogger.Println(v)
}

func Fatal(v any) {
	ErrorLogger.Fatalln(v)
}
