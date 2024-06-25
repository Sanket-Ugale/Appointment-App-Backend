package utils

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs an info message
func LogInfo(message string) {
	InfoLogger.Println(message)
}

// LogError logs an error message
func LogError(err error) {
	if err != nil {
		ErrorLogger.Println(err)
	}
}
