/*
The MessageLevelByIdRange implementation returns the logger.Level based on the value of the message number.
*/
package messagelevel

import (
	"fmt"
	"sort"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelByIdRange type is for determining log level base on which range a message number resides in.
type MessageLevelByIdRange struct {
	DefaultLogLevel logger.Level
	IdLevelRanges   map[int]logger.Level
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLevel method returns a logger.level based on the message number.
func (messageLevel *MessageLevelByIdRange) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	result := messageLevel.DefaultLogLevel

	// First priority:  Log level explicitly given in details parameter.
	// Last occurance of logger.Level wins.

	foundInDetails := false
	for _, value := range details {
		switch typedValue := value.(type) {
		case logger.Level:
			foundInDetails = true
			result = typedValue
		}
	}
	if foundInDetails {
		return result, err
	}

	// Second priority: Message in a range.

	if messageLevel.IdLevelRanges != nil {

		// Since maps aren't sorted, create a list of sorted keys.

		messageLevelKeys := make([]int, 0, len(messageLevel.IdLevelRanges))
		for key := range messageLevel.IdLevelRanges {
			messageLevelKeys = append(messageLevelKeys, key)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(messageLevelKeys)))

		// Using the sorted message number, find the level.

		for _, messageLevelKey := range messageLevelKeys {
			if messageNumber >= messageLevelKey {
				return messageLevel.IdLevelRanges[messageLevelKey], err
			}
		}
	}

	err = fmt.Errorf("could not find error range for message number %d. Setting to level %d", messageNumber, result)
	return result, err
}
