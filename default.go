package glog

import (
	"os"
	"log"
)

var stdLog = New(LOG_INFO, os.Stderr, "", log.LstdFlags, 3)

func SetLevelFromString(level string) error{
	return stdLog.SetLevelFromString(level)
}

func SetLevel(level int) {
	stdLog.SetLevel(level)
}

func SetFlags(flags int) {
	stdLog.SetFlags(flags)
}

func Tracef(format string, v ...interface{}) {
	stdLog.Tracef(format, v)
}

func Traceln(v ...interface{}) {
	stdLog.Traceln(v)
}

func Debugf(format string, v ...interface{}) {
	stdLog.Debugf(format, v)
}

func Debugln(v ...interface{}) {
	stdLog.Debugln(v)
}

func Infof(format string, v ...interface{}) {
	stdLog.Infof(format, v)
}

func Infoln(v ...interface{}) {
	stdLog.Infoln(v)
}

func Warnf(format string, v ...interface{}) {
	stdLog.Warnf(format, v)
}

func Warnln(v ...interface{}) {
	stdLog.Warnln(v)
}

func Errf(format string, v ...interface{}) {
	stdLog.Errf(format, v)
}

func Errln(v ...interface{}) {
	stdLog.Errln(v)
}

func Critf(format string, v ...interface{}) {
	stdLog.Critf(format, v)
}

func Critln(v ...interface{}) {
	stdLog.Critln(v)
}

func Emergf(format string, v ...interface{}) {
	stdLog.Emergf(format, v)
}

func Emergln(v ...interface{}) {
	stdLog.Emergln(v)
}
