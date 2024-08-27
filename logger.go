package logger

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	LogLevelInfo  LogLevel = iota // default, always on
	LogLevelDebug                 // includes all info logs as well debug
)

type Logger struct {
	Level LogLevel
}

var (
	logger = NewLogger()
)

func NewLogger() *Logger {
	return &Logger{
		Level: LogLevelInfo,
	}
}

func (this *Logger) Logln(level LogLevel, text ...interface{}) {
	if level <= this.Level {
		err := log.Output(3, fmt.Sprintln(text...))
		if err != nil {
			log.Println("[ERROR] in Logln:", err)
		}
	}
}

func (this *Logger) Logf(level LogLevel, format string, text ...interface{}) {
	if level <= this.Level {
		err := log.Output(3, fmt.Sprintf(format, text...))
		if err != nil {
			log.Println("[ERROR] in Logf:", err)
		}
	}
}

/////////////////////////////////////////

func Info(text ...interface{}) {
	logger.Logln(LogLevelInfo, append([]interface{}{"[INFO]"}, text...)...)
}
func Infof(format string, text ...interface{}) {
	logger.Logf(LogLevelInfo, "[INFO] "+format, text...)
}
func Debug(text ...interface{}) {
	logger.Logln(LogLevelDebug, append([]interface{}{"[DEBUG]"}, text...)...)
}
func Debugf(format string, text ...interface{}) {
	logger.Logf(LogLevelDebug, "[DEBUG] "+format, text...)
}

func SetDebug() {
	logger.Level = LogLevelDebug
}

// Only need [INFO] and [DEBUG] https://dave.cheney.net/2015/11/05/lets-talk-about-logging
