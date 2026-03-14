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
	Path string
}

type MultiLogger struct {
	Loggers *[]Logger
}

func (ConsoleLogger) Log(level, message string) {
	t := time.Now()
	fmt.Fprint(os.Stdout, t.Format("2006-01-02 15:04:05"),
		strings.ToUpper(level), message)
}

func (f FileLogger) Log(level, message string) {
	file, err := os.OpenFile(f.Path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		cons := ConsoleLogger{}
		cons.Log("ERROR", "Failed to create the file for the error message")
		return
	}
	defer file.Close()

	errorStr := fmt.Sprint(strings.ToUpper(level), message)
	file.WriteString(errorStr)
}
