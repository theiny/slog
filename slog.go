package slog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const envLevel = "LOG_LEVEL"

const (
	debugLevel = "DEBUG"
	infoLevel  = "INFO"
	errorLevel = "ERROR"
)

var (
	debugColor = green
	infoColor  = yellow
	errColor   = red
)

var (
	green  = color("\033[1;32m%s\033[0m")
	yellow = color("\033[1;33m%s\033[0m")
	red    = color("\033[1;31m%s\033[0m")
)

type log struct {
	Level string
}

var l log

func init() {
	l.Level = os.Getenv(envLevel)
	SetLevel(l.Level)
}

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// SetLevel sets the log level based on DEBUG, INFO or ERROR values.
func SetLevel(lv string) {
	switch lv {
	case debugLevel:
		l.Level = debugLevel
	case infoLevel:
		l.Level = infoLevel
	case errorLevel:
		l.Level = errorLevel
	default:
		l.Level = debugLevel
	}
}

// Debug prints statements if log level is set to DEBUG.
func Debug(msg interface{}) {
	_, file, line, _ := runtime.Caller(1)
	if l.Level == debugLevel {
		fmt.Printf(debugColor("[%s] [%s] [%s line %d] %s \n"), time.Now().Format("2006-01-02 15:04:05"), debugLevel, file[strings.LastIndex(file, "/")+1:], line, msg)
	}
}

// Info prints statements if log level is set to DEBUG or INFO.
func Info(msg interface{}) {
	_, file, line, _ := runtime.Caller(1)
	if l.Level == debugLevel || l.Level == infoLevel {
		fmt.Printf(infoColor("[%s] [%s] [%s line %d] %s \n"), time.Now().Format("2006-01-02 15:04:05"), infoLevel, file[strings.LastIndex(file, "/")+1:], line, msg)
	}
}

// Error statements always print regardless of log level.
func Error(msg interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf(errColor("[%s] [%s] [%s line %d] %s \n"), time.Now().Format("2006-01-02 15:04:05"), errorLevel, file[strings.LastIndex(file, "/")+1:], line, msg)
}
