/*
The logger package is a decorator over Go's log package.
*/
package logger

import "log"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The Level type is used in LevelXxxxx constants.
type Level int

// The LoggerInterface type defines guards, logging methods and get/set of logging level.
type LoggerInterface interface {
	Debug(v ...interface{}) LoggerInterface
	Debugf(format string, v ...interface{}) LoggerInterface
	Error(v ...interface{}) LoggerInterface
	Errorf(format string, v ...interface{}) LoggerInterface
	Fatal(v ...interface{}) LoggerInterface
	Fatalf(format string, v ...interface{}) LoggerInterface
	GetLogLevel() Level
	GetLogLevelAsString() string
	Info(v ...interface{}) LoggerInterface
	Infof(format string, v ...interface{}) LoggerInterface
	IsDebug() bool
	IsError() bool
	IsFatal() bool
	IsInfo() bool
	IsPanic() bool
	IsTrace() bool
	IsWarn() bool
	Panic(v ...interface{}) LoggerInterface
	Panicf(format string, v ...interface{}) LoggerInterface
	SetLogLevel(level Level) LoggerInterface
	SetLogLevelFromString(levelString string) LoggerInterface
	Trace(v ...interface{}) LoggerInterface
	Tracef(format string, v ...interface{}) LoggerInterface
	Warn(v ...interface{}) LoggerInterface
	Warnf(format string, v ...interface{}) LoggerInterface
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const noFormat = ""

/*
LevelXxxx values are an enumeration of typed integers representing logging levels.
Order is important for the LevelXxxx variables.
*/
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// Strings printed when logging.
const (
	LevelDebugName = "DEBUG"
	LevelErrorName = "ERROR"
	LevelFatalName = "FATAL"
	LevelInfoName  = "INFO"
	LevelPanicName = "PANIC"
	LevelTraceName = "TRACE"
	LevelWarnName  = "WARN"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Map from Log level as typed integer to string.
var levelToTextMap = map[Level]string{
	LevelTrace: LevelTraceName,
	LevelDebug: LevelDebugName,
	LevelInfo:  LevelInfoName,
	LevelWarn:  LevelWarnName,
	LevelError: LevelErrorName,
	LevelFatal: LevelFatalName,
	LevelPanic: LevelPanicName,
}

// Default logger instance.
var loggerInstance *LoggerDefault

// Map from Log level as string to typed integer.
var textToLevelMap = map[string]Level{
	LevelTraceName: LevelTrace,
	LevelDebugName: LevelDebug,
	LevelInfoName:  LevelInfo,
	LevelWarnName:  LevelWarn,
	LevelErrorName: LevelError,
	LevelFatalName: LevelFatal,
	LevelPanicName: LevelPanic,
}

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

// Create a new default instance of the logger.
func New() *LoggerDefault {
	result := &LoggerDefault{}
	result.SetLogLevel(LevelInfo)
	return result
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

// Initialize default instance.
func init() {
	loggerInstance = New()
}

// ----------------------------------------------------------------------------
// Public functions for default logger instance.
// ----------------------------------------------------------------------------

func Debug(v ...interface{}) LoggerInterface {
	if loggerInstance.IsDebug() {
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

func Error(v ...interface{}) LoggerInterface {
	if loggerInstance.IsError() {
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

func Fatal(v ...interface{}) LoggerInterface {
	if loggerInstance.IsFatal() {
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

func Info(v ...interface{}) LoggerInterface {
	if loggerInstance.IsInfo() {
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

func SetLogLevel(level Level) LoggerInterface {
	return loggerInstance.SetLogLevel(level)
}

func SetLogLevelFromString(levelString string) LoggerInterface {
	return loggerInstance.SetLogLevelFromString(levelString)
}

func GetLogLevel() Level {
	return loggerInstance.GetLogLevel()
}

func GetLogLevelAsString() string {
	return loggerInstance.GetLogLevelAsString()
}

func IsDebug() bool {
	return loggerInstance.IsDebug()
}

func IsError() bool {
	return loggerInstance.IsError()
}

func IsFatal() bool {
	return loggerInstance.IsFatal()
}

func IsInfo() bool {
	return loggerInstance.IsInfo()
}

func IsPanic() bool {
	return loggerInstance.IsPanic()
}

func IsTrace() bool {
	return loggerInstance.IsTrace()
}

func IsWarn() bool {
	return loggerInstance.IsWarn()
}

func Panic(v ...interface{}) LoggerInterface {
	if loggerInstance.IsPanic() {
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

func Trace(v ...interface{}) LoggerInterface {
	if loggerInstance.IsTrace() {
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

func Warn(v ...interface{}) LoggerInterface {
	if loggerInstance.IsWarn() {
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
