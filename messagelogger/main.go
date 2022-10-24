/*
Package logger provides...
*/
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
	Error(messageNumber int, details ...interface{}) error
	GetLogLevel() Level
	GetLogLevelAsString() string
	Log(messageNumber int, details ...interface{}) error
	Message(messageNumber int, details ...interface{}) (string, error)
	SetIdTemplate(idTemplate string) MessageLoggerInterface
	SetLogLevel(level Level) MessageLoggerInterface
	SetLogLevelFromString(levelString string) MessageLoggerInterface
	SetTextTemplates(messages map[int]string) MessageLoggerInterface
}
