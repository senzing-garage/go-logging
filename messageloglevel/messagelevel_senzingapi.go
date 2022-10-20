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

type MessageLogLevelSenzingApi struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageLogLevelSenzingApi) CalculateMessageLogLevel(errorNumber int, message string) (logger.Level, error) {
	var err error = nil
	return logger.LevelError, err
}
