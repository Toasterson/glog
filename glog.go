package glog

import (
	"log"
	"io"
	"strings"
	"fmt"
)

const (
	//To Debug the Application gives a lot of information try grep
	LOG_TRACE = iota
	//To Analyse the application
	LOG_DEBUG = iota
	//Let the use know something
	LOG_INFO = iota
	//Warn the user he may have a problem
	LOG_WARN = iota
	//There was an error need to stop routine
	LOG_ERR = iota
	//Some Client has not gotten Data
	LOG_CRIT = iota
	//Application Panic that quits Even a Server
	LOG_EMERG = iota
)

type Logger struct {
	callDepth int
	Level int
	log   *log.Logger
}

func New(level int, output io.Writer, prefix string, flags int, depth int) *Logger {
	if depth == 0 {
		depth = 2
	}
	return &Logger{
		Level:     level,
		log:       log.New(output, prefix, flags),
		callDepth: depth,
	}
}

func (l *Logger) SetCallDepth(depth int) {
	l.callDepth = depth
}

func (l *Logger) SetLevelFromString(level string) error {
	if !strings.HasPrefix(level, "LOG_") {
		level = "LOG_" + level
	}
	level = strings.ToUpper(level)
	switch level {
	case "LOG_TRACE":
		l.SetLevel(LOG_TRACE)
		l.SetFlags(log.Llongfile | log.Ltime)
	case "LOG_DEBUG":
		l.SetLevel(LOG_DEBUG)
		l.SetFlags(log.Lshortfile | log.Ltime)
	case "LOG_INFO":
		l.SetLevel(LOG_INFO)
	case "LOG_WARN":
		l.SetLevel(LOG_WARN)
	case "LOG_ERR":
		l.SetLevel(LOG_ERR)
	case "LOG_CRIT":
		l.SetLevel(LOG_CRIT)
	case "LOG_EMERG":
		l.SetLevel(LOG_EMERG)
	default:
		fmt.Errorf("Cannot set Level to %s use trace|debug|info|warn|err|crit|emerg", level)
	}
	return nil
}

func (l *Logger) SetLevel(level int) {
	l.Level = level
}

func (l *Logger) SetFlags(flags int) {
	l.log.SetFlags(flags)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.Level == LOG_TRACE {
		l.log.Output(l.callDepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Traceln(v ...interface{}) {
	if l.Level == LOG_TRACE {
		l.log.Output(l.callDepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.Level <= LOG_DEBUG {
		l.log.Output(l.callDepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debugln(v ...interface{}) {
	if l.Level <= LOG_DEBUG {
		l.log.Output(l.callDepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.Level <= LOG_INFO {
		l.log.Output(l.callDepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Infoln(v ...interface{}) {
	if l.Level <= LOG_INFO {
		l.log.Output(l.callDepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.Level <= LOG_WARN {
		l.log.Output(l.callDepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warnln(v ...interface{}) {
	if l.Level <= LOG_WARN {
		l.log.Output(l.callDepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Errf(format string, v ...interface{}) {
	if l.Level <= LOG_ERR {
		l.log.Output(l.callDepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Errln(v ...interface{}) {
	if l.Level <= LOG_ERR {
		l.log.Output(l.callDepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Emergf(format string, v ...interface{}) {
	if l.Level <= LOG_EMERG {
		l.log.Output(l.callDepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Emergln(v ...interface{}) {
	if l.Level <= LOG_EMERG {
		l.log.Output(l.callDepth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Critf(format string, v ...interface{}) {
	if l.Level <= LOG_CRIT {
		l.log.Output(l.callDepth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Critln(v ...interface{}) {
	if l.Level <= LOG_CRIT {
		l.log.Output(l.callDepth, fmt.Sprintln(v...))
	}
}
