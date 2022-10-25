/*
Package helper ...
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

// TODO:
func (messagelogger *MessageLoggerDefault) Error(messageNumber int, details ...interface{}) error {
	errorMessage, err := messagelogger.Message(messageNumber, details...)
	if err != nil {
		return err
	}
	return errors.New(errorMessage)
}

// TODO:
func (messagelogger *MessageLoggerDefault) GetLogLevel() Level {
	return Level(messagelogger.Logger.GetLogLevel())
}

// TODO:
func (messagelogger *MessageLoggerDefault) GetLogLevelAsString() string {
	return (messagelogger.Logger.GetLogLevelAsString())
}

// TODO:
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

// TODO:
func (messagelogger *MessageLoggerDefault) Message(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil

	id := fmt.Sprintf("%d", messageNumber)
	if messagelogger.MessageId != nil {
		id, err = messagelogger.MessageId.MessageId(messageNumber, details...)
		if err != nil {
			return "", err
		}
	}

	text := ""
	if messagelogger.MessageText != nil {
		text, err = messagelogger.MessageText.MessageText(messageNumber, details...)
		if err != nil {
			return "", err
		}
	}

	status := ""
	if messagelogger.MessageStatus != nil {
		status, err = messagelogger.MessageStatus.MessageStatus(messageNumber, details...)
		if err != nil {
			return "", err
		}
	}

	result, err := messagelogger.MessageFormat.Message(id, status, text, details...)
	if err != nil {
		return "", err
	}

	return result, err
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetIdTemplate(idTemplate string) MessageLoggerInterface {
	messagelogger.MessageId.SetIdTemplate(idTemplate)
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetLogger(logger logger.LoggerInterface) MessageLoggerInterface {
	messagelogger.Logger = logger
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetLogLevel(level Level) MessageLoggerInterface {
	messagelogger.Logger.SetLogLevel(logger.Level(level))
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetLogLevelFromString(levelString string) MessageLoggerInterface {
	logger.SetLogLevelFromString(levelString)
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	messagelogger.MessageFormat = messageFormat
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetMessageId(messageId messageid.MessageIdInterface) MessageLoggerInterface {
	messagelogger.MessageId = messageId
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface {
	messagelogger.MessageLogLevel = messageLogLevel
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetMessageStatus(messageStatus messagestatus.MessageStatusInterface) MessageLoggerInterface {
	messagelogger.MessageStatus = messageStatus
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetMessageText(messageText messagetext.MessageTextInterface) MessageLoggerInterface {
	messagelogger.MessageText = messageText
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerDefault) SetTextTemplates(messages map[int]string) MessageLoggerInterface {
	messagelogger.MessageText.SetTextTemplates(messages)
	return messagelogger
}
