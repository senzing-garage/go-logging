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

type Logger struct {
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
	LevelDebugName  = "DEBUG"
	LevelErrorName  = "ERROR"
	LevelFatalName  = "FATAL"
	LevelInfoName   = "INFO"
	LevelPanicName  = "PANIC"
	LevelTraceName  = "TRACE"
	LevelWarnName   = "WARN"
	MessageIdFormat = "senzing-6511%04d"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var logger *Logger
