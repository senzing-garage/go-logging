/*
Package messagelogger generates and logs messages.
*/
package messagelogger

import (
	"fmt"

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

// For use with logging levels. (e.g. LevelInfo, LevelWarn, etc.)
type Level int

type MessageLoggerInterface interface {
	Error(messageNumber int, details ...interface{}) error
	GetLogLevel() Level
	GetLogLevelAsString() string
	Log(messageNumber int, details ...interface{}) error
	Message(messageNumber int, details ...interface{}) (string, error)
	// SetIdTemplate(idTemplate string) MessageLoggerInterface
	// SetLogger(logger logger.LoggerInterface) MessageLoggerInterface
	SetLogLevel(level Level) MessageLoggerInterface
	SetLogLevelFromString(levelString string) MessageLoggerInterface
	// SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface
	// SetMessageId(messageId messageid.MessageIdInterface) MessageLoggerInterface
	// SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface
	// SetMessageStatus(messageStatus messagestatus.MessageStatusInterface) MessageLoggerInterface
	// SetMessageText(messageText messagetext.MessageTextInterface) MessageLoggerInterface
	// SetTextTemplates(messages map[int]string) MessageLoggerInterface
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

// var messageLoggerInstance *MessageLoggerDefault

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

/*
The New method creates a new instance of MessageLoggerDefault.
The default message logger uses default and null subcomponents.
To use non-default subcomponents,
parameters to New() can specify the subcomponent desired.
The parameters can be of the following type and in any order:

	logger.LoggerInterface
	messageformat.MessageFormatInterface
	messageid.MessageIdInterface
	messageloglevel.MessageLogLevelInterface
	messagestatus.MessageStatusInterface
	messagetext.MessageTextInterface
	logger.Level

If a type is specified multiple times,
the last instance instance of the type specified wins.
*/
func New(interfaces ...interface{}) (MessageLoggerInterface, error) {
	var err error = nil
	logLevel := LevelInfo
	result := &MessageLoggerDefault{
		Logger:          &logger.LoggerDefault{},
		MessageFormat:   &messageformat.MessageFormatDefault{},
		MessageId:       &messageid.MessageIdDefault{},
		MessageLogLevel: &messageloglevel.MessageLogLevelDefault{},
		MessageStatus:   &messagestatus.MessageStatusNull{},
		MessageText:     &messagetext.MessageTextNull{},
	}

	var errorsList []interface{}
	if len(interfaces) > 0 {
		for _, value := range interfaces {
			switch typedValue := value.(type) {
			case logger.LoggerInterface:
				result.Logger = typedValue
			case messageformat.MessageFormatInterface:
				result.MessageFormat = typedValue
			case messageid.MessageIdInterface:
				result.MessageId = typedValue
			case messageloglevel.MessageLogLevelInterface:
				result.MessageLogLevel = typedValue
			case messagestatus.MessageStatusInterface:
				result.MessageStatus = typedValue
			case messagetext.MessageTextInterface:
				result.MessageText = typedValue
			case logger.Level:
				logLevelCandidate, ok := value.(logger.Level)
				if ok {
					logLevel = Level(logLevelCandidate)
				}
			default:
				errorsList = append(errorsList, typedValue)
			}
		}
	}

	if len(errorsList) > 0 {
		err = fmt.Errorf("unsupported interfaces: %#v", errorsList)
	}

	result.SetLogLevel(logLevel)
	return result, err
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

// func init() {
// 	messageLoggerInstance = New()
// }

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// func Error(messageNumber int, details ...interface{}) error {
// 	return messageLoggerInstance.Error(messageNumber, details...)
// }

// func GetLogLevel() Level {
// 	return messageLoggerInstance.GetLogLevel()
// }

// func GetLogLevelAsString() string {
// 	return messageLoggerInstance.GetLogLevelAsString()
// }

// func Log(messageNumber int, details ...interface{}) error {
// 	return messageLoggerInstance.Log(messageNumber, details...)
// }

// func Message(messageNumber int, details ...interface{}) (string, error) {
// 	return messageLoggerInstance.Message(messageNumber, details...)
// }

// func MsgNumber(messageNumber int) messagetext.MessageNumber {
// 	return messagetext.MsgNumber(messageNumber)
// }

// func GetMessageLogger() *MessageLoggerDefault {
// 	return messageLoggerInstance
// }

// func SetIdTemplate(idTemplate string) MessageLoggerInterface {
// 	return messageLoggerInstance.SetIdTemplate(idTemplate)
// }

// func SetLogger(logger logger.LoggerInterface) MessageLoggerInterface {
// 	return messageLoggerInstance.SetLogger(logger)
// }

// func SetLogLevel(level Level) MessageLoggerInterface {
// 	return messageLoggerInstance.SetLogLevel(level)
// }

// func SetLogLevelFromString(levelString string) MessageLoggerInterface {
// 	return messageLoggerInstance.SetLogLevelFromString(levelString)
// }

// func SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
// 	return messageLoggerInstance.SetMessageFormat(messageFormat)
// }

// func SetMessageId(messageId messageid.MessageIdInterface) MessageLoggerInterface {
// 	return messageLoggerInstance.SetMessageId(messageId)
// }

// func SetMessageLogger(messageLogger *MessageLoggerDefault) {
// 	messageLoggerInstance = messageLogger
// }

// func SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface {
// 	return messageLoggerInstance.SetMessageLogLevel(messageLogLevel)
// }

// func SetMessageStatus(messageStatus messagestatus.MessageStatusInterface) MessageLoggerInterface {
// 	return messageLoggerInstance.SetMessageStatus(messageStatus)
// }

// func SetMessageText(messageText messagetext.MessageTextInterface) MessageLoggerInterface {
// 	return messageLoggerInstance.SetMessageText(messageText)
// }

// func SetTextTemplates(messages map[int]string) MessageLoggerInterface {
// 	return messageLoggerInstance.SetTextTemplates(messages)
// }
