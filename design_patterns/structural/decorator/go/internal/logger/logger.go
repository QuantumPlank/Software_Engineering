package logger

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Printf("%s\n", message)
}

type LoggerDecorator struct {
	Logger Logger
}

func (d *LoggerDecorator) Log(message string) {
	d.Logger.Log(message)
}

type TimestampDecorator struct {
	LoggerDecorator
}

func (t *TimestampDecorator) Log(message string) {
	t.LoggerDecorator.Log(fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), message))
}

type LevelDecorator struct {
	LoggerDecorator
	Level string
}

func (l *LevelDecorator) Log(message string) {
	l.LoggerDecorator.Log(fmt.Sprintf("[%s] %s", l.Level, message))
}

type FileDecorator struct {
	LoggerDecorator
	Filename string
}

func (f *FileDecorator) Log(message string) {
	f.LoggerDecorator.Log(fmt.Sprintf("[%s] %s", f.Filename, message))
}
