/*
The MessageLogLevelStatic implementation always returns the logger.Level of INFO.
*/
package messageloglevel

import (
	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLogLevelStatic type is for always returning the same log level.
type MessageLogLevelStatic struct {
	LogLevel logger.Level
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLogLevel method always return log level found in MessageLogLevelStatic.LogLevel.
func (messageLogLevel *MessageLogLevelStatic) MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	return messageLogLevel.LogLevel, err
}
