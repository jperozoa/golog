package golog

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const timeFormat = "2006-01-02 15:04:05.000000"
const humanLogFormat = "%s - %s - %s"

type logger struct {
	currLevel Level
	w         *os.File
}

var instance *logger
var mu sync.Mutex
var initialized int32

func GetInstance() *logger {
	if atomic.LoadInt32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &logger{LevelInfo, os.Stdout}
		atomic.StoreInt32(&initialized, 1)
	}

	return instance
}

// SetOutputByName func
func (l *logger) SetOutputByName(fName string) (*os.File, error) {
	logF, err := os.OpenFile(fName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	l.w = logF
	return l.w, err
}

// GetLogLevel func
func (l *logger) GetLogLevel() Level {
	return l.currLevel
}

// SetLogLevel func
func (l *logger) SetLogLevel(level Level) {
	l.currLevel = level
}

// Debug func
func Debug(msg string, a ...interface{}) {
	GetInstance().Log(LevelDebug, msg, a...)
}

// Info func
func Info(msg string, a ...interface{}) {
	GetInstance().Log(LevelInfo, msg, a...)
}

// Warning func
func Warning(msg string, a ...interface{}) {
	GetInstance().Log(LevelWarning, msg, a...)
}

// Error func
func Error(msg string, a ...interface{}) {
	GetInstance().Log(LevelError, msg, a...)
}

// Critical func
func Critical(msg string, a ...interface{}) {
	GetInstance().Log(LevelFatal, msg, a...)
	os.Exit(-1)
}

// Log func
func (l *logger) Log(level Level, msg string, a ...interface{}) {
	message := fmt.Sprintf(msg, a...)
	l.humanLog(level, message)
}

func (l *logger) humanLog(level Level, message string) {
	if level <= l.currLevel {
		c := Level2Color[level]
		l.w.Write(c)
		fmt.Fprintf(l.w, humanLogFormat+"\n",
			time.Now().Format(timeFormat),
			level_to_string[level],
			message)
		l.w.Write(reset)
	}
}
