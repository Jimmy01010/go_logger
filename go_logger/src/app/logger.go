package app

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
)

type Logger struct {
	Error *log.Logger
	Info *log.Logger
	Warn *log.Logger
}

// Log the date and time
var logFlags = log.Ldate | log.Ltime

// Tmethods

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Info.Printf(format, v...)
}

func (l *Logger) Errorf (format string, v ...interface{}){
	l.Error.Printf(format, v...)
}

func (l *Logger) Warnf (format string, v ...interface{}){
	l.Error.Printf(format, v...)
}


func NewSysLogger(name string) (*Logger, error) {
	errorStr := "Error initializing syslog Logger"
	var err error

	errorLog, err := syslog.New(syslog.LOG_ERR, name)
	if err != nil {
		err = fmt.Errorf("%s: %v", errorStr, err)
		return nil, err
	}

	infoLog, err := syslog.New(syslog.LOG_INFO, name)
	if err != nil {
		err = fmt.Errorf("%s: %v", errorStr, err)
		return nil, err
	}

	warnLog, err := syslog.New(syslog.LOG_WARNING, name)
	if err != nil {
		err = fmt.Errorf("%s: %v", errorStr, err)
		return nil, err
	}

	logger := new(Logger)
	logger.Error = log.New(errorLog, "", 0)
	logger.Info = log.New(infoLog, "", 0)
	logger.Warn = log.New(warnLog, "", 0)

	return logger, nil
}

func NewTerminaLogger(logPrefix string) *Logger{
	logger := new(Logger)

	prefix := logPrefix + ":"
	logger.Error = log.New(os.Stderr, prefix, logFlags)
	logger.Info = log.New(os.Stdout, prefix, logFlags)
	logger.Warn = log.New(os.Stdout, prefix, logFlags)

	return logger
}