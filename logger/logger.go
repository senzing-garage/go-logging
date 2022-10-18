/*
Package helper ...
*/
package logger

import (
	"fmt"
	"log"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func init() {
	logger = New()
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (logger *Logger) printf(debugLevelName string, format string, v ...interface{}) *Logger {
	var message string
	calldepth := 3
	if format == noFormat {
		v := append(v, 0)
		copy(v[1:], v[0:])
		v[0] = debugLevelName + " "
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(debugLevelName+" "+format, v...)
	}
	log.Output(calldepth, message)
	return logger
}

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

func SetLevel(level Level) *Logger { return logger.SetLevel(level) }
func (logger *Logger) SetLevel(level Level) *Logger {
	logger.isPanic = level <= LevelPanic
	logger.isFatal = level <= LevelFatal
	logger.isError = level <= LevelError
	logger.isWarn = level <= LevelWarn
	logger.isInfo = level <= LevelInfo
	logger.isDebug = level <= LevelDebug
	logger.isTrace = level <= LevelTrace
	return logger
}

func New() *Logger {
	return new(Logger)
}

// ----------------------------------------------------------------------------
// Public IsXXX() routines
// ----------------------------------------------------------------------------

func IsPanic() bool { return logger.IsPanic() }
func (logger *Logger) IsPanic() bool {
	return logger.isPanic
}

func IsFatal() bool { return logger.IsFatal() }
func (logger *Logger) IsFatal() bool {
	return logger.isFatal
}

func IsError() bool { return logger.IsError() }
func (logger *Logger) IsError() bool {
	return logger.isError
}

func IsWarn() bool { return logger.IsWarn() }
func (logger *Logger) IsWarn() bool {
	return logger.isWarn
}

func IsInfo() bool { return logger.IsInfo() }
func (logger *Logger) IsInfo() bool {
	return logger.isInfo
}

func IsDebug() bool { return logger.IsDebug() }
func (logger *Logger) IsDebug() bool {
	return logger.isDebug
}

func IsTrace() bool { return logger.IsTrace() }
func (logger *Logger) IsTrace() bool {
	return logger.isTrace
}

// ----------------------------------------------------------------------------
// Public XXX() routines for leveled logging.
// ----------------------------------------------------------------------------

// --- Trace ------------------------------------------------------------------

func Trace(v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Trace(v ...interface{}) *Logger {
	if logger.isTrace {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

func Tracef(format string, v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

func (logger *Logger) Tracef(format string, v ...interface{}) *Logger {
	if logger.isTrace {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

// --- Debug ------------------------------------------------------------------

func Debug(v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Debug(v ...interface{}) *Logger {
	if logger.isDebug {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

func Debugf(format string, v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

func (logger *Logger) Debugf(format string, v ...interface{}) *Logger {
	if logger.isDebug {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

// --- Info -------------------------------------------------------------------

func Info(v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Info(v ...interface{}) *Logger {
	if logger.isInfo {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

func Infof(format string, v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

func (logger *Logger) Infof(format string, v ...interface{}) *Logger {
	if logger.isInfo {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

// --- Warn -------------------------------------------------------------------

func Warn(v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Warn(v ...interface{}) *Logger {
	if logger.isWarn {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

func Warnf(format string, v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}

func (logger *Logger) Warnf(format string, v ...interface{}) *Logger {
	if logger.isWarn {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}

// --- Error ------------------------------------------------------------------

func Error(v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Error(v ...interface{}) *Logger {
	if logger.isError {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

func Errorf(format string, v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

func (logger *Logger) Errorf(format string, v ...interface{}) *Logger {
	if logger.isError {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

// --- Fatal ------------------------------------------------------------------

func Fatal(v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

func (logger *Logger) Fatal(v ...interface{}) *Logger {
	if logger.isFatal {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

func Fatalf(format string, v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

func (logger *Logger) Fatalf(format string, v ...interface{}) *Logger {
	if logger.isFatal {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

// --- Panic ------------------------------------------------------------------

func Panic(v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

func (logger *Logger) Panic(v ...interface{}) *Logger {
	if logger.isPanic {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

func Panicf(format string, v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}

func (logger *Logger) Panicf(format string, v ...interface{}) *Logger {
	if logger.isPanic {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}
