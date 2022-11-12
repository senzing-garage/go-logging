/*
The MessageLevelDefault implementation returns the logger.Level based on a logger.Level in demessagelogleveltails parameter,
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
type MessageLevelSenzing struct {
	DefaultLogLevel        logger.Level
	IdLevels               map[int]logger.Level // Specific message ids and the corresponding logger level.
	IdRanges               map[int]logger.Level // The "low-bound" of a range and the corresponding logger level.
	sortedMessageLevelKeys []int                // The keys of IdRanges in sorted order.
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (messageLogLevel *MessageLevelSenzing) getSortedMessageLevelKeys() []int {
	if messageLogLevel.sortedMessageLevelKeys == nil {
		messageLogLevel.sortedMessageLevelKeys = make([]int, 0, len(messageLogLevel.IdRanges))
		for key := range messageLogLevel.IdRanges {
			messageLogLevel.sortedMessageLevelKeys = append(messageLogLevel.sortedMessageLevelKeys, key)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(messageLogLevel.sortedMessageLevelKeys)))
	}
	return messageLogLevel.sortedMessageLevelKeys
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLevel method returns a logger.level based on one or more logger.level types in the details parameter.
func (messageLogLevel *MessageLevelSenzing) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil

	// First priority:  Log level explicitly given in details parameter.

	for _, value := range details {
		switch typedValue := value.(type) {
		case logger.Level:
			return typedValue, err
		}
	}

	// Second priority: Message Id exact match to an entry in IdLevels.

	if messageLogLevel.IdLevels != nil {
		result, ok := messageLogLevel.IdLevels[messageNumber]
		if ok {
			return result, err
		}
	}

	// Third priority: Message in a range.

	if messageLogLevel.IdRanges != nil {
		sortedMessageLevelKeys := messageLogLevel.getSortedMessageLevelKeys()
		for _, messageLevelKey := range sortedMessageLevelKeys {
			if messageNumber >= messageLevelKey {
				return messageLogLevel.IdRanges[messageLevelKey], err
			}
		}
	}

	// Last priority, the default value.

	return messageLogLevel.DefaultLogLevel, err
}
