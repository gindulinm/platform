package logging

import (
	"fmt"
	"log"
)

// Defining the data structure
type DefaultLogger struct {
	minLevel     LogLevel                 // minimum logging level
	loggers      map[LogLevel]*log.Logger // This is a map
	triggerPanic bool                     // Determines whether the application should panic (crash) when a logging error occurs.
}

// This function returns the minimum log level.
func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

// This feature logs a message if the importance level of that message is not lower than the minimum level that the journalist is set to.
func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

// It passes the Trace log level and the msg message to the write method to write it to the log.
func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

// Tracef writes a trace-level message to the logger using a formatted template.
func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(template, vals...))
}

// It passes the Debug log level and the msg message to the write method to write it to the log.
func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

// Debugf writes a trace-level message to the logger using a formatted template.
func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(template, vals...))
}

// It passes the Info log level and the msg message to the write method to write it to the log.
func (l *DefaultLogger) Info(msg string) {
	l.write(Information, msg)
}

// Infof writes a trace-level message to the logger using a formatted template.
func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
	l.write(Information, fmt.Sprintf(template, vals...))
}

// It passes the Warn log level and the msg message to the write method to write it to the log.
func (l *DefaultLogger) Warn(msg string) {
	l.write(Warning, msg)
}

// Warnf writes a trace-level message to the logger using a formatted template.
func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
	l.write(Warning, fmt.Sprintf(template, vals...))
}

// will be recorded in the journal and will cause panic if true
func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

// writes a formatted error message to the log with a log level of "Fatal" and, if necessary, causes a panic, stopping the program.
func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
	formattedMsg := fmt.Sprintf(template, vals...)
	l.write(Fatal, formattedMsg)
	if l.triggerPanic {
		panic(formattedMsg)
	}
}
