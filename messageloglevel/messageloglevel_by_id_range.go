/*
The MessageLogLevelById implementation returns the logger.Level based on the value of the message number.
*/
package messageloglevel

import (
	"fmt"
	"sort"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLogLevelByIdRange struct {
	DefaultLogLevel logger.Level
	IdRanges        map[int]logger.Level
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Return a logger.level based on the message number.
func (messageLogLevel *MessageLogLevelByIdRange) MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	result := logger.LevelPanic

	// Create a list of sorted keys.

	messageLevelKeys := make([]int, 0, len(messageLogLevel.IdRanges))
	for key := range messageLogLevel.IdRanges {
		messageLevelKeys = append(messageLevelKeys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(messageLevelKeys)))

	// Using the sorted message number, find the level.

	for _, messageLevelKey := range messageLevelKeys {
		if messageNumber >= messageLevelKey {
			return messageLogLevel.IdRanges[messageLevelKey], err
		}
	}

	result = messageLogLevel.DefaultLogLevel
	err = fmt.Errorf("could not find error range for message number %d. Setting to level %d", messageNumber, result)
	return result, err
}
