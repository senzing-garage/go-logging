/*
The MessageLevelSenzingApi implementation returns the logger.Level based on the "status" value.
*/
package messagelevel

import (
	"sort"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagestatus"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelSenzingApi type is for calculating the log level based on the status value.
type MessageLevelSenzingApi struct {
	IdLevelRanges       map[int]logger.Level
	IdStatuses          map[int]string
	sortedIdLevelRanges []int // The keys of IdLevelRanges in sorted order.
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

// The MessageLevel method returns a log level based on the status value.
func (messageLevel *MessageLevelSenzingApi) MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	var result = logger.LevelError

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

	// Second priority: Calculate log level from the status.

	if messageLevel.IdStatuses != nil {
		messageStatus := &messagestatus.MessageStatusSenzingApi{
			IdStatuses: messageLevel.IdStatuses,
		}
		status, err := messageStatus.MessageStatus(messageNumber, details...)
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

	// Anything else is an "ERROR"

	return logger.LevelError, err
}
