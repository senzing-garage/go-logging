/*
Package helper ...
*/
package messagelogger

import (
	"fmt"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messagelevel"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLoggerImpl struct {
	idTemplate    string
	messages      map[int]string
	messageformat messageformat.MessageFormatInterface
	messagelevel  messagelevel.MessageLevelInterface
	logger        logger.LoggerInterface
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messageLoggerInstance *MessageLoggerImpl

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

func New() *MessageLoggerImpl {
	result := new(MessageLoggerImpl)
	result.SetLevel(LevelWarn)
	result.SetMessageFormat(&messageformat.MessageFormatJson{})
	result.SetMessageLevel(&messagelevel.MessageLevelSenzingApi{})
	return result
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func init() {
	messageLoggerInstance = New()
}

// ----------------------------------------------------------------------------
// Public Setters and Getters
// ----------------------------------------------------------------------------

// --- IdTemplate -------------------------------------------------------------

func SetIdTemplate(idTemplate string) MessageLoggerInterface {
	return messageLoggerInstance.SetIdTemplate(idTemplate)
}
func (messagelogger *MessageLoggerImpl) SetIdTemplate(idTemplate string) MessageLoggerInterface {
	messagelogger.idTemplate = idTemplate
	return messagelogger
}

func GetIdTemplate() string { return messageLoggerInstance.GetIdTemplate() }
func (messagelogger *MessageLoggerImpl) GetIdTemplate() string {
	return messagelogger.idTemplate
}

// --- Level ------------------------------------------------------------------

func SetLevel(level Level) MessageLoggerInterface { return messageLoggerInstance.SetLevel(level) }
func (messagelogger *MessageLoggerImpl) SetLevel(level Level) MessageLoggerInterface {
	messagelogger.logger.SetLevel(logger.Level(level))
	return messagelogger
}

func GetLevel() Level { return messageLoggerInstance.GetLevel() }
func (messagelogger *MessageLoggerImpl) GetLevel() Level {
	return Level(messagelogger.logger.GetLevel())
}

// --- LevelFromString --------------------------------------------------------

func SetLevelFromString(levelString string) MessageLoggerInterface {
	return messageLoggerInstance.SetLevelFromString(levelString)
}
func (messagelogger *MessageLoggerImpl) SetLevelFromString(levelString string) MessageLoggerInterface {
	logger.SetLevelFromString(levelString)
	return messagelogger
}

func GetLevelAsString() string { return messageLoggerInstance.GetLevelAsString() }
func (messagelogger *MessageLoggerImpl) GetLevelAsString() string {
	return (messagelogger.logger.GetLevelAsString())
}

// --- Messages ---------------------------------------------------------------

func SetMessages(messages map[int]string) MessageLoggerInterface {
	return messageLoggerInstance.SetMessages(messages)
}
func (messagelogger *MessageLoggerImpl) SetMessages(messages map[int]string) MessageLoggerInterface {
	messagelogger.messages = messages
	return messagelogger
}

func GetMessages() map[int]string { return messageLoggerInstance.GetMessages() }
func (messagelogger *MessageLoggerImpl) GetMessages() map[int]string {
	return messagelogger.messages
}

// --- MessageFormat ----------------------------------------------------------

func SetMessageLevel(messageLevel messagelevel.MessageLevelInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageLevel(messageLevel)
}
func (messagelogger *MessageLoggerImpl) SetMessageLevel(messageLevel messagelevel.MessageLevelInterface) MessageLoggerInterface {
	messagelogger.messagelevel = messageLevel
	return messagelogger
}

// --- MessageFormat ----------------------------------------------------------

func SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageFormat(messageFormat)
}
func (messagelogger *MessageLoggerImpl) SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	messagelogger.messageformat = messageFormat
	return messagelogger
}

// ----------------------------------------------------------------------------
// Instance functions
// ----------------------------------------------------------------------------

// TODO:
func Log(errorNumber int, details ...interface{}) error {
	return messageLoggerInstance.Log(errorNumber, details...)
}

// Write log record based on message level function.
func LogBasedOnLevel(level Level, messageBody string) {
	messageLoggerInstance.LogBasedOnLevel(level, messageBody)
}

// Inspect the error to see what the level is and log based on the level function.
// func LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
// 	return messageLoggerInstance.LogMessage(idTemplate, errorNumber, message, details...)
// }

// Inspect the error to see what the level is and log based on the level function.
// func LogMessageFromError(idTemplate string, errorNumber int, message string, err error, details ...interface{}) error {
// 	return messageLoggerInstance.LogMessageFromError(idTemplate, errorNumber, message, err, details...)
// }

// Inspect the error to see what the level is and log based on the level function.
// func LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, err error, details map[string]interface{}) error {
// 	return messageLoggerInstance.LogMessageFromErrorUsingMap(idTemplate, errorNumber, message, err, details)
// }

// Inspect the error to see what the level is and log based on the level function.
// func LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
// 	return messageLoggerInstance.LogMessageUsingMap(idTemplate, errorNumber, message, details)
// }

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Write log record based on message level method.
func (messagelogger *MessageLoggerImpl) LogBasedOnLevel(level Level, messageBody string) {
	switch level {
	case Level(logger.LevelInfo):
		logger.Info(messageBody)
	case Level(logger.LevelWarn):
		logger.Warn(messageBody)
	case Level(logger.LevelError):
		logger.Error(messageBody)
	case Level(logger.LevelDebug):
		logger.Debug(messageBody)
	case Level(logger.LevelTrace):
		logger.Trace(messageBody)
	case Level(logger.LevelFatal):
		logger.Fatal(messageBody)
	case Level(logger.LevelPanic):
		logger.Panic(messageBody)
	default:
		logger.Info(messageBody)
	}
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLoggerImpl) Log(errorNumber int, details ...interface{}) error {
	var err error = nil

	status := ""

	text := ""
	textTemplate, ok := messagelogger.messages[errorNumber]
	if ok {
		text = fmt.Sprintf(textTemplate, details...)
	}

	messageLevel, err := messagelogger.messagelevel.CalculateMessageLevel(errorNumber, text)
	messageBody := messagelogger.messageformat.BuildMessage(messagelogger.idTemplate, status, text, details...)
	messagelogger.LogBasedOnLevel(Level(messageLevel), messageBody)
	return err
}

// Inspect the error to see what the level is and log based on the level method.
// func (messagelogger *MessageLoggerImpl) LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
// 	var err error = nil
// 	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
// 	messageJson := messagebuilder.BuildMessage(idTemplate, errorNumber, message, details...)
// 	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
// 	return err
// }

// Inspect the error to see what the level is and log based on the level method.
// func (messagelogger *MessageLoggerImpl) LogMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) error {
// 	var err error = nil
// 	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
// 	messageJson := messagebuilder.BuildMessageFromError(idTemplate, errorNumber, message, anError, details...)
// 	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
// 	return err
// }

// Inspect the error to see what the level is and log based on the level method.
// func (messagelogger *MessageLoggerImpl) LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) error {
// 	var err error = nil

// 	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
// 	messageJson := messagebuilder.BuildMessageFromErrorUsingMap(idTemplate, errorNumber, message, anError, details)
// 	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
// 	return err
// }

// Inspect the error to see what the level is and log based on the level method.
// func (messagelogger *MessageLoggerImpl) LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
// 	var err error = nil
// 	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
// 	messageJson := messagebuilder.BuildMessageUsingMap(idTemplate, errorNumber, message, details)
// 	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
// 	return err
// }
