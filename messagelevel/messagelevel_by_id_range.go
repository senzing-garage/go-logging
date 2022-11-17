/*
The MessageLevelByIdRange implementation returns the logger.Level based on the value of the message number
compared to a list of ranges.
*/
package messagelevel

import (
	"sort"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelByIdRange type is for determining log level base on which range a message number resides in.
type MessageLevelByIdRange struct {
	DefaultLogLevel     logger.Level         // User specified default.
	IdLevelRanges       map[int]logger.Level // The "low-bound" of a range and the corresponding logger level.
	sortedIdLevelRanges []int                // The keys of IdRanges in sorted order.
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (messageLevel *MessageLevelByIdRange) getSortedIdLevelRanges() []int {
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
// 2) the range for the message id, or 3) the DefaultLogLevel.
func (messageLevel *MessageLevelByIdRange) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
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

	// Second priority: Message in a range.

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
