package logging

import (
	"context"
	"errors"
	"time"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-messaging/messenger"
	"golang.org/x/exp/slog"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicLogging is an type-struct for an implementation of the loggingInterface.
type BasicLogging struct {
	// Using Ctx is not a preferred practice, but used to simplify Log() calls.
	Ctx          context.Context //nolint
	messenger    messenger.Messenger
	logger       *slog.Logger
	leveler      *slog.LevelVar
	logLevelName string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Error method returns an error with a JSON message based on the messageNumber and details.

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.

Output
  - error
*/
func (loggingImpl *BasicLogging) NewError(messageNumber int, details ...interface{}) error {
	transformedDetails := transformDetails(details...)
	return errors.New(loggingImpl.messenger.NewJSON(messageNumber, transformedDetails...)) //nolint
}

/*
The GetLogLevel method retrieves the current log level name.

Output
  - One of the following string values: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", or "PANIC"
*/
func (loggingImpl *BasicLogging) GetLogLevel() string {
	return loggingImpl.logLevelName
}

/*
The Is method is used to determine if a log message will be printed.

Output
  - True, if message would be logged at the logLevelName level.
*/
func (loggingImpl *BasicLogging) Is(logLevelName string) bool {
	result := false
	logLevel, ok := TextToLevelMap[logLevelName]

	if ok {
		result = loggingImpl.logger.Enabled(loggingImpl.Ctx, logLevel)
	}

	return result
}

/*
The IsDebug method is used to determine if DEBUG messages will be logged.

Output
  - If true, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC messages will be logged.
*/
func (loggingImpl *BasicLogging) IsDebug() bool {
	return loggingImpl.Is(LevelDebugName)
}

/*
The IsError method is used to determine if ERROR messages will be logged.

Output
  - If true, ERROR, FATAL, and PANIC messages will be logged.
*/
func (loggingImpl *BasicLogging) IsError() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelErrorName])
}

/*
The IsFatal method is used to determine if FATAL messages will be logged.

Output
  - If true, FATAL and PANIC messages will be logged.
*/
func (loggingImpl *BasicLogging) IsFatal() bool {
	return loggingImpl.Is(LevelFatalName)
}

/*
The IsInfo method is used to determine if INFO messages will be logged.

Output
  - If true, INFO, WARN, ERROR, FATAL, and PANIC messages will be logged.
*/
func (loggingImpl *BasicLogging) IsInfo() bool {
	return loggingImpl.Is(LevelInfoName)
}

/*
The IsPanic method is used to determine if PANIC messages will be logged.

Output
  - If true, PANIC messages will be logged.
*/
func (loggingImpl *BasicLogging) IsPanic() bool {
	return loggingImpl.Is(LevelPanicName)
}

/*
The IsTrace method is used to determine if TRACE messages will be logged.

Output
  - If true, TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC messages will be logged.
*/
func (loggingImpl *BasicLogging) IsTrace() bool {
	return loggingImpl.Is(LevelTraceName)
}

/*
The IsWarn method is used to determine if WARN messages will be logged.

Output
  - If true, WARN, ERROR, FATAL, and PANIC messages will be logged.
*/
func (loggingImpl *BasicLogging) IsWarn() bool {
	return loggingImpl.Is(LevelWarnName)
}

/*
The Json method returns a JSON string based on the messageNumber and details.

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.

Output
  - JSON string with message key/value pairs.
*/
func (loggingImpl *BasicLogging) JSON(messageNumber int, details ...interface{}) string {
	transformedDetails := transformDetails(details...)

	return loggingImpl.messenger.NewJSON(messageNumber, transformedDetails...)
}

/*
The Log method writes a log record to the output specified at LoggingImpl creation

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.
*/
func (loggingImpl *BasicLogging) Log(messageNumber int, details ...interface{}) {
	transformedDetails := transformDetails(details...)
	message, logLevel, newDetails := loggingImpl.messenger.NewSlogLevel(
		messageNumber,
		transformedDetails...,
	)
	newTransformedDetails := transformDetails(newDetails...)
	loggingImpl.logger.Log(loggingImpl.Ctx, logLevel, message, newTransformedDetails...)
}

/*
The SetLogLevel method changes the level of log messages generated.

Input
  - logLevelName: One of these strings:  "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", or "PANIC".

Output
  - error
*/
func (loggingImpl *BasicLogging) SetLogLevel(logLevelName string) error {
	var err error

	slogLevel, ok := TextToLevelMap[logLevelName]
	if !ok {
		return wraperror.Errorf(
			errLogging,
			"unknown error level: %s error: %w",
			logLevelName,
			errLogging,
		)
	}

	loggingImpl.leveler.Set(slogLevel)
	loggingImpl.logLevelName = logLevelName

	return err
}

// ----------------------------------------------------------------------------
// Private methods
// ----------------------------------------------------------------------------

func (loggingImpl *BasicLogging) initialize() {
	if loggingImpl.Ctx == nil {
		loggingImpl.Ctx = context.Background()
	}

	if loggingImpl.messenger == nil {
		panic("LoggingImpl.messenger is nil")
	}

	if loggingImpl.logger == nil {
		panic("LoggingImpl.logger is nil")
	}

	if loggingImpl.leveler == nil {
		panic("LoggingImpl.leveler is nil")
	}

	if loggingImpl.logLevelName == "" {
		loggingImpl.logLevelName = LevelInfoName
	}
}

// ----------------------------------------------------------------------------
// Private methods
// ----------------------------------------------------------------------------

func transformDetails(details ...interface{}) []interface{} {
	result := []interface{}{}

	for _, value := range details {
		switch typedValue := value.(type) {
		case MessageCode:
			result = append(result, messenger.MessageCode{Value: typedValue.Value})
		case MessageDuration:
			result = append(result, messenger.MessageDuration{Value: typedValue.Value})
		case MessageID:
			result = append(result, messenger.MessageID{Value: typedValue.Value})
		case MessageLevel:
			result = append(result, messenger.MessageID{Value: typedValue.Value})
		case MessageLocation:
			result = append(result, messenger.MessageLocation{Value: typedValue.Value})
		case MessageReason:
			result = append(result, messenger.MessageReason{Value: typedValue.Value})
		case MessageStatus:
			result = append(result, messenger.MessageStatus{Value: typedValue.Value})
		case MessageText:
			result = append(result, messenger.MessageText{Value: typedValue.Value})
		case MessageTime:
			result = append(result, messenger.MessageTime{Value: typedValue.Value})
		case OptionCallerSkip:
			result = append(result, messenger.OptionCallerSkip{Value: typedValue.Value})
		case time.Duration:
			result = append(result, messenger.MessageDuration{Value: typedValue.Nanoseconds()})
		default:
			result = append(result, typedValue)
		}
	}

	return result
}
