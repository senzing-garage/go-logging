/*
Package helper ...
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

// TODO:
func (messagelevel *MessageLogLevelNull) CalculateMessageLogLevel(errorNumber int, message string) (logger.Level, error) {
	var err error = nil
	return logger.LevelWarn, err
}
