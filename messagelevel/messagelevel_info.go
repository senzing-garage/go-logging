/*
Package helper ...
*/
package messagelevel

import (
	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLevelInfo struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageLevelInfo) CalculateMessageLevel(errorNumber int, message string) (logger.Level, error) {
	var err error = nil
	return logger.LevelInfo, err
}
