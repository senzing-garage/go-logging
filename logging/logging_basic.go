package logging

import (
	"context"
	"errors"
	"fmt"

	"github.com/senzing-garage/go-messaging/messenger"
	"golang.org/x/exp/slog"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicLogging is an type-struct for an implementation of the loggingInterface.
type BasicLogging struct {
	Ctx          context.Context // Not a preferred practice, but used to simplify Log() calls.
	messenger    messenger.Interface
	logger       *slog.Logger
	leveler      *slog.LevelVar
	logLevelName string
}

// ----------------------------------------------------------------------------
// Private methods
// ----------------------------------------------------------------------------

func (loggingImpl *BasicLogging) initialize() error {
	var err error

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
func (loggingImpl *BasicLogging) NewError(messageNumber int, details ...interface{}) error {
	return errors.New(loggingImpl.messenger.NewJSON(messageNumber, details...))
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
	return loggingImpl.messenger.NewJSON(messageNumber, details...)
}

/*
The Log method writes a log record to the output specified at LoggingImpl creation

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.
*/
func (loggingImpl *BasicLogging) Log(messageNumber int, details ...interface{}) {
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
func (loggingImpl *BasicLogging) SetLogLevel(logLevelName string) error {
	var err error
	slogLevel, ok := TextToLevelMap[logLevelName]
	if !ok {
		err := fmt.Errorf("unknown error level: %s", logLevelName)
		return err
	}
	loggingImpl.leveler.Set(slogLevel)
	loggingImpl.logLevelName = logLevelName
	return err
}
