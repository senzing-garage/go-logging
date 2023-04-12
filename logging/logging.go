package logging

import (
	"context"
	"errors"
	"fmt"

	"github.com/senzing/go-messaging/messenger"
	"golang.org/x/exp/slog"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// loggingImpl is an type-struct for an implementation of the loggingInterface.
type LoggingImpl struct {
	Ctx          context.Context // Not a preferred practice, but used to simplify Log() calls.
	messenger    messenger.MessengerInterface
	logger       *slog.Logger
	leveler      *slog.LevelVar
	logLevelName string
}

// ----------------------------------------------------------------------------
// Private methods
// ----------------------------------------------------------------------------

func (loggingImpl *LoggingImpl) initialize() error {
	var err error = nil

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

	return err
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
func (loggingImpl *LoggingImpl) Error(messageNumber int, details ...interface{}) error {
	return errors.New(loggingImpl.Json(messageNumber, details...))
}

/*
The GetLogLevel method retrieves the current log level name.

Output
  - One of the following string values: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", or "PANIC"
*/
func (loggingImpl *LoggingImpl) GetLogLevel() string {
	return loggingImpl.logLevelName
}

func (loggingImpl *LoggingImpl) IsDebug() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelDebugName])
}

func (loggingImpl *LoggingImpl) IsError() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelErrorName])
}

func (loggingImpl *LoggingImpl) IsFatal() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelFatalName])
}

func (loggingImpl *LoggingImpl) IsInfo() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelInfoName])
}

func (loggingImpl *LoggingImpl) IsPanic() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelPanicName])
}

func (loggingImpl *LoggingImpl) IsTrace() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelTraceName])
}

func (loggingImpl *LoggingImpl) IsWarn() bool {
	return loggingImpl.logger.Enabled(loggingImpl.Ctx, TextToLevelMap[LevelWarnName])
}

/*
The Json method returns a JSON string based on the messageNumber and details.

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.

Output
  - JSON string with message key/value pairs.
*/
func (loggingImpl *LoggingImpl) Json(messageNumber int, details ...interface{}) string {
	return loggingImpl.messenger.NewJson(messageNumber, details...)
}

/*
The Log method writes a log record to the output specified at LoggingImpl creation

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.
*/
func (loggingImpl *LoggingImpl) Log(messageNumber int, details ...interface{}) {
	message, logLevel, details := loggingImpl.messenger.NewSlogLevel(messageNumber, details...)
	loggingImpl.logger.Log(loggingImpl.Ctx, logLevel, message, details...)
}

/*
The SetLogLevel method changes the level of log messages generated.

Input
  - logLevelName: One of these strings:  "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", or "PANIC".

Output
  - error
*/
func (loggingImpl *LoggingImpl) SetLogLevel(logLevelName string) error {
	var err error = nil
	slogLevel, ok := TextToLevelMap[logLevelName]
	if !ok {
		err := fmt.Errorf("unknown error level: %s", logLevelName)
		return err
	}
	loggingImpl.leveler.Set(slogLevel)
	loggingImpl.logLevelName = logLevelName
	return err
}
