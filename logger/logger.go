package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

type CustomLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
}

// ANSI escape codes for colors
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func NewLogger() *CustomLogger {
	return &CustomLogger{
		debug: log.New(os.Stdout, "", 0),
		info:  log.New(os.Stdout, "", 0),
		warn:  log.New(os.Stdout, "", 0),
		err:   log.New(os.Stderr, "", 0),
	}
}

func formatLog(level string, color string, v ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s[%s] : %s %s%s", color, level, timestamp, fmt.Sprint(v...), colorReset)
}

func (l *CustomLogger) Debug(v ...interface{}) {
	l.debug.Println(formatLog("DEBUG", colorBlue, v...))
}

func (l *CustomLogger) Info(v ...interface{}) {
	l.info.Println(formatLog("INFO", colorGreen, v...))
}

func (l *CustomLogger) Warn(v ...interface{}) {
	l.warn.Println(formatLog("WARN", colorYellow, v...))
}

func (l *CustomLogger) Error(v ...interface{}) {
	l.err.Println(formatLog("ERROR", colorRed, v...))
}

var MyLogger = NewLogger()
