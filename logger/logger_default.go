/*
The logger package is a decorator over Go's log package.

The LoggerDefault implementation  provides a layer over go's log to
add Trace, Debug, Info, Warn, Error, Fatal, and Panic levels.

It also implements IsXxxx() functions
that can be used as guards
(https://en.wikipedia.org/wiki/Guard_(computer_science))
*/
package logger

import (
	"fmt"
	"log"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

/*
The LoggerDefault type is for logging messages based on the following levels:
TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC.
*/
type LoggerDefault struct {
	level   Level
	isDebug bool
	isError bool
	isFatal bool
	isInfo  bool
	isPanic bool
	isTrace bool
	isWarn  bool
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (logger *LoggerDefault) printf(debugLevelName string, format string, v ...interface{}) LoggerInterface {
	var message string
	calldepth := 3
	if format == noFormat {
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(format, v...)
	}
	log.Output(calldepth, message)
	return loggerInstance
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Debug() logs a DEBUG message.
func (logger *LoggerDefault) Debug(v ...interface{}) LoggerInterface {
	if logger.isDebug {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

// Debugf() logs a formatted DEBUG message.
func (logger *LoggerDefault) Debugf(format string, v ...interface{}) LoggerInterface {
	if logger.isDebug {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

// Error() logs a ERROR message.
func (logger *LoggerDefault) Error(v ...interface{}) LoggerInterface {
	if logger.isError {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

// Errorf() logs a formatted ERROR message.
func (logger *LoggerDefault) Errorf(format string, v ...interface{}) LoggerInterface {
	if logger.isError {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

// Fatal() logs a FATAL message.
func (logger *LoggerDefault) Fatal(v ...interface{}) LoggerInterface {
	if logger.isFatal {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

// Fatalf() logs a formatted FATAL message.
func (logger *LoggerDefault) Fatalf(format string, v ...interface{}) LoggerInterface {
	if logger.isFatal {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

// GetLogLevel() gets the logger instance logging level.
func (logger *LoggerDefault) GetLogLevel() Level {
	return logger.level
}

// GetLogLevelAsString() gets the logger instance logging level in string representation.
func (logger *LoggerDefault) GetLogLevelAsString() string {
	return LevelToTextMap[logger.level]
}

// Info() logs a INFO message.
func (logger *LoggerDefault) Info(v ...interface{}) LoggerInterface {
	if logger.isInfo {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

// Infof() logs a formatted INFO message.
func (logger *LoggerDefault) Infof(format string, v ...interface{}) LoggerInterface {
	if logger.isInfo {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

// IsDebug() returns true if the logger instance will log a DEBUG message.
func (logger *LoggerDefault) IsDebug() bool {
	return logger.isDebug
}

// IsError() returns true if the logger instance will log a ERROR message.
func (logger *LoggerDefault) IsError() bool {
	return logger.isError
}

// IsFatal() returns true if the logger instance will log a FATAL message.
func (logger *LoggerDefault) IsFatal() bool {
	return logger.isFatal
}

// IsInfo() returns true if the logger instance will log a INFO message.
func (logger *LoggerDefault) IsInfo() bool {
	return logger.isInfo
}

// IsPanic() returns true if the logger instance will log a PANIC message.
func (logger *LoggerDefault) IsPanic() bool {
	return logger.isPanic
}

// IsTrace() returns true if the logger instance will log a TRACE message.
func (logger *LoggerDefault) IsTrace() bool {
	return logger.isTrace
}

// IsWarn() returns true if the logger instance will log a WARN message.
func (logger *LoggerDefault) IsWarn() bool {
	return logger.isWarn
}

// Panic() logs a PANIC message.
func (logger *LoggerDefault) Panic(v ...interface{}) LoggerInterface {
	if logger.isPanic {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

// Panicf() logs a formatted PANIC message.
func (logger *LoggerDefault) Panicf(format string, v ...interface{}) LoggerInterface {
	if logger.isPanic {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}

// SetLogLevel() sets the logger instance logging level.
func (logger *LoggerDefault) SetLogLevel(level Level) LoggerInterface {
	logger.level = level
	logger.isPanic = level <= LevelPanic
	logger.isFatal = level <= LevelFatal
	logger.isError = level <= LevelError
	logger.isWarn = level <= LevelWarn
	logger.isInfo = level <= LevelInfo
	logger.isDebug = level <= LevelDebug
	logger.isTrace = level <= LevelTrace
	return logger
}

// SetLogLevelFromString() sets the logger instance logging level using a string representation.
func (logger *LoggerDefault) SetLogLevelFromString(levelString string) LoggerInterface {
	upperLevelString := strings.ToUpper(levelString)
	level, ok := TextToLevelMap[upperLevelString]
	if !ok {
		level = LevelPanic
	}
	logger.SetLogLevel(level)
	return logger
}

// Trace() logs a TRACE message.
func (logger *LoggerDefault) Trace(v ...interface{}) LoggerInterface {
	if logger.isTrace {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

// Tracef() logs a formatted TRACE message.
func (logger *LoggerDefault) Tracef(format string, v ...interface{}) LoggerInterface {
	if logger.isTrace {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

// Warn() logs a WARN message.
func (logger *LoggerDefault) Warn(v ...interface{}) LoggerInterface {
	if logger.isWarn {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

// Warnf() logs a formatted WARN message.
func (logger *LoggerDefault) Warnf(format string, v ...interface{}) LoggerInterface {
	if logger.isWarn {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}
