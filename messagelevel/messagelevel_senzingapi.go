/*
The MessageLevelSenzingApi implementation returns the logger.Level based on a logger.Level in details parameter,
a specific "status" value, or a message id in a range.
*/
package messagelevel

import (
	"sort"
	"sync"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagestatus"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelSenzingApi type calculates the logger.Level based on the status value.
type MessageLevelSenzingApi struct {
	DefaultLogLevel     logger.Level                         // User specified default.
	IdLevelRanges       map[int]logger.Level                 // The "low-bound" of a range and the corresponding logger level.
	IdStatuses          map[int]string                       // Passed to MessageStatusSenzingApi for message number to specific status.
	messageStatus       messagestatus.MessageStatusInterface // Local instance of MessageStatus
	lock                sync.Mutex                           // Lock for serializing creation of sortedIdLevelRanges.
	sortedIdLevelRanges []int                                // The keys of IdLevelRanges in sorted order.
}

// statusToLevelMap maps the constants in messagestatus_senzingapi.go to log levels.
var statusToLevelMap = map[string]logger.Level{
	messagestatus.Debug:              logger.LevelDebug,
	messagestatus.Error:              logger.LevelError,
	messagestatus.ErrorBadUserInput:  logger.LevelError,
	messagestatus.ErrorRetryable:     logger.LevelError,
	messagestatus.ErrorUnrecoverable: logger.LevelError,
	messagestatus.Fatal:              logger.LevelFatal,
	messagestatus.Info:               logger.LevelInfo,
	messagestatus.Panic:              logger.LevelPanic,
	messagestatus.Trace:              logger.LevelTrace,
	messagestatus.Warn:               logger.LevelWarn,
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (messageLevel *MessageLevelSenzingApi) getSortedIdLevelRanges() []int {
	if messageLevel.sortedIdLevelRanges == nil {
		messageLevel.lock.Lock()
		defer messageLevel.lock.Unlock()
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
// 2) the "status" value for the message id, 3) the range for the message id, or 4) the DefaultLogLevel.
func (messageLevel *MessageLevelSenzingApi) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
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

	// Second priority: Calculate log level from the status.

	if messageLevel.IdStatuses != nil {
		if messageLevel.messageStatus == nil {
			messageLevel.messageStatus = &messagestatus.MessageStatusSenzingApi{
				IdStatuses: messageLevel.IdStatuses,
			}
		}
		status, err := messageLevel.messageStatus.MessageStatus(messageNumber, details...)
		if err != nil {
			return result, err
		}
		result, ok := statusToLevelMap[status]
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
