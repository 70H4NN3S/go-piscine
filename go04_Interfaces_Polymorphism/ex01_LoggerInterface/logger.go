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
	File *os.File
}

type MultiLogger struct {
	Loggers *[]Logger
}

func (c ConsoleLogger) Logger(level, message string) {
	t := time.Now()
	fmt.Fprint(os.Stdout, t.Format("2006-01-02 15:04:05"),
		strings.ToUpper(level), message)
}
