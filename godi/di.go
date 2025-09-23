package godi

import (
	"fmt"
)

type Logger interface {
	log(message string)
}

type FileLogger struct {

}

type DBLogger struct {

}

func (f FileLogger) log (message string) {
	fmt.Printf("Logging to file: %s\n", message)
}

func (db DBLogger) log(message string) {
	fmt.Printf("Logging to db: %s\n", message)
}

func consumer (logger Logger) {
	logger.log("info")
}

func Run () {
	fileLogger := FileLogger{}
	consumer(fileLogger)

	dbLogger := DBLogger{}
	consumer(dbLogger)
}