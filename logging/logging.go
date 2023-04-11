package logging

import (
	"context"
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
The Log method ...

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.

Output
  - error
*/
func (loggingImpl *LoggingImpl) Log(messageNumber int, details ...interface{}) {
	message, logLevel, details := loggingImpl.messenger.NewSlogLevel(messageNumber, details...)
	loggingImpl.logger.Log(loggingImpl.Ctx, logLevel, message, details...)
}

/*
The GetLogLevel method retrieves the current log level name.

Output
  - One of the following string values: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", or "PANIC"
*/
func (loggingImpl *LoggingImpl) GetLogLevel() string {
	return loggingImpl.logLevelName
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
