package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LogLevel int

const (
	_     LogLevel = iota // start enum at 1 instead of 0
	DEBUG                 // LogLevel DEBUG = 1
	INFO                  // LogLevel INFO = 2
	WARN                  // LogLevel WARN = 3
	ERROR                 // LogLevel ERROR = 4
	FATAL                 // LogLevel FATAL = 5
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return fmt.Sprintf("%s%sDEBUG", Green, Bold)
	case INFO:
		return fmt.Sprintf("%s%sINFO", Green, Bold)
	case WARN:
		return fmt.Sprintf("%s%sWARN", Yellow, Bold)
	case ERROR:
		return fmt.Sprintf("%s%sERROR", Purple, Bold)
	case FATAL:
		return fmt.Sprintf("%s%sFATAL", Red, Bold)
	default:
		return fmt.Sprintf("%s%sUNKNOWN", Cyan, Bold)
	}
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

type Logger struct {
	logger *log.Logger
	format string
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		format: "[%s] %s - %s",
	}
}

func (logger *Logger) Print(l LogLevel, message string) {
	date := time.Now().Local().Format(time.DateTime)
	logger.logger.Printf(logger.format, l, date, message)
}

func PrintLog(l LogLevel, title string, message string) {
	logger := &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		format: "%s%s %s:%d - %s%s%s %s",
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	filePath := filepath.Base(file)

	logger.logger.Printf(logger.format, l.String(), Reset, filePath, line, Blue, title, Reset, message)
}

func PrintError(l LogLevel, title string, err error) {
	logger := &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		format: "%s%s %s:%d - %s%s%s %s",
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	filePath := filepath.Base(file)

	logger.logger.Printf(logger.format, l.String(), Reset, filePath, line, Blue, title, Reset, err.Error())
}
