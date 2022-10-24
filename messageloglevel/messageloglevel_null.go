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
func (messagelevel *MessageLogLevelNull) MessageLogLevel(errorNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	return logger.LevelInfo, err
}
