package main

import (
	"fmt"
	"io"
	"os"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct {
	file *os.File
}

func (f *FileLogger) Log(message string) {
	_, err := f.file.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

type ConsoleLogger struct {
	out io.ReadWriter
}

func (c *ConsoleLogger) Log(message string) {
	fmt.Fprintln(c.out, message)
}

type LogSystem struct {
	logger Logger
}

type LogOption func(*LogSystem)

func NewLogSystem(options ...LogOption) *LogSystem {
	logSystem := &LogSystem{}

	for _, option := range options {
		option(logSystem)
	}
	return logSystem
}

func WithLogger(logger Logger) LogOption {
	return func(l *LogSystem) {
		l.logger = logger
	}
}

func (l *LogSystem) Log(message string) {
	if l.logger != nil {
		l.logger.Log(message)
	} else {
		fmt.Println("No logger set")
	}
}

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := &FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
}
