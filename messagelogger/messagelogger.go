/*
Package helper ...
*/
package messagelogger

import (
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagebuilder"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func init() {
	messagelogger = New()
}

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

func SetLevel(level Level) *MessageLogger { return messagelogger.SetLevel(level) }
func (messagelogger *MessageLogger) SetLevel(level Level) *MessageLogger {
	logger.SetLevel(logger.Level(level))
	return messagelogger
}

func New() *MessageLogger {
	return new(MessageLogger)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Write log record based on message level function.
func LogBasedOnLevel(messageLevel string, messageJson string) {
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
}

// Write log record based on message level method.
func (messagelogger *MessageLogger) LogBasedOnLevel(messageLevel string, messageJson string) {
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

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Inspect the error to see what the level is and log based on the level function.
func LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	return messagelogger.LogMessage(idTemplate, errorNumber, message, details...)
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLogger) LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	var err error = nil
	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessage(idTemplate, errorNumber, message, details...)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageFromError(idTemplate string, errorNumber int, message string, err error, details ...interface{}) error {
	return messagelogger.LogMessageFromError(idTemplate, errorNumber, message, err, details...)
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLogger) LogMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) error {
	var err error = nil

	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessageFromError(idTemplate, errorNumber, message, anError, details...)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, err error, details map[string]interface{}) error {
	return messagelogger.LogMessageFromErrorUsingMap(idTemplate, errorNumber, message, err, details)
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLogger) LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) error {
	var err error = nil

	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessageFromErrorUsingMap(idTemplate, errorNumber, message, anError, details)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
	return messagelogger.LogMessageUsingMap(idTemplate, errorNumber, message, details)
}

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLogger) LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
	var err error = nil
	messageLevel := messagebuilder.BuildMessageLevel(errorNumber, message)
	messageJson := messagebuilder.BuildMessageUsingMap(idTemplate, errorNumber, message, details)
	messagelogger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}
