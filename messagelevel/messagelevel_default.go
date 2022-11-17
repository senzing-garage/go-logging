/*
The MessageLevelDefault implementation returns the logger.Level based on a logger.Level in details parameter,
a specific message id, or a message id in a range.
*/
package messagelevel

import (
	"sort"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelDefault type returns the logger.Level based on the last logger.Level in details parameter.
type MessageLevelDefault struct {
	DefaultLogLevel     logger.Level         // User specified default.
	IdLevels            map[int]logger.Level // Specific message ids and the corresponding logger level.
	IdLevelRanges       map[int]logger.Level // The "low-bound" of a range and the corresponding logger level.
	sortedIdLevelRanges []int                // The keys of IdLevelRanges in sorted order.
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (messageLevel *MessageLevelDefault) getSortedIdLevelRanges() []int {
	if messageLevel.sortedIdLevelRanges == nil {
		messageLevel.sortedIdLevelRanges = make([]int, 0, len(messageLevel.IdLevelRanges))
		for key := range messageLevel.IdLevelRanges {
			messageLevel.sortedIdLevelRanges = append(messageLevel.sortedIdLevelRanges, key)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(messageLevel.sortedIdLevelRanges)))
	}
	return messageLevel.sortedIdLevelRanges
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLevel method returns a logger.Level based on: 1) the highest logger.Level types in the details parameter,
// 2) a specific message for the message id, 3) the range for the message id, or 4) the DefaultLogLevel.
func (messageLevel *MessageLevelDefault) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	var result = messageLevel.DefaultLogLevel

	// First priority:  Log level explicitly given in details parameter.
	// Highest value of logger.Level is used.

	var explicitResult logger.Level
	foundInDetails := false
	for _, value := range details {
		switch typedValue := value.(type) {
		case logger.Level:
			foundInDetails = true
			if typedValue > explicitResult {
				explicitResult = typedValue
			}
		}
	}
	if foundInDetails {
		return explicitResult, err
	}

	// Second priority: Message Id exact match to an entry in IdLevels.

	if messageLevel.IdLevels != nil {
		result, ok := messageLevel.IdLevels[messageNumber]
		if ok {
			return result, err
		}
	}

	// Third priority: Message in a range.

	if messageLevel.IdLevelRanges != nil {
		sortedMessageLevelKeys := messageLevel.getSortedIdLevelRanges()
		for _, messageLevelKey := range sortedMessageLevelKeys {
			if messageNumber >= messageLevelKey {
				return messageLevel.IdLevelRanges[messageLevelKey], err
			}
		}
	}

	// Last priority, the default value.

	return result, err
}
