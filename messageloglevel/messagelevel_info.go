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

type MessageLogLevelInfo struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageLogLevelInfo) MessageLogLevel(errorNumber int, message string) (logger.Level, error) {
	var err error = nil
	return logger.LevelInfo, err
}
