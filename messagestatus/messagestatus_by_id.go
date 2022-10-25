/*
The MessageStatusById implementation returns a status based on the message id.
Message ranges:

	0000-0999 INFO
	1000-1999 WARN
	2000-2999 ERROR
	3000-3999 DEBUG
	4000-4999 TRACE
	5000-5999 FATAL
	6000-6999 PANIC
*/
package messagestatus

import (
	"errors"
	"sort"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusById struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

/*
The numbers used to divide the logging levels.

Important:  The number listed is one more than the highest number for the MessageLevel.
This should be kept in sync with go-logging/logger/main.go
Message ranges:
0000-0999 INFO
1000-1999 WARN
2000-2999 ERROR
3000-3999 DEBUG
4000-4999 TRACE
5000-5999 FATAL
6000-6999 PANIC
*/
var messageLevelMapById = map[int]logger.Level{
	1000: logger.LevelInfo,
	2000: logger.LevelWarn,
	3000: logger.LevelError,
	4000: logger.LevelDebug,
	5000: logger.LevelTrace,
	6000: logger.LevelFatal,
	7000: logger.LevelPanic,
}

var messageLevelToStringMapById = map[logger.Level]string{
	logger.LevelInfo:  logger.LevelInfoName,
	logger.LevelWarn:  logger.LevelWarnName,
	logger.LevelError: logger.LevelErrorName,
	logger.LevelDebug: logger.LevelDebugName,
	logger.LevelTrace: logger.LevelTraceName,
	logger.LevelFatal: logger.LevelFatalName,
	logger.LevelPanic: logger.LevelPanicName,
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Get the "status" value given the message id and it's details.
func (messagelevel *MessageStatusById) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	// Create a list of sorted keys.

	messageLevelKeys := make([]int, 0, len(messageLevelMapById))
	for key := range messageLevelMapById {
		messageLevelKeys = append(messageLevelKeys, key)
	}
	sort.Ints(messageLevelKeys)

	// Using the sorted message number, find the level.

	finalMessageLevel := logger.LevelPanic
	for _, messageLevelKey := range messageLevelKeys {
		if messageNumber < messageLevelKey {
			finalMessageLevel = messageLevelMapById[messageLevelKey]
			break
		}
	}

	result, ok := messageLevelToStringMapById[finalMessageLevel]
	if ok {
		return result, err
	}

	// --- At this point, failed to find status -------------------------------

	return "", errors.New("could not determine status")
}
