// The messagelogger package is a set of methods logging messages.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagelogger

import (
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageloglevel"
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
	GetLogLevel() Level
	GetLogLevelAsString() string
	GetMessages() map[int]string
	Log(errorNumber int, details ...interface{}) error
	LogBasedOnLevel(level Level, messageBody string)
	SetIdTemplate(idTemplate string) MessageLoggerInterface
	SetLogger(logger logger.LoggerInterface) MessageLoggerInterface
	SetLogLevel(level Level) MessageLoggerInterface
	SetLogLevelFromString(levelString string) MessageLoggerInterface
	SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface
	SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface
	SetMessages(messages map[int]string) MessageLoggerInterface
}
