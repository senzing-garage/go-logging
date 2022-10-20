/*
Package helper ...
*/
package messagelogger

import (
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagebuilder"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLoggerImpl struct {
	idTemplate string
	level      Level
	messages   map[int]string
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messageLoggerInstance *MessageLoggerImpl

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

func New() *MessageLoggerImpl {
	return new(MessageLoggerImpl)
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
	logger.SetLevel(logger.Level(level))
	return messagelogger
}

func GetLevel() Level { return messageLoggerInstance.GetLevel() }
func (messagelogger *MessageLoggerImpl) GetLevel() Level {
	return messagelogger.level
}

// --- Messages ---------------------------------------------------------------

func SetMessages(messages map[int]string) MessageLoggerInterface {
	return messageLoggerInstance.SetMessages(messages)
}
func (messagelogger *MessageLoggerImpl) SetMessages(messages map[int]string) {
	messagelogger.messages = messages
	return messagelogger
}

func GetMessages() map[int]string { return messageLoggerInstance.GetMessages() }
func (messagelogger *MessageLoggerImpl) GetMessages() map[int]string {
	return messagelogger.messages
}

// ----------------------------------------------------------------------------
// Instance functions
// ----------------------------------------------------------------------------

// Write log record based on message level function.
func LogBasedOnLevel(messageLevel string, messageJson string) {
	messageLoggerInstance.LogBasedOnLevel(messageLevel, messageJson)
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	return messageLoggerInstance.LogMessage(idTemplate, errorNumber, message, details...)
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageFromError(idTemplate string, errorNumber int, message string, err error, details ...interface{}) error {
	return messageLoggerInstance.LogMessageFromError(idTemplate, errorNumber, message, err, details...)
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, err error, details map[string]interface{}) error {
	return messageLoggerInstance.LogMessageFromErrorUsingMap(idTemplate, errorNumber, message, err, details)
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
	return messageLoggerInstance.LogMessageUsingMap(idTemplate, errorNumber, message, details)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Write log record based on message level method.
func (messagelogger *MessageLoggerImpl) LogBasedOnLevel(messageLevel string, messageJson string) {
	switch messageLevel {
	case "info":
		logger.Info(messageJson)
	case "warning":
		logger.Warn(messageJson)
	case "error":
		logger.Error(messageJson)
	case "debug":
		logger.Debug(messageJson)
	case "trace":
		logger.Trace(messageJson)
	case "retryable":
		logger.Info(messageJson)
	case "reserved":
		logger.Info(messageJson)
	case "fatal":
		logger.Fatal(messageJson)
	case "panic":
		logger.Panic(messageJson)
	default:
		logger.Info(messageJson)
	}
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLoggerImpl) LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	var err error = nil
	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessage(idTemplate, errorNumber, message, details...)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLoggerImpl) LogMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) error {
	var err error = nil
	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessageFromError(idTemplate, errorNumber, message, anError, details...)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLoggerImpl) LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) error {
	var err error = nil

	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessageFromErrorUsingMap(idTemplate, errorNumber, message, anError, details)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLoggerImpl) LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
	var err error = nil
	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessageUsingMap(idTemplate, errorNumber, message, details)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}
