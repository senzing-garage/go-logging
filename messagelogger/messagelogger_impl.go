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
	"github.com/senzing/go-logging/messagestatus"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLoggerImpl struct {
	IdTemplate      string
	Logger          logger.LoggerInterface
	MessageFormat   messageformat.MessageFormatInterface
	MessageLogLevel messageloglevel.MessageLogLevelInterface
	Messages        map[int]string
	MessageStatus   messagestatus.MessageStatusInterface
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
	result := &MessageLoggerImpl{
		IdTemplate:      DefaultIdTemplate,
		Logger:          &logger.LoggerImpl{},
		MessageFormat:   &messageformat.MessageFormatJson{},
		MessageLogLevel: &messageloglevel.MessageLogLevelSenzingApi{},
		MessageStatus:   &messagestatus.MessageStatusSenzingApi{},
	}
	result.SetLogLevel(LevelError)
	return result
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

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

func SetLogLevel(level Level) MessageLoggerInterface { return messageLoggerInstance.SetLogLevel(level) }
func (messagelogger *MessageLoggerImpl) SetLogLevel(level Level) MessageLoggerInterface {
	messagelogger.Logger.SetLogLevel(logger.Level(level))
	return messagelogger
}

func GetLogLevel() Level { return messageLoggerInstance.GetLogLevel() }
func (messagelogger *MessageLoggerImpl) GetLogLevel() Level {
	return Level(messagelogger.Logger.GetLogLevel())
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
	return (messagelogger.Logger.GetLogLevelAsString())
}

// --- MessageLogger ----------------------------------------------------------

func GetMessageLogger() *MessageLoggerImpl { return messageLoggerInstance }

// ----------------------------------------------------------------------------
// Instance functions
// ----------------------------------------------------------------------------

// TODO:
func Log(errorNumber int, details ...interface{}) error {
	return messageLoggerInstance.Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Inspect the error to see what the level is and log based on the level method.
func (messagelogger *MessageLoggerImpl) Log(errorNumber int, details ...interface{}) error {
	var err error = nil

	idTemplate := "%d"
	if messagelogger.IdTemplate != "" {
		idTemplate = messagelogger.IdTemplate
	}
	id := fmt.Sprintf(idTemplate, errorNumber)

	text := ""
	textTemplate, ok := messagelogger.Messages[errorNumber]
	if ok {
		textRaw := fmt.Sprintf(textTemplate, details...)
		text = strings.Split(textRaw, "%!(")[0]
	}

	status, err := messageLoggerInstance.MessageStatus.CalculateMessageStatus(errorNumber, text)
	if err != nil {
		return err
	}

	messageLevel, err := messagelogger.MessageLogLevel.CalculateMessageLogLevel(errorNumber, text)
	if err != nil {
		return err
	}
	messageBody := messagelogger.MessageFormat.BuildMessage(id, status, text, details...)
	messagelogger.logBasedOnLevel(Level(messageLevel), messageBody)
	return err
}
