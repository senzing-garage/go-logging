// The logger package is a set of method to help with common tasks.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package logger

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Level int

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const noFormat = ""

// Order is important for the LevelXxxx variables

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

const (
	LevelDebugName = "DEBUG"
	LevelErrorName = "ERROR"
	LevelFatalName = "FATAL"
	LevelInfoName  = "INFO"
	LevelPanicName = "PANIC"
	LevelTraceName = "TRACE"
	LevelWarnName  = "WARN"
	// MessageIdFormat = "senzing-6511%04d"
)

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type LoggerInterface interface {
	Debug(v ...interface{}) LoggerInterface
	Debugf(format string, v ...interface{}) LoggerInterface
	Error(v ...interface{}) LoggerInterface
	Errorf(format string, v ...interface{}) LoggerInterface
	Fatal(v ...interface{}) LoggerInterface
	Fatalf(format string, v ...interface{}) LoggerInterface
	GetLevel() Level
	GetLevelAsString() string
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
	SetLevel(level Level) LoggerInterface
	SetLevelFromString(levelString string) LoggerInterface
	Trace(v ...interface{}) LoggerInterface
	Tracef(format string, v ...interface{}) LoggerInterface
	Warn(v ...interface{}) LoggerInterface
	Warnf(format string, v ...interface{}) LoggerInterface
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var levelToTextMap = map[Level]string{
	LevelTrace: LevelTraceName,
	LevelDebug: LevelDebugName,
	LevelInfo:  LevelInfoName,
	LevelWarn:  LevelWarnName,
	LevelError: LevelErrorName,
	LevelFatal: LevelFatalName,
	LevelPanic: LevelPanicName,
}

var textToLevelMap = map[string]Level{
	LevelTraceName: LevelTrace,
	LevelDebugName: LevelDebug,
	LevelInfoName:  LevelInfo,
	LevelWarnName:  LevelWarn,
	LevelErrorName: LevelError,
	LevelFatalName: LevelFatal,
	LevelPanicName: LevelPanic,
}
