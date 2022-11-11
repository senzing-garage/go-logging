/*
The MessageLoggerDefault implementation aggregates id, status, text, and details to return a formatted string.
*/
package messagelogger

import (
	"errors"
	"fmt"
	"time"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagedate"
	"github.com/senzing/go-logging/messageduration"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelocation"
	"github.com/senzing/go-logging/messageloglevel"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
	"github.com/senzing/go-logging/messagetime"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLoggerDefault type is for constructing and logging messages.
type MessageLoggerDefault struct {
	Logger          logger.LoggerInterface
	MessageDate     messagedate.MessageDateInterface
	MessageDuration messageduration.MessageDurationInterface
	MessageFormat   messageformat.MessageFormatInterface
	MessageId       messageid.MessageIdInterface
	MessageLocation messagelocation.MessageLocationInterface
	MessageLogLevel messageloglevel.MessageLogLevelInterface
	MessageStatus   messagestatus.MessageStatusInterface
	MessageText     messagetext.MessageTextInterface
	MessageTime     messagetime.MessageTimeInterface
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

// The IsDebug method return true if logging is DEBUG or lower.
func (messagelogger *MessageLoggerDefault) IsDebug() bool {
	return messagelogger.Logger.IsDebug()
}

// The IsError method return true if logging is ERROR or lower.
func (messagelogger *MessageLoggerDefault) IsError() bool {
	return messagelogger.Logger.IsError()
}

// The IsFatal method return true if logging is RATAL or lower.
func (messagelogger *MessageLoggerDefault) IsFatal() bool {
	return messagelogger.Logger.IsFatal()
}

// The IsInfo method return true if logging is INFO or lower.
func (messagelogger *MessageLoggerDefault) IsInfo() bool {
	return messagelogger.Logger.IsInfo()
}

// The IsPanic method return true if logging is PANIC or lower.
func (messagelogger *MessageLoggerDefault) IsPanic() bool {
	return messagelogger.Logger.IsPanic()
}

// The IsTrace method return true if logging is TRACE or lower.
func (messagelogger *MessageLoggerDefault) IsTrace() bool {
	return messagelogger.Logger.IsTrace()
}

// The IsWarn method return true if logging is WARN or lower.
func (messagelogger *MessageLoggerDefault) IsWarn() bool {
	return messagelogger.Logger.IsWarn()
}

// The Log method sends the formatted message to the Go log framework.
func (messagelogger *MessageLoggerDefault) Log(messageNumber int, details ...interface{}) error {

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
	var err error
	now := time.Now()

	date := ""
	if messagelogger.MessageDate != nil {
		date, _ = messagelogger.MessageDate.MessageDate(messageNumber, now, details...)
	}

	time := ""
	if messagelogger.MessageTime != nil {
		time, _ = messagelogger.MessageTime.MessageTime(messageNumber, now, details...)
	}

	id := fmt.Sprintf("%d", messageNumber)
	if messagelogger.MessageId != nil {
		id, err = messagelogger.MessageId.MessageId(messageNumber, details...)
		if err != nil {
			id = fmt.Sprintf("%d", messageNumber)
		}
	}

	location := ""
	if messagelogger.MessageLocation != nil {
		location, _ = messagelogger.MessageLocation.MessageLocation(messageNumber, details...)
	}

	level := ""
	if messagelogger.MessageLogLevel != nil {
		levelAsLevel, _ := messagelogger.MessageLogLevel.MessageLogLevel(messageNumber, details...)
		var ok bool
		level, ok = logger.LevelToTextMap[levelAsLevel]
		if !ok {
			level = ""
		}
	}

	status := ""
	if messagelogger.MessageStatus != nil {
		status, _ = messagelogger.MessageStatus.MessageStatus(messageNumber, details...)
	}

	text := ""
	if messagelogger.MessageText != nil {
		text, _ = messagelogger.MessageText.MessageText(messageNumber, details...)
	}

	duration := int64(0)
	if messagelogger.MessageDuration != nil {
		duration, _ = messagelogger.MessageDuration.MessageDuration(messageNumber, details...)
	}

	result, err := messagelogger.MessageFormat.Message(date, time, level, location, id, status, text, duration, details...)
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
