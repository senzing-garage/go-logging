/*
Package logger provides...
*/
package messagelogger

import (
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messageloglevel"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
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
	Error(messageNumber int, details ...interface{}) error
	GetLogLevel() Level
	GetLogLevelAsString() string
	Log(messageNumber int, details ...interface{}) error
	Message(messageNumber int, details ...interface{}) (string, error)
	SetIdTemplate(idTemplate string) MessageLoggerInterface
	SetLogger(logger logger.LoggerInterface) MessageLoggerInterface
	SetLogLevel(level Level) MessageLoggerInterface
	SetLogLevelFromString(levelString string) MessageLoggerInterface
	SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface
	SetMessageId(messageId messageid.MessageIdInterface) MessageLoggerInterface
	SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface
	SetMessageStatus(messageStatus messagestatus.MessageStatusInterface) MessageLoggerInterface
	SetMessageText(messageText messagetext.MessageTextInterface) MessageLoggerInterface
	SetTextTemplates(messages map[int]string) MessageLoggerInterface
}
