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

func (logger *LoggerImpl) printf(debugLevelName string, format string, v ...interface{}) LoggerInterface {
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
	return loggerInstance
}

// ----------------------------------------------------------------------------
// Public Setters and Getters
// ----------------------------------------------------------------------------

func SetLevel(level Level) LoggerInterface { return loggerInstance.SetLevel(level) }
func (logger *LoggerImpl) SetLevel(level Level) LoggerInterface {
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
func (logger *LoggerImpl) GetLevel() Level {
	return logger.level
}

func New() *LoggerImpl {
	return new(LoggerImpl)
}

// ----------------------------------------------------------------------------
// Public IsXXX() routines
// ----------------------------------------------------------------------------

func IsPanic() bool { return loggerInstance.IsPanic() }
func (logger *LoggerImpl) IsPanic() bool {
	return logger.isPanic
}

func IsFatal() bool { return loggerInstance.IsFatal() }
func (logger *LoggerImpl) IsFatal() bool {
	return logger.isFatal
}

func IsError() bool { return loggerInstance.IsError() }
func (logger *LoggerImpl) IsError() bool {
	return logger.isError
}

func IsWarn() bool { return loggerInstance.IsWarn() }
func (logger *LoggerImpl) IsWarn() bool {
	return logger.isWarn
}

func IsInfo() bool { return loggerInstance.IsInfo() }
func (logger *LoggerImpl) IsInfo() bool {
	return logger.isInfo
}

func IsDebug() bool { return loggerInstance.IsDebug() }
func (logger *LoggerImpl) IsDebug() bool {
	return logger.isDebug
}

func IsTrace() bool { return loggerInstance.IsTrace() }
func (logger *LoggerImpl) IsTrace() bool {
	return logger.isTrace
}

// ----------------------------------------------------------------------------
// Public XXX() routines for leveled logging.
// ----------------------------------------------------------------------------

// --- Trace ------------------------------------------------------------------

func Trace(v ...interface{}) LoggerInterface {
	if loggerInstance.IsTrace() {
		loggerInstance.printf(LevelTraceName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Trace(v ...interface{}) LoggerInterface {
	if loggerInstance.isTrace {
		loggerInstance.printf(LevelTraceName, noFormat, v...)
	}
	return loggerInstance
}

func Tracef(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsTrace() {
		loggerInstance.printf(LevelTraceName, format, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Tracef(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.isTrace {
		loggerInstance.printf(LevelTraceName, format, v...)
	}
	return loggerInstance
}

// --- Debug ------------------------------------------------------------------

func Debug(v ...interface{}) LoggerInterface {
	if loggerInstance.IsDebug() {
		loggerInstance.printf(LevelDebugName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Debug(v ...interface{}) LoggerInterface {
	if loggerInstance.isDebug {
		loggerInstance.printf(LevelDebugName, noFormat, v...)
	}
	return loggerInstance
}

func Debugf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsDebug() {
		loggerInstance.printf(LevelDebugName, format, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Debugf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.isDebug {
		loggerInstance.printf(LevelDebugName, format, v...)
	}
	return loggerInstance
}

// --- Info -------------------------------------------------------------------

func Info(v ...interface{}) LoggerInterface {
	if loggerInstance.IsInfo() {
		loggerInstance.printf(LevelInfoName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Info(v ...interface{}) LoggerInterface {
	if loggerInstance.isInfo {
		loggerInstance.printf(LevelInfoName, noFormat, v...)
	}
	return loggerInstance
}

func Infof(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsInfo() {
		loggerInstance.printf(LevelInfoName, format, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Infof(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.isInfo {
		loggerInstance.printf(LevelInfoName, format, v...)
	}
	return loggerInstance
}

// --- Warn -------------------------------------------------------------------

func Warn(v ...interface{}) LoggerInterface {
	if loggerInstance.IsWarn() {
		loggerInstance.printf(LevelWarnName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Warn(v ...interface{}) LoggerInterface {
	if loggerInstance.isWarn {
		loggerInstance.printf(LevelWarnName, noFormat, v...)
	}
	return loggerInstance
}

func Warnf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsWarn() {
		loggerInstance.printf(LevelWarnName, format, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Warnf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.isWarn {
		loggerInstance.printf(LevelWarnName, format, v...)
	}
	return loggerInstance
}

// --- Error ------------------------------------------------------------------

func Error(v ...interface{}) LoggerInterface {
	if loggerInstance.IsError() {
		loggerInstance.printf(LevelErrorName, noFormat, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Error(v ...interface{}) LoggerInterface {
	if loggerInstance.isError {
		loggerInstance.printf(LevelErrorName, noFormat, v...)
	}
	return loggerInstance
}

func Errorf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsError() {
		loggerInstance.printf(LevelErrorName, format, v...)
	}
	return loggerInstance
}

func (logger *LoggerImpl) Errorf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.isError {
		loggerInstance.printf(LevelErrorName, format, v...)
	}
	return loggerInstance
}

// --- Fatal ------------------------------------------------------------------

func Fatal(v ...interface{}) LoggerInterface {
	if loggerInstance.IsFatal() {
		loggerInstance.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return loggerInstance
}

func (logger *LoggerImpl) Fatal(v ...interface{}) LoggerInterface {
	if loggerInstance.isFatal {
		loggerInstance.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return loggerInstance
}

func Fatalf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsFatal() {
		loggerInstance.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return loggerInstance
}

func (logger *LoggerImpl) Fatalf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.isFatal {
		loggerInstance.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return loggerInstance
}

// --- Panic ------------------------------------------------------------------

func Panic(v ...interface{}) LoggerInterface {
	if loggerInstance.IsPanic() {
		loggerInstance.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return loggerInstance
}

func (logger *LoggerImpl) Panic(v ...interface{}) LoggerInterface {
	if loggerInstance.isPanic {
		loggerInstance.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return loggerInstance
}

func Panicf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsPanic() {
		loggerInstance.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return loggerInstance
}

func (logger *LoggerImpl) Panicf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.isPanic {
		loggerInstance.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return loggerInstance
}
