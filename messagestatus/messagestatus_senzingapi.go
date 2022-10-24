/*
Package helper ...
*/
package messagestatus

import (
	"errors"
	"sort"
	"strings"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusSenzingApi struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// --- Exception hierarchy ----------------------------------------------------

const (
	ErrorRetryable     = "retryable"
	ErrorBadUserInput  = "bad-input"
	ErrorUnrecoverable = "unrecoverable"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var senzingApiErrorsMap = map[string]string{
	"0002E":  logger.LevelInfoName,
	"0019E":  ErrorUnrecoverable,
	"0099E":  ErrorRetryable, // Fake error.
	"0007E":  "error",
	"0023E":  "error",
	"0024E":  "error",
	"0025E":  "error",
	"0026E":  "error",
	"0027E":  "error",
	"0032E":  "error",
	"0034E":  "error",
	"0035E":  "error",
	"0036E":  "error",
	"0048E":  "fatal",
	"0051E":  "error",
	"0053E":  "fatal",
	"0054E":  "error",
	"0061E":  "error",
	"0062E":  "error",
	"0064E":  "error",
	"1007E":  "error",
	"2134E":  "error",
	"30020":  "error",
	"30103E": "error",
	"30110E": "error",
	"30111E": "error",
	"30112E": "error",
	"30121E": "error",
	"30122E": "error",
	"30123E": "error",
	"9000E":  "error",
}

// Important:  The number listed is one more than the highest number for the MessageLevel.
// This should be kept in sync with go-logging/logger/main.go

// Message ranges:
// 0000-0999 info
// 1000-1999 warning
// 2000-2999 error
// 3000-3999 debug
// 4000-4999 trace
// 5000-5999 fatal
// 6000-6999 panic
var messageLevelMap = map[int]logger.Level{
	1000: logger.LevelInfo,
	2000: logger.LevelWarn,
	3000: logger.LevelError,
	4000: logger.LevelDebug,
	5000: logger.LevelTrace,
	6000: logger.LevelFatal,
	7000: logger.LevelPanic,
}

var messageLevelToStringMap = map[logger.Level]string{
	logger.LevelInfo:  logger.LevelInfoName,
	logger.LevelWarn:  logger.LevelWarnName,
	logger.LevelError: logger.LevelErrorName,
	logger.LevelDebug: logger.LevelDebugName,
	logger.LevelTrace: logger.LevelTraceName,
	logger.LevelFatal: logger.LevelFatalName,
	logger.LevelPanic: logger.LevelPanicName,
}

var MessagePrecedence = []string{
	logger.LevelPanicName,
	logger.LevelFatalName,
	ErrorUnrecoverable,
	ErrorBadUserInput,
	ErrorRetryable,
	logger.LevelErrorName,
	logger.LevelWarnName,
	logger.LevelInfoName,
	logger.LevelDebugName,
	logger.LevelTraceName,
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageStatusSenzingApi) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	// --- Status based on Senzing error number -------------------------------

	// Create a list of Senzing errors.

	var senzingErrors []string
	if len(details) > 0 {
		for index := len(details) - 1; index >= 0; index-- {
			detail := details[index]
			switch typedDetail := detail.(type) {
			case error:
				errorMessage := typedDetail.Error()
				messageSplits := strings.Split(errorMessage, "|")
				for key, value := range senzingApiErrorsMap {
					if messageSplits[0] == key {
						senzingErrors = append(senzingErrors, value)
					}
				}
			}
		}
	}

	// In the list of Senzing errors, determine the highest priority error.

	if len(senzingErrors) > 0 {
		for _, MessagePrecedenceLevel := range MessagePrecedence {
			for _, senzingError := range senzingErrors {
				if senzingError == MessagePrecedenceLevel {
					return senzingError, err
				}
			}
		}
	}

	// --- Status based on messageNumber ----------------------------------------

	// Create a list of sorted keys.

	messageLevelKeys := make([]int, 0, len(messageLevelMap))
	for key := range messageLevelMap {
		messageLevelKeys = append(messageLevelKeys, key)
	}
	sort.Ints(messageLevelKeys)

	// Using the sorted message number, find the level.

	finalMessageLevel := logger.LevelPanic
	for _, messageLevelKey := range messageLevelKeys {
		if messageNumber < messageLevelKey {
			finalMessageLevel = messageLevelMap[messageLevelKey]
			break
		}
	}

	result, ok := messageLevelToStringMap[finalMessageLevel]
	if ok {
		return result, err
	}

	// --- At this point, failed to find status -------------------------------

	return "", errors.New("could not determine status")
}
