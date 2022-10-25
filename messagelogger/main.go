/*
Package messagelogger generates and logs messages.
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
// Variables
// ----------------------------------------------------------------------------

var messageLoggerInstance *MessageLoggerDefault

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

func New() *MessageLoggerDefault {
	result := &MessageLoggerDefault{
		Logger:          &logger.LoggerDefault{},
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

func init() {
	messageLoggerInstance = New()
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

func Error(messageNumber int, details ...interface{}) error {
	return messageLoggerInstance.Error(messageNumber, details...)
}

func GetLogLevel() Level {
	return messageLoggerInstance.GetLogLevel()
}

func GetLogLevelAsString() string {
	return messageLoggerInstance.GetLogLevelAsString()
}

func Log(messageNumber int, details ...interface{}) error {
	return messageLoggerInstance.Log(messageNumber, details...)
}

func Message(messageNumber int, details ...interface{}) (string, error) {
	return messageLoggerInstance.Message(messageNumber, details...)
}

func MsgNumber(messageNumber int) messagetext.MessageNumber {
	return messagetext.MsgNumber(messageNumber)
}

func GetMessageLogger() *MessageLoggerDefault {
	return messageLoggerInstance
}

func SetIdTemplate(idTemplate string) MessageLoggerInterface {
	return messageLoggerInstance.SetIdTemplate(idTemplate)
}

func SetLogger(logger logger.LoggerInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetLogger(logger)
}

func SetLogLevel(level Level) MessageLoggerInterface {
	return messageLoggerInstance.SetLogLevel(level)
}

func SetLogLevelFromString(levelString string) MessageLoggerInterface {
	return messageLoggerInstance.SetLogLevelFromString(levelString)
}

func SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageFormat(messageFormat)
}

func SetMessageId(messageId messageid.MessageIdInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageId(messageId)
}

func SetMessageLogger(messageLogger *MessageLoggerDefault) {
	messageLoggerInstance = messageLogger
}

func SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageLogLevel(messageLogLevel)
}

func SetMessageStatus(messageStatus messagestatus.MessageStatusInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageStatus(messageStatus)
}

func SetMessageText(messageText messagetext.MessageTextInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageText(messageText)
}

func SetTextTemplates(messages map[int]string) MessageLoggerInterface {
	return messageLoggerInstance.SetTextTemplates(messages)
}
