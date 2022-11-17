/*
The logger package is a decorator over Go's log package.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/logger/logger_test.go
*/
package logger

import "log"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The Level type is used to identify the integer is the detail parameters
// and is used in LevelXxxxx constants.
type Level int

// The LoggerInterface type defines guards, logging methods, and get/set of logging level.
type LoggerInterface interface {
	Debug(v ...interface{}) LoggerInterface                   // Log a DEBUG message.
	Debugf(format string, v ...interface{}) LoggerInterface   // Log a formatted DEBUG message.
	Error(v ...interface{}) LoggerInterface                   // Log an ERROR message.
	Errorf(format string, v ...interface{}) LoggerInterface   // Log a formatted ERROR message.
	Fatal(v ...interface{}) LoggerInterface                   // Log a FATAL message.
	Fatalf(format string, v ...interface{}) LoggerInterface   // Log a formatted FATAL message.
	GetLogLevel() Level                                       // Gets the logger instance logging level.
	GetLogLevelAsString() string                              // Gets the logger instance logging level in string representation.
	Info(v ...interface{}) LoggerInterface                    // Log an INFO message.
	Infof(format string, v ...interface{}) LoggerInterface    // Log a formatted INFO message.
	IsDebug() bool                                            // Returns true if a DEBUG message will be logged.
	IsError() bool                                            // Returns true if an ERROR message will be logged.
	IsFatal() bool                                            // Returns true if a FATAL message will be logged.
	IsInfo() bool                                             // Returns true if an INFO message will be logged.
	IsPanic() bool                                            // Returns true if a PANIC message will be logged.
	IsTrace() bool                                            // Returns true if a TRACE message will be logged.
	IsWarn() bool                                             // Returns true if a WARN message will be logged.
	Panic(v ...interface{}) LoggerInterface                   // Log a PANIC message.
	Panicf(format string, v ...interface{}) LoggerInterface   // Log a formatted PANIC message.
	SetLogLevel(level Level) LoggerInterface                  // Sets the logger instance logging level.
	SetLogLevelFromString(levelString string) LoggerInterface // Sets the logger instance logging level using a string representation.
	Trace(v ...interface{}) LoggerInterface                   // Log a TRACE message.
	Tracef(format string, v ...interface{}) LoggerInterface   // Log a formatted TRACE message.
	Warn(v ...interface{}) LoggerInterface                    // Log a WARN message.
	Warnf(format string, v ...interface{}) LoggerInterface    // Log a formatted WARN message.
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

// Strings representing the supported logging levels.
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

// Map from Log level as typed integer to string representation.
var LevelToTextMap = map[Level]string{
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

// Map from string representation to Log level as typed integer.
var TextToLevelMap = map[string]Level{
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

// Debug() logs a DEBUG message.
func Debug(v ...interface{}) LoggerInterface {
	if loggerInstance.IsDebug() {
		loggerInstance.printf(LevelDebugName, noFormat, v...)
	}
	return loggerInstance
}

// Debugf() logs a formatted DEBUG message.
func Debugf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsDebug() {
		loggerInstance.printf(LevelDebugName, format, v...)
	}
	return loggerInstance
}

// Error() logs a ERROR message.
func Error(v ...interface{}) LoggerInterface {
	if loggerInstance.IsError() {
		loggerInstance.printf(LevelErrorName, noFormat, v...)
	}
	return loggerInstance
}

// Errorf() logs a formatted ERROR message.
func Errorf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsError() {
		loggerInstance.printf(LevelErrorName, format, v...)
	}
	return loggerInstance
}

// Fatal() logs a FATAL message.
func Fatal(v ...interface{}) LoggerInterface {
	if loggerInstance.IsFatal() {
		loggerInstance.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return loggerInstance
}

// Fatalf() logs a formatted FATAL message.
func Fatalf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsFatal() {
		loggerInstance.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return loggerInstance
}

// GetLogLevel() gets the logger instance logging level.
func GetLogLevel() Level {
	return loggerInstance.GetLogLevel()
}

// GetLogLevelAsString() gets the logger instance logging level in string representation.
func GetLogLevelAsString() string {
	return loggerInstance.GetLogLevelAsString()
}

// Info() logs a INFO message.
func Info(v ...interface{}) LoggerInterface {
	if loggerInstance.IsInfo() {
		loggerInstance.printf(LevelInfoName, noFormat, v...)
	}
	return loggerInstance
}

// Infof() logs a formatted INFO message.
func Infof(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsInfo() {
		loggerInstance.printf(LevelInfoName, format, v...)
	}
	return loggerInstance
}

// IsDebug() returns true if the logger instance will log a DEBUG message.
func IsDebug() bool {
	return loggerInstance.IsDebug()
}

// IsError() returns true if the logger instance will log a ERROR message.
func IsError() bool {
	return loggerInstance.IsError()
}

// IsFatal() returns true if the logger instance will log a FATAL message.
func IsFatal() bool {
	return loggerInstance.IsFatal()
}

// IsInfo() returns true if the logger instance will log a INFO message.
func IsInfo() bool {
	return loggerInstance.IsInfo()
}

// IsPanic() returns true if the logger instance will log a PANIC message.
func IsPanic() bool {
	return loggerInstance.IsPanic()
}

// IsTrace() returns true if the logger instance will log a TRACE message.
func IsTrace() bool {
	return loggerInstance.IsTrace()
}

// IsWarn() returns true if the logger instance will log a WARN message.
func IsWarn() bool {
	return loggerInstance.IsWarn()
}

// Panic() logs a PANIC message.
func Panic(v ...interface{}) LoggerInterface {
	if loggerInstance.IsPanic() {
		loggerInstance.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return loggerInstance
}

// Panicf() logs a formatted PANIC message.
func Panicf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsPanic() {
		loggerInstance.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return loggerInstance
}

// SetLogLevel() sets the logger instance logging level.
func SetLogLevel(level Level) LoggerInterface {
	return loggerInstance.SetLogLevel(level)
}

// SetLogLevelFromString() sets the logger instance logging level using a string representation.
func SetLogLevelFromString(levelString string) LoggerInterface {
	return loggerInstance.SetLogLevelFromString(levelString)
}

// Trace() logs a TRACE message.
func Trace(v ...interface{}) LoggerInterface {
	if loggerInstance.IsTrace() {
		loggerInstance.printf(LevelTraceName, noFormat, v...)
	}
	return loggerInstance
}

// Tracef() logs a formatted TRACE message.
func Tracef(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsTrace() {
		loggerInstance.printf(LevelTraceName, format, v...)
	}
	return loggerInstance
}

// Warn() logs a WARN message.
func Warn(v ...interface{}) LoggerInterface {
	if loggerInstance.IsWarn() {
		loggerInstance.printf(LevelWarnName, noFormat, v...)
	}
	return loggerInstance
}

// Warnf() logs a formatted WARN message.
func Warnf(format string, v ...interface{}) LoggerInterface {
	if loggerInstance.IsWarn() {
		loggerInstance.printf(LevelWarnName, format, v...)
	}
	return loggerInstance
}
