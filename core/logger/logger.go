package logger

import (
	"log"
	"os"
)

var (
	Loger Log
)

type Log interface {
	Debug(logTag string, msg ...any)
	Info(logTag string, msgs ...any)
	Warn(logTag string, msgs ...any)
	Error(logTag string, msgs ...any)
	Fatal(logTag string, msgs ...any)
}

type logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func init() {
	Loger = logger{
		debugLogger:   log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime),
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime),
	}

}

func (self logger) Debug(logTag string, msgs ...any) {
	self.debugLogger.Println(msgs...)
}

func (self logger) Info(logTag string, msgs ...any) {
	self.infoLogger.Println(msgs...)
}

func (self logger) Warn(logTag string, msgs ...any) {
	self.warningLogger.Println(msgs...)
}
func (self logger) Error(logTag string, msgs ...any) {
	self.errorLogger.Println(msgs...)
}

func (self logger) Fatal(logTag string, msgs ...any) {
	self.errorLogger.Fatal(msgs...)
}
