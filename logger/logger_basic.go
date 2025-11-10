/*
The [BasicLogger] implementation of the [Logger] interface
provides a layer over go's log to
add Trace, Debug, Info, Warn, Error, Fatal, and Panic levels.
It also implements IsXxxx() functions that can be used as [guards].

[guards]: https://en.wikipedia.org/wiki/Guard_(computer_science)
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
The BasicLogger type is for logging messages based on the following levels:
TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC.
*/
type BasicLogger struct {
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
// Interface methods
// ----------------------------------------------------------------------------

// Debug() logs a DEBUG message.
func (logger *BasicLogger) Debug(v ...interface{}) Logger {
	if logger.isDebug {
		logger.print(LevelDebugName, v...)
	}

	return logger
}

// Debugf() logs a formatted DEBUG message.
func (logger *BasicLogger) Debugf(format string, v ...interface{}) Logger {
	if logger.isDebug {
		logger.printf(LevelDebugName, format, v...)
	}

	return logger
}

// Error() logs a ERROR message.
func (logger *BasicLogger) Error(v ...interface{}) Logger {
	if logger.isError {
		logger.print(LevelErrorName, v...)
	}

	return logger
}

// Errorf() logs a formatted ERROR message.
func (logger *BasicLogger) Errorf(format string, v ...interface{}) Logger {
	if logger.isError {
		logger.printf(LevelErrorName, format, v...)
	}

	return logger
}

// Fatal() logs a FATAL message.
func (logger *BasicLogger) Fatal(v ...interface{}) Logger {
	if logger.isFatal {
		logger.print(LevelFatalName, v...)
		log.Fatal("")
	}

	return logger
}

// Fatalf() logs a formatted FATAL message.
func (logger *BasicLogger) Fatalf(format string, v ...interface{}) Logger {
	if logger.isFatal {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}

	return logger
}

// GetLogLevel() gets the logger instance logging level.
func (logger *BasicLogger) GetLogLevel() Level {
	return logger.level
}

// GetLogLevelAsString() gets the logger instance logging level in string representation.
func (logger *BasicLogger) GetLogLevelAsString() string {
	return LevelToTextMap[logger.level]
}

// Info() logs a INFO message.
func (logger *BasicLogger) Info(v ...interface{}) Logger {
	if logger.isInfo {
		logger.print(LevelInfoName, v...)
	}

	return logger
}

// Infof() logs a formatted INFO message.
func (logger *BasicLogger) Infof(format string, v ...interface{}) Logger {
	if logger.isInfo {
		logger.printf(LevelInfoName, format, v...)
	}

	return logger
}

// IsDebug() returns true if the logger instance will log a DEBUG message.
func (logger *BasicLogger) IsDebug() bool {
	return logger.isDebug
}

// IsError() returns true if the logger instance will log a ERROR message.
func (logger *BasicLogger) IsError() bool {
	return logger.isError
}

// IsFatal() returns true if the logger instance will log a FATAL message.
func (logger *BasicLogger) IsFatal() bool {
	return logger.isFatal
}

// IsInfo() returns true if the logger instance will log a INFO message.
func (logger *BasicLogger) IsInfo() bool {
	return logger.isInfo
}

// IsPanic() returns true if the logger instance will log a PANIC message.
func (logger *BasicLogger) IsPanic() bool {
	return logger.isPanic
}

// IsTrace() returns true if the logger instance will log a TRACE message.
func (logger *BasicLogger) IsTrace() bool {
	return logger.isTrace
}

// IsWarn() returns true if the logger instance will log a WARN message.
func (logger *BasicLogger) IsWarn() bool {
	return logger.isWarn
}

// Panic() logs a PANIC message.
func (logger *BasicLogger) Panic(v ...interface{}) Logger {
	if logger.isPanic {
		logger.print(LevelPanicName, v...)
		log.Panic("")
	}

	return logger
}

// Panicf() logs a formatted PANIC message.
func (logger *BasicLogger) Panicf(format string, v ...interface{}) Logger {
	if logger.isPanic {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}

	return logger
}

// SetLogLevel() sets the logger instance logging level.
func (logger *BasicLogger) SetLogLevel(level Level) Logger {
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
func (logger *BasicLogger) SetLogLevelFromString(levelString string) Logger {
	upperLevelString := strings.ToUpper(levelString)

	level, ok := TextToLevelMap[upperLevelString]
	if !ok {
		level = LevelPanic
	}

	logger.SetLogLevel(level)

	return logger
}

// Trace() logs a TRACE message.
func (logger *BasicLogger) Trace(v ...interface{}) Logger {
	if logger.isTrace {
		logger.print(LevelTraceName, v...)
	}

	return logger
}

// Tracef() logs a formatted TRACE message.
func (logger *BasicLogger) Tracef(format string, v ...interface{}) Logger {
	if logger.isTrace {
		logger.printf(LevelTraceName, format, v...)
	}

	return logger
}

// Warn() logs a WARN message.
func (logger *BasicLogger) Warn(v ...interface{}) Logger {
	if logger.isWarn {
		logger.print(LevelWarnName, v...)
	}

	return logger
}

// Warnf() logs a formatted WARN message.
func (logger *BasicLogger) Warnf(format string, v ...interface{}) Logger {
	if logger.isWarn {
		logger.printf(LevelWarnName, format, v...)
	}

	return logger
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (logger *BasicLogger) print(
	debugLevelName string,
	messages ...interface{},
) {
	var message string

	_ = debugLevelName
	calldepth := 3

	message = fmt.Sprint(messages...)

	err := log.Output(calldepth, message)
	if err != nil {
		panic(err)
	}
}

func (logger *BasicLogger) printf(
	debugLevelName string,
	format string,
	messages ...interface{},
) {
	var message string

	_ = debugLevelName
	calldepth := 3

	if format == "" {
		message = fmt.Sprint(messages...)
	} else {
		message = fmt.Sprintf(format, messages...)
	}

	err := log.Output(calldepth, message)
	if err != nil {
		panic(err)
	}
}
