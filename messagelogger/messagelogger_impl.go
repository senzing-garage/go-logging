/*
Package helper ...
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

type MessageLoggerImpl struct {
	Logger          logger.LoggerInterface
	MessageFormat   messageformat.MessageFormatInterface
	MessageId       messageid.MessageIdInterface
	MessageLogLevel messageloglevel.MessageLogLevelInterface
	Messages        map[int]string
	MessageStatus   messagestatus.MessageStatusInterface
	MessageText     messagetext.MessageTextInterface
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

// TODO:
func New() *MessageLoggerImpl {
	result := &MessageLoggerImpl{
		MessageId:       &messageid.MessageIdDefault{},
		Logger:          &logger.LoggerImpl{},
		MessageFormat:   &messageformat.MessageFormatJson{},
		MessageLogLevel: &messageloglevel.MessageLogLevelNull{},
		MessageStatus:   &messagestatus.MessageStatusNull{},
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
// Public Setters and Getters
// ----------------------------------------------------------------------------

// --- LogLevel ---------------------------------------------------------------

// TODO:
func SetLogLevel(level Level) MessageLoggerInterface { return messageLoggerInstance.SetLogLevel(level) }

// TODO:
func (messagelogger *MessageLoggerImpl) SetLogLevel(level Level) MessageLoggerInterface {
	messagelogger.Logger.SetLogLevel(logger.Level(level))
	return messagelogger
}

// TODO:
func GetLogLevel() Level { return messageLoggerInstance.GetLogLevel() }

// TODO:
func (messagelogger *MessageLoggerImpl) GetLogLevel() Level {
	return Level(messagelogger.Logger.GetLogLevel())
}

// --- LogLevelFromString -----------------------------------------------------

// TODO:
func SetLogLevelFromString(levelString string) MessageLoggerInterface {
	return messageLoggerInstance.SetLogLevelFromString(levelString)
}

// TODO:
func (messagelogger *MessageLoggerImpl) SetLogLevelFromString(levelString string) MessageLoggerInterface {
	logger.SetLogLevelFromString(levelString)
	return messagelogger
}

// TODO:
func GetLogLevelAsString() string { return messageLoggerInstance.GetLogLevelAsString() }

// TODO:
func (messagelogger *MessageLoggerImpl) GetLogLevelAsString() string {
	return (messagelogger.Logger.GetLogLevelAsString())
}

// --- MessageLogger ----------------------------------------------------------

// TODO:
func GetMessageLogger() *MessageLoggerImpl { return messageLoggerInstance }

// ----------------------------------------------------------------------------
// Instance functions
// ----------------------------------------------------------------------------

// TODO:
func Log(errorNumber int, details ...interface{}) error {
	return messageLoggerInstance.Log(errorNumber, details...)
}

// TODO:
func Message(errorNumber int, details ...interface{}) (string, error) {
	return messageLoggerInstance.Message(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelogger *MessageLoggerImpl) Message(errorNumber int, details ...interface{}) (string, error) {
	var err error = nil

	id, err := messagelogger.MessageId.MessageId(errorNumber)
	if err != nil {
		return "", err
	}

	text, err := messagelogger.MessageText.MessageText(errorNumber, details...)
	if err != nil {
		return "", err
	}

	status, err := messageLoggerInstance.MessageStatus.MessageStatus(errorNumber, text)
	if err != nil {
		return "", err
	}

	result := messagelogger.MessageFormat.BuildMessage(id, status, text, details...)
	return result, err
}

// TODO:
func (messagelogger *MessageLoggerImpl) Log(errorNumber int, details ...interface{}) error {
	var err error = nil

	messageBody, err := messagelogger.Message(errorNumber, details...)
	if err != nil {
		return err
	}
	messageLevel, err := messagelogger.MessageLogLevel.MessageLogLevel(errorNumber, messageBody)
	if err != nil {
		return err
	}
	messagelogger.logBasedOnLevel(Level(messageLevel), messageBody)
	return err
}
