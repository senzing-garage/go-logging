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

type MessageLoggerImpl struct {
	Logger          logger.LoggerInterface
	MessageFormat   messageformat.MessageFormatInterface
	MessageId       messageid.MessageIdInterface
	MessageLogLevel messageloglevel.MessageLogLevelInterface
	MessageStatus   messagestatus.MessageStatusInterface
	MessageText     messagetext.MessageTextInterface
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messageLoggerInstance *MessageLoggerImpl

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

// TODO:
func New() *MessageLoggerImpl {
	result := &MessageLoggerImpl{
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
// Internal methods
// ----------------------------------------------------------------------------

// Write log record based on message level method.
func (messagelogger *MessageLoggerImpl) logBasedOnLevel(level Level, messageBody string) {
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
// Instance functions
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
func GetMessageLogger() *MessageLoggerImpl { return messageLoggerInstance }

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
func SetMessageLogger(messageLogger *MessageLoggerImpl) {
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

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelogger *MessageLoggerImpl) Error(messageNumber int, details ...interface{}) error {
	errorMessage, err := messagelogger.Message(messageNumber, details...)
	if err != nil {
		return err
	}
	return errors.New(errorMessage)
}

// TODO:
func (messagelogger *MessageLoggerImpl) GetLogLevel() Level {
	return Level(messagelogger.Logger.GetLogLevel())
}

// TODO:
func (messagelogger *MessageLoggerImpl) GetLogLevelAsString() string {
	return (messagelogger.Logger.GetLogLevelAsString())
}

// TODO:
func (messagelogger *MessageLoggerImpl) Log(messageNumber int, details ...interface{}) error {
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
func (messagelogger *MessageLoggerImpl) Message(messageNumber int, details ...interface{}) (string, error) {
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
func (messagelogger *MessageLoggerImpl) SetIdTemplate(idTemplate string) MessageLoggerInterface {
	messagelogger.MessageId.SetIdTemplate(idTemplate)
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetLogger(logger logger.LoggerInterface) MessageLoggerInterface {
	messagelogger.Logger = logger
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetLogLevel(level Level) MessageLoggerInterface {
	messagelogger.Logger.SetLogLevel(logger.Level(level))
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetLogLevelFromString(levelString string) MessageLoggerInterface {
	logger.SetLogLevelFromString(levelString)
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetMessageFormat(messageFormat messageformat.MessageFormatInterface) MessageLoggerInterface {
	messagelogger.MessageFormat = messageFormat
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetMessageId(messageId messageid.MessageIdInterface) MessageLoggerInterface {
	messagelogger.MessageId = messageId
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetMessageLogLevel(messageLogLevel messageloglevel.MessageLogLevelInterface) MessageLoggerInterface {
	messagelogger.MessageLogLevel = messageLogLevel
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetMessageStatus(messageStatus messagestatus.MessageStatusInterface) MessageLoggerInterface {
	messagelogger.MessageStatus = messageStatus
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetMessageText(messageText messagetext.MessageTextInterface) MessageLoggerInterface {
	messagelogger.MessageText = messageText
	return messagelogger
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetTextTemplates(messages map[int]string) MessageLoggerInterface {
	messagelogger.MessageText.SetTextTemplates(messages)
	return messagelogger
}
