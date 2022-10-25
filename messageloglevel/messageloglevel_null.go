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

type MessageLogLevelNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Always return log level of INFO.
func (messagelevel *MessageLogLevelNull) MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	return logger.LevelInfo, err
}
