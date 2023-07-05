package logger

import (
	"log"
	"os"
	"sync"
)

var (
	instanceLock sync.RWMutex
	instanceLog  Log
)

type Log interface {
	Debug(msgs ...interface{})
	Info(msgs ...interface{})
	Warn(msgs ...interface{})
	Error(msgs ...interface{})
	Fatal(msgs ...interface{})
}

type logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func NewLogger() Log {
	instanceLock.Lock()
	defer instanceLock.Unlock()
	if instanceLog == nil {
		instanceLog = logger{
			debugLogger:   log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
			infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
			warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
			errorLogger:   log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	}
	return instanceLog
}

func (self logger) Debug(msgs ...interface{}) {
	self.debugLogger.Println(msgs...)
}

func (self logger) Info(msgs ...interface{}) {
	self.infoLogger.Println(msgs...)
}

func (self logger) Warn(msgs ...interface{}) {
	self.warningLogger.Println(msgs...)
}
func (self logger) Error(msgs ...interface{}) {
	self.errorLogger.Println(msgs...)
}

func (self logger) Fatal(msgs ...interface{}) {
	self.errorLogger.Fatal(msgs...)
}
