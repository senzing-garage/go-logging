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

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messageLoggerInstance *MessageLoggerDefault

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

// TODO:
func New() *MessageLoggerDefault {
	result := &MessageLoggerDefault{
		Logger:          &logger.LoggerImpl{},
		MessageFormat:   &messageformat.MessageFormatJson{},
		MessageId:       &messageid.MessageIdDefault{},
		MessageLogLevel: &messageloglevel.MessageLogLevelDefault{},
		MessageStatus:   &messagestatus.MessageStatusNull{},
		MessageText:     &messagetext.MessageTextDefault{},
	}
	result.SetLogLevel(LevelInfo)
	return result
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

// TODO:
func init() {
	messageLoggerInstance = New()
}

// ----------------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------------

// TODO:
func Error(messageNumber int, details ...interface{}) error {
	return messageLoggerInstance.Error(messageNumber, details...)
}

// TODO:
func GetLogLevel() Level { return messageLoggerInstance.GetLogLevel() }

// TODO:
func GetLogLevelAsString() string { return messageLoggerInstance.GetLogLevelAsString() }

// TODO:
func Log(messageNumber int, details ...interface{}) error {
	return messageLoggerInstance.Log(messageNumber, details...)
}

// TODO:
func Message(messageNumber int, details ...interface{}) (string, error) {
	return messageLoggerInstance.Message(messageNumber, details...)
}

// TODO:
func GetMessageLogger() *MessageLoggerDefault { return messageLoggerInstance }

// TODO:
func SetIdTemplate(idTemplate string) MessageLoggerInterface {
	return messageLoggerInstance.SetIdTemplate(idTemplate)
}

// TODO:
func SetLogger(logger logger.LoggerInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetLogger(logger)
}

// TODO:
func SetLogLevel(level Level) MessageLoggerInterface { return messageLoggerInstance.SetLogLevel(level) }

// TODO:
func SetLogLevelFromString(levelString string) MessageLoggerInterface {
	return messageLoggerInstance.SetLogLevelFromString(levelString)
}

// TODO:
func SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageFormat(messageFormat)
}

// TODO:
func SetMessageId(messageId messageid.MessageIdInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageId(messageId)
}

// TODO:
func SetMessageLogger(messageLogger *MessageLoggerDefault) {
	messageLoggerInstance = messageLogger
}

// TODO:
func SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageLogLevel(messageLogLevel)
}

// TODO:
func SetMessageStatus(messageStatus messagestatus.MessageStatusInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageStatus(messageStatus)
}

// TODO:
func SetMessageText(messageText messagetext.MessageTextInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageText(messageText)
}

// TODO:
func SetTextTemplates(messages map[int]string) MessageLoggerInterface {
	return messageLoggerInstance.SetTextTemplates(messages)
}
