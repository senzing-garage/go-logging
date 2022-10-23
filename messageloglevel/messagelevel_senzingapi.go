/*
Package helper ...
*/
package messageloglevel

import (
	"sort"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLogLevelSenzingApi struct{}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Important:  The number listed is one more than the highest number for the MessageLevel.
// Message ranges:
// 0000-0999 info
// 1000-1999 warning
// 2000-2999 error
// 3000-3999 debug
// 4000-4999 trace
// 5000-5999 fatal
// 6000-6999 panic
var MessageLevelMap = map[int]logger.Level{
	1000: logger.LevelInfo,
	2000: logger.LevelWarn,
	3000: logger.LevelError,
	4000: logger.LevelDebug,
	5000: logger.LevelTrace,
	6000: logger.LevelFatal,
	7000: logger.LevelPanic,
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageLogLevelSenzingApi) MessageLogLevel(errorNumber int, message string) (logger.Level, error) {
	var err error = nil
	var result = logger.LevelPanic

	// Create a list of sorted keys.

	messageLevelKeys := make([]int, 0, len(MessageLevelMap))
	for key := range MessageLevelMap {
		messageLevelKeys = append(messageLevelKeys, key)
	}
	sort.Ints(messageLevelKeys)

	// Using the sorted message number, find the level.

	for _, messageLevelKey := range messageLevelKeys {
		if errorNumber < messageLevelKey {
			result = MessageLevelMap[messageLevelKey]
			break
		}
	}

	return result, err
}
