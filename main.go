// package simpleLogExtended
package main

import (
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

func (l *LogExtended) SetLogLevel(logLevel LogLevel) {
	if !logLevel.IsValid() {
		return
	}
	l.logLevel = logLevel
}

func (l *LogExtended) println(srcLogLevel LogLevel, prefix string, msg string) {
	if l.logLevel > srcLogLevel {
		return
	}

	l.Logger.Println(prefix + msg)
}

func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERR ", msg)
}

func NewLogExtended(out io.Writer, prefix string, flag int) *LogExtended {
	return &LogExtended{
		Logger:   log.New(out, prefix, flag),
		logLevel: LogLevelError,
	}
}

func main() {
	logger := NewLogExtended(os.Stderr, "", log.LstdFlags)
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Nothing to write")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
