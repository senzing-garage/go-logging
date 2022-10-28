/*
The MessageLoggerDefault implementation aggregates id, status, text, and details to return a formatted string.
*/
package messagelogger

import (
	"errors"
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

// The MessageLoggerDefault type is for constructing and logging messages.
type MessageLoggerDefault struct {
	Logger          logger.LoggerInterface
	MessageFormat   messageformat.MessageFormatInterface
	MessageId       messageid.MessageIdInterface
	MessageLogLevel messageloglevel.MessageLogLevelInterface
	MessageStatus   messagestatus.MessageStatusInterface
	MessageText     messagetext.MessageTextInterface
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Write log record based on message level method.
func (messagelogger *MessageLoggerDefault) logBasedOnLevel(level Level, messageBody string) {
	switch level {
	case Level(logger.LevelInfo):
		messagelogger.Logger.Info(messageBody)
	case Level(logger.LevelWarn):
		messagelogger.Logger.Warn(messageBody)
	case Level(logger.LevelError):
		messagelogger.Logger.Error(messageBody)
	case Level(logger.LevelDebug):
		messagelogger.Logger.Debug(messageBody)
	case Level(logger.LevelTrace):
		messagelogger.Logger.Trace(messageBody)
	case Level(logger.LevelFatal):
		messagelogger.Logger.Fatal(messageBody)
	case Level(logger.LevelPanic):
		messagelogger.Logger.Panic(messageBody)
	default:
		messagelogger.Logger.Info(messageBody)
	}
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The Error method returns an error with the formatted message.
func (messagelogger *MessageLoggerDefault) Error(messageNumber int, details ...interface{}) error {
	errorMessage, err := messagelogger.Message(messageNumber, details...)
	if err != nil {
		return err
	}
	return errors.New(errorMessage)
}

// The GetLogLevel method returns the current log level as a typed int.
func (messagelogger *MessageLoggerDefault) GetLogLevel() Level {
	return Level(messagelogger.Logger.GetLogLevel())
}

// The GetLogLevelAsString method returns the current log level as a string.
func (messagelogger *MessageLoggerDefault) GetLogLevelAsString() string {
	return (messagelogger.Logger.GetLogLevelAsString())
}

// The Log method sends the formatted message to the Go log framework.
func (messagelogger *MessageLoggerDefault) Log(messageNumber int, details ...interface{}) error {
	var err error = nil

	messageBody, err := messagelogger.Message(messageNumber, details...)
	if err != nil {
		return err
	}

	messageLevel := logger.LevelInfo
	if messagelogger.MessageLogLevel != nil {
		messageLevel, err = messagelogger.MessageLogLevel.MessageLogLevel(messageNumber, details...)
		if err != nil {
			return err
		}
	}
	messagelogger.logBasedOnLevel(Level(messageLevel), messageBody)
	return err
}

// The Message method returns a string with the formatted message.
func (messagelogger *MessageLoggerDefault) Message(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil

	id := fmt.Sprintf("%d", messageNumber)
	if messagelogger.MessageId != nil {
		id, err = messagelogger.MessageId.MessageId(messageNumber, details...)
		if err != nil {
			id = fmt.Sprintf("%d", messageNumber)
		}
	}

	text := ""
	if messagelogger.MessageText != nil {
		text, _ = messagelogger.MessageText.MessageText(messageNumber, details...)
	}

	status := ""
	if messagelogger.MessageStatus != nil {
		status, _ = messagelogger.MessageStatus.MessageStatus(messageNumber, details...)
	}

	result, err := messagelogger.MessageFormat.Message(id, status, text, details...)
	if err != nil {
		return "", err
	}

	return result, err
}

// The SetLogLevel method sets the log level given a typed int.
func (messagelogger *MessageLoggerDefault) SetLogLevel(level Level) MessageLoggerInterface {
	messagelogger.Logger.SetLogLevel(logger.Level(level))
	return messagelogger
}

// The SetLogLevelFromString method sets the log level given a string.
// Acceptable string values: TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC.
func (messagelogger *MessageLoggerDefault) SetLogLevelFromString(levelString string) MessageLoggerInterface {
	messagelogger.Logger.SetLogLevelFromString(levelString)
	return messagelogger
}
