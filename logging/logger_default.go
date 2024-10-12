package logging

import (
	"log"
	// "fmt"
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
