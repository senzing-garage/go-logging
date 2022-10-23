// The messagelogger package is a set of methods logging messages.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagelogger

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Level int

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Must match what's in logger/main.go

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageLoggerInterface interface {
	GetLogLevel() Level
	GetLogLevelAsString() string
	Log(errorNumber int, details ...interface{}) error
	Message(errorNumber int, details ...interface{}) (string, error)
	SetLogLevel(level Level) MessageLoggerInterface
	SetLogLevelFromString(levelString string) MessageLoggerInterface
	SetMessageIdTemplate(idTemplate string) MessageLoggerInterface
	SetMessages(messages map[int]string) MessageLoggerInterface
}
