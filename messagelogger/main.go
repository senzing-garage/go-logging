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

type MessageLoggerImpl struct {
	level Level
}

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
	GetLevel() Level
	LogBasedOnLevel(messageLevel string, messageJson string)
	LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error
	LogMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) error
	LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) error
	LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error
	SetLevel(level Level) MessageLoggerInterface
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messageLoggerInstance *MessageLoggerImpl
