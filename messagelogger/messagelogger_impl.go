/*
Package helper ...
*/
package messagelogger

import (
	"fmt"
	"strings"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageloglevel"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLoggerImpl struct {
	IdTemplate      string
	Messages        map[int]string
	MessageFormat   messageformat.MessageFormatInterface
	MessageLogLevel messageloglevel.MessageLogLevelInterface
	logger          logger.LoggerInterface
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Must match what's in logger/main.go

const (
	DefaultIdTemplate = "senzing-9999%04d"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messageLoggerInstance *MessageLoggerImpl

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

func New() *MessageLoggerImpl {
	result := new(MessageLoggerImpl)
	result.SetLogger(&logger.LoggerImpl{})
	result.SetLogLevel(LevelError)
	result.SetIdTemplate(DefaultIdTemplate)
	result.SetMessageFormat(&messageformat.MessageFormatJson{})
	result.SetMessageLogLevel(&messageloglevel.MessageLogLevelSenzingApi{})
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
	messagelogger.IdTemplate = idTemplate
	return messagelogger
}

func GetIdTemplate() string { return messageLoggerInstance.GetIdTemplate() }
func (messagelogger *MessageLoggerImpl) GetIdTemplate() string {
	return messagelogger.IdTemplate
}

// --- LogLevel ---------------------------------------------------------------

func SetLogLevel(level Level) MessageLoggerInterface { return messageLoggerInstance.SetLogLevel(level) }
func (messagelogger *MessageLoggerImpl) SetLogLevel(level Level) MessageLoggerInterface {
	messagelogger.logger.SetLogLevel(logger.Level(level))
	return messagelogger
}

func GetLogLevel() Level { return messageLoggerInstance.GetLogLevel() }
func (messagelogger *MessageLoggerImpl) GetLogLevel() Level {
	return Level(messagelogger.logger.GetLogLevel())
}

// --- LogLevelFromString -----------------------------------------------------

func SetLogLevelFromString(levelString string) MessageLoggerInterface {
	return messageLoggerInstance.SetLogLevelFromString(levelString)
}
func (messagelogger *MessageLoggerImpl) SetLogLevelFromString(levelString string) MessageLoggerInterface {
	logger.SetLogLevelFromString(levelString)
	return messagelogger
}

func GetLogLevelAsString() string { return messageLoggerInstance.GetLogLevelAsString() }
func (messagelogger *MessageLoggerImpl) GetLogLevelAsString() string {
	return (messagelogger.logger.GetLogLevelAsString())
}

// --- Logger -----------------------------------------------------------------

func SetLogger(logger logger.LoggerInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetLogger(logger)
}
func (messagelogger *MessageLoggerImpl) SetLogger(logger logger.LoggerInterface) MessageLoggerInterface {
	messagelogger.logger = logger
	return messagelogger
}

// --- Messages ---------------------------------------------------------------

func SetMessages(messages map[int]string) MessageLoggerInterface {
	return messageLoggerInstance.SetMessages(messages)
}
func (messagelogger *MessageLoggerImpl) SetMessages(messages map[int]string) MessageLoggerInterface {
	messagelogger.Messages = messages
	return messagelogger
}

func GetMessages() map[int]string { return messageLoggerInstance.GetMessages() }
func (messagelogger *MessageLoggerImpl) GetMessages() map[int]string {
	return messagelogger.Messages
}

// --- MessageLevel -----------------------------------------------------------

func SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageLogLevel(messageLogLevel)
}
func (messagelogger *MessageLoggerImpl) SetMessageLogLevel(messageLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface {
	messagelogger.MessageLogLevel = messageLevel
	return messagelogger
}

// --- MessageFormat ----------------------------------------------------------

func SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	return messageLoggerInstance.SetMessageFormat(messageFormat)
}
func (messagelogger *MessageLoggerImpl) SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	messagelogger.MessageFormat = messageFormat
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
		messagelogger.logger.Info(messageBody)
	case Level(logger.LevelWarn):
		messagelogger.logger.Warn(messageBody)
	case Level(logger.LevelError):
		messagelogger.logger.Error(messageBody)
	case Level(logger.LevelDebug):
		messagelogger.logger.Debug(messageBody)
	case Level(logger.LevelTrace):
		messagelogger.logger.Trace(messageBody)
	case Level(logger.LevelFatal):
		messagelogger.logger.Fatal(messageBody)
	case Level(logger.LevelPanic):
		messagelogger.logger.Panic(messageBody)
	default:
		messagelogger.logger.Info(messageBody)
	}
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLoggerImpl) Log(errorNumber int, details ...interface{}) error {
	var err error = nil

	idTemplate := "%d"
	if messagelogger.IdTemplate != "" {
		idTemplate = messagelogger.IdTemplate
	}
	id := fmt.Sprintf(idTemplate, errorNumber)

	status := ""

	text := ""
	textTemplate, ok := messagelogger.Messages[errorNumber]

	if ok {
		textRaw := fmt.Sprintf(textTemplate, details...)
		text = strings.Split(textRaw, "%!(")[0]

	}

	messageLevel, err := messagelogger.MessageLogLevel.CalculateMessageLogLevel(errorNumber, text)
	messageBody := messagelogger.MessageFormat.BuildMessage(id, status, text, details...)
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
