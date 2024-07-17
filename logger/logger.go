package logger
import(
	"log"
	"os"
)
type Logger interface{
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}
type CustomLogger struct{
	debug *log.Logger
	info *log.Logger
	warn *log.Logger
	err *log.Logger
}

//func to initialize
func NewLogger() *CustomLogger{
	return &CustomLogger{
		debug: log.New(os.Stdout,"[DEBUG] : ",log.LstdFlags),
		info:log.New(os.Stdout,"[INFO] : ",log.LstdFlags),
		warn: log.New(os.Stdout,"[WARN] : ",log.LstdFlags),
		err: log.New(os.Stderr,"[ERROR] : ",log.LstdFlags),
	}
}

//implementing debug method for logger
func(l *CustomLogger) Debug(v ...interface{}){
	l.debug.Println(v...)
}
func(l *CustomLogger) Info(v ...interface{}){
	l.info.Println(v...)
}
func(l *CustomLogger) Warn(v ...interface{}){
	l.warn.Println(v...)
}
func(l *CustomLogger) Error(v ...interface{}){
	l.err.Println(v...)
}