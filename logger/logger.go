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
	loggerInstance = New()
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
// Public Setters and Getters
// ----------------------------------------------------------------------------

func SetLevel(level Level) *Logger { return loggerInstance.SetLevel(level) }
func (logger *Logger) SetLevel(level Level) *Logger {
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

func GetLevel() Level { return loggerInstance.GetLevel() }
func (logger *Logger) GetLevel() Level {
	return logger.level
}

func New() *Logger {
	return new(Logger)
}

// ----------------------------------------------------------------------------
// Public IsXXX() routines
// ----------------------------------------------------------------------------

func IsPanic() bool { return loggerInstance.IsPanic() }
func (logger *Logger) IsPanic() bool {
	return logger.isPanic
}

func IsFatal() bool { return loggerInstance.IsFatal() }
func (logger *Logger) IsFatal() bool {
	return logger.isFatal
}

func IsError() bool { return loggerInstance.IsError() }
func (logger *Logger) IsError() bool {
	return logger.isError
}

func IsWarn() bool { return loggerInstance.IsWarn() }
func (logger *Logger) IsWarn() bool {
	return logger.isWarn
}

func IsInfo() bool { return loggerInstance.IsInfo() }
func (logger *Logger) IsInfo() bool {
	return logger.isInfo
}

func IsDebug() bool { return loggerInstance.IsDebug() }
func (logger *Logger) IsDebug() bool {
	return logger.isDebug
}

func IsTrace() bool { return loggerInstance.IsTrace() }
func (logger *Logger) IsTrace() bool {
	return logger.isTrace
}

// ----------------------------------------------------------------------------
// Public XXX() routines for leveled logging.
// ----------------------------------------------------------------------------

// --- Trace ------------------------------------------------------------------

func Trace(v ...interface{}) *Logger {
	if loggerInstance.IsTrace() {
		loggerInstance.printf(LevelTraceName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *Logger) Trace(v ...interface{}) *Logger {
	if loggerInstance.isTrace {
		loggerInstance.printf(LevelTraceName, noFormat, v...)
	}
	return loggerInstance
}

func Tracef(format string, v ...interface{}) *Logger {
	if loggerInstance.IsTrace() {
		loggerInstance.printf(LevelTraceName, format, v...)
	}
	return loggerInstance
}

func (logger *Logger) Tracef(format string, v ...interface{}) *Logger {
	if loggerInstance.isTrace {
		loggerInstance.printf(LevelTraceName, format, v...)
	}
	return loggerInstance
}

// --- Debug ------------------------------------------------------------------

func Debug(v ...interface{}) *Logger {
	if loggerInstance.IsDebug() {
		loggerInstance.printf(LevelDebugName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *Logger) Debug(v ...interface{}) *Logger {
	if loggerInstance.isDebug {
		loggerInstance.printf(LevelDebugName, noFormat, v...)
	}
	return loggerInstance
}

func Debugf(format string, v ...interface{}) *Logger {
	if loggerInstance.IsDebug() {
		loggerInstance.printf(LevelDebugName, format, v...)
	}
	return loggerInstance
}

func (logger *Logger) Debugf(format string, v ...interface{}) *Logger {
	if loggerInstance.isDebug {
		loggerInstance.printf(LevelDebugName, format, v...)
	}
	return loggerInstance
}

// --- Info -------------------------------------------------------------------

func Info(v ...interface{}) *Logger {
	if loggerInstance.IsInfo() {
		loggerInstance.printf(LevelInfoName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *Logger) Info(v ...interface{}) *Logger {
	if loggerInstance.isInfo {
		loggerInstance.printf(LevelInfoName, noFormat, v...)
	}
	return loggerInstance
}

func Infof(format string, v ...interface{}) *Logger {
	if loggerInstance.IsInfo() {
		loggerInstance.printf(LevelInfoName, format, v...)
	}
	return loggerInstance
}

func (logger *Logger) Infof(format string, v ...interface{}) *Logger {
	if loggerInstance.isInfo {
		loggerInstance.printf(LevelInfoName, format, v...)
	}
	return loggerInstance
}

// --- Warn -------------------------------------------------------------------

func Warn(v ...interface{}) *Logger {
	if loggerInstance.IsWarn() {
		loggerInstance.printf(LevelWarnName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *Logger) Warn(v ...interface{}) *Logger {
	if loggerInstance.isWarn {
		loggerInstance.printf(LevelWarnName, noFormat, v...)
	}
	return loggerInstance
}

func Warnf(format string, v ...interface{}) *Logger {
	if loggerInstance.IsWarn() {
		loggerInstance.printf(LevelWarnName, format, v...)
	}
	return loggerInstance
}

func (logger *Logger) Warnf(format string, v ...interface{}) *Logger {
	if loggerInstance.isWarn {
		loggerInstance.printf(LevelWarnName, format, v...)
	}
	return loggerInstance
}

// --- Error ------------------------------------------------------------------

func Error(v ...interface{}) *Logger {
	if loggerInstance.IsError() {
		loggerInstance.printf(LevelErrorName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *Logger) Error(v ...interface{}) *Logger {
	if loggerInstance.isError {
		loggerInstance.printf(LevelErrorName, noFormat, v...)
	}
	return loggerInstance
}

func Errorf(format string, v ...interface{}) *Logger {
	if loggerInstance.IsError() {
		loggerInstance.printf(LevelErrorName, format, v...)
	}
	return loggerInstance
}

func (logger *Logger) Errorf(format string, v ...interface{}) *Logger {
	if loggerInstance.isError {
		loggerInstance.printf(LevelErrorName, format, v...)
	}
	return loggerInstance
}

// --- Fatal ------------------------------------------------------------------

func Fatal(v ...interface{}) *Logger {
	if loggerInstance.IsFatal() {
		loggerInstance.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return loggerInstance
}

func (logger *Logger) Fatal(v ...interface{}) *Logger {
	if loggerInstance.isFatal {
		loggerInstance.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return loggerInstance
}

func Fatalf(format string, v ...interface{}) *Logger {
	if loggerInstance.IsFatal() {
		loggerInstance.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return loggerInstance
}

func (logger *Logger) Fatalf(format string, v ...interface{}) *Logger {
	if loggerInstance.isFatal {
		loggerInstance.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return loggerInstance
}

// --- Panic ------------------------------------------------------------------

func Panic(v ...interface{}) *Logger {
	if loggerInstance.IsPanic() {
		loggerInstance.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return loggerInstance
}

func (logger *Logger) Panic(v ...interface{}) *Logger {
	if loggerInstance.isPanic {
		loggerInstance.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return loggerInstance
}

func Panicf(format string, v ...interface{}) *Logger {
	if loggerInstance.IsPanic() {
		loggerInstance.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return loggerInstance
}

func (logger *Logger) Panicf(format string, v ...interface{}) *Logger {
	if loggerInstance.isPanic {
		loggerInstance.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return loggerInstance
}
