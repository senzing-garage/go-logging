/*
The MessageLogLevelNull implementation always returns the logger.Level of INFO.
*/
package messageloglevel

import (
	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLogLevelStatic struct {
	LogLevel logger.Level
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Always return log level of INFO.
func (messageLogLevel *MessageLogLevelStatic) MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	return messageLogLevel.LogLevel, err
}
