// The messagelogger package is a set of methods logging messages.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagelogger

import (
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messagelevel"
)

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
	GetIdTemplate() string
	GetLevel() Level
	GetLevelAsString() string
	GetMessages() map[int]string
	Log(errorNumber int, details ...interface{}) error
	LogBasedOnLevel(level Level, messageBody string)
	SetIdTemplate(idTemplate string) MessageLoggerInterface
	SetLevel(level Level) MessageLoggerInterface
	SetLevelFromString(levelString string) MessageLoggerInterface
	SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface
	SetMessageLevel(messageLevel messagelevel.MessageLevelInterface) MessageLoggerInterface
	SetMessages(messages map[int]string) MessageLoggerInterface
}
