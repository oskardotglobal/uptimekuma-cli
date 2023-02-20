package util

import "log"

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(nil, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(nil, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(nil, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(v ...any) {
	InfoLogger.Println(v)
}

func Warn(v ...any) {
	WarningLogger.Println(v)
}

func Error(v ...any) {
	ErrorLogger.Println(v)
}

func Fatal(v ...any) {
	ErrorLogger.Fatalln(v)
}
