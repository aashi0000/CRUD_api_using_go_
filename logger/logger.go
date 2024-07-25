package logger

import (
	"fmt"
	"log"
	"os"
	"time"
	"runtime"
	//"path/filepath"
	//"strings"
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
//const projectDirName = "baic_api"
func NewLogger() *CustomLogger {
	return &CustomLogger{
		debug: log.New(os.Stdout, "", 0),
		info:  log.New(os.Stdout, "", 0),
		warn:  log.New(os.Stdout, "", 0),
		err:   log.New(os.Stderr, "", 0),
	}
}

// func formatLog(level string, color string, v ...interface{}) string {
// 	timestamp := time.Now().Format("2006-01-02 15:04:05")
// 	return fmt.Sprintf("%s[%s] : %s %s%s", color, level, timestamp, fmt.Sprint(v...), colorReset)
// }


func getCallingFunction() string {
	pc, _, _, _ := runtime.Caller(3) // Adjust the depth to get the correct calling function
	fn := runtime.FuncForPC(pc)
	funcName := "unknown"
	if fn != nil {
		funcName = fn.Name()
	}
	return funcName
}

func formatLog(level string, color string, v ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// _, file, line, ok := runtime.Caller(3) // Adjust the depth as needed
	// if ok {
	// 	projectRoot, err := getProjectRoot()
	// 	if err == nil {
	// 		relPath, err := filepath.Rel(projectRoot, file)
	// 		if err == nil && !strings.HasPrefix(relPath, "..") && !filepath.IsAbs(relPath) {
	// 			file = relPath
	// 		} else {
	// 			file = filepath.Base(file)
	// 		}
	// 	} else {
	// 		file = filepath.Base(file)
	// 	}

		funcName := getCallingFunction()

		return fmt.Sprintf("%s[%s] : %s | %s() | %s%s", color, level, timestamp,  funcName, fmt.Sprint(v...), colorReset)
	//}
	//return fmt.Sprintf("%s[%s] : %s %s%s", color, level, timestamp, fmt.Sprint(v...), colorReset)
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
