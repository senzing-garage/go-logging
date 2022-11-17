/*
The MessageLevelStatic implementation always returns the logger.Level set by the MessageLevelStatic.LogLevel value.
*/
package messagelevel

import (
	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelStatic type is for always returning the same log level.
type MessageLevelStatic struct {
	LogLevel logger.Level // User specified default.
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLevel method always return logger.Level found in MessageLevelStatic.LogLevel.
func (messageLogLevel *MessageLevelStatic) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	return messageLogLevel.LogLevel, err
}
