package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger interface {
	Log(level, message string)
}

type ConsoleLogger struct{}

type FileLogger struct {
	path string
}

type MultiLogger struct {
	loggers []Logger
}

func (ConsoleLogger) Log(level, message string) {
	fmt.Printf("%s [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"), strings.ToUpper(level), message)
}

func (f FileLogger) Log(level, message string) {
	file, err := os.OpenFile(f.path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		cons := ConsoleLogger{}
		cons.Log("ERROR", "Failed to create the file for the error message")
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "%s [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"), strings.ToUpper(level), message)
}
