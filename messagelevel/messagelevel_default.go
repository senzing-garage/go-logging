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

// The MessageLevelDefault type returns the logger.Level based on a any logger.Level in details parameter.
type MessageLevelDefault struct {
	DefaultLogLevel        logger.Level
	IdLevels               map[int]logger.Level // Specific message ids and the corresponding logger level.
	IdRanges               map[int]logger.Level // The "low-bound" of a range and the corresponding logger level.
	sortedMessageLevelKeys []int                // The keys of IdRanges in sorted order.
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (messageLevel *MessageLevelDefault) getSortedMessageLevelKeys() []int {
	if messageLevel.sortedMessageLevelKeys == nil {
		messageLevel.sortedMessageLevelKeys = make([]int, 0, len(messageLevel.IdRanges))
		for key := range messageLevel.IdRanges {
			messageLevel.sortedMessageLevelKeys = append(messageLevel.sortedMessageLevelKeys, key)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(messageLevel.sortedMessageLevelKeys)))
	}
	return messageLevel.sortedMessageLevelKeys
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLevel method returns a logger.level based on one or more logger.level types in the details parameter.
func (messageLevel *MessageLevelDefault) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	var result logger.Level

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

	// Second priority: Message Id exact match to an entry in IdLevels.

	if messageLevel.IdLevels != nil {
		result, ok := messageLevel.IdLevels[messageNumber]
		if ok {
			return result, err
		}
	}

	// Third priority: Message in a range.

	if messageLevel.IdRanges != nil {
		sortedMessageLevelKeys := messageLevel.getSortedMessageLevelKeys()
		for _, messageLevelKey := range sortedMessageLevelKeys {
			if messageNumber >= messageLevelKey {
				return messageLevel.IdRanges[messageLevelKey], err
			}
		}
	}

	// Last priority, the default value.

	return messageLevel.DefaultLogLevel, err
}
