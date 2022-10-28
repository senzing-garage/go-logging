/*
The MessageStatusSenzingApi implementation calculates a status value based on message id and Senzing return code.
*/
package messagestatus

import (
	"fmt"
	"sort"
	"strings"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

/*
The MessageStatusSenzingApi type is for constructing status values by first
looking at the Senzing error code.
If it doesn't exist, use the messageNumber to calculate a status.
*/
type MessageStatusSenzingApi struct {
	IdRanges map[int]string
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Types of Senzing errors
const (
	ErrorRetryable     = "retryable"
	ErrorBadUserInput  = "bad-input"
	ErrorUnrecoverable = "unrecoverable"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A map of Senzing errors to the corresponding error level.
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

var messagePrecedence = []string{
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

// The MessageStatus method returns a status based on a message number indexed into senzingApiErrorsMap.
func (messageStatus *MessageStatusSenzingApi) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	// --- Status based on Senzing error passed in via details ----------------

	// Create a list of Senzing errors by looking at details in reverse order.

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
		for _, messagePrecedenceLevel := range messagePrecedence {
			for _, senzingError := range senzingErrors {
				if senzingError == messagePrecedenceLevel {
					return senzingError, err
				}
			}
		}
	}

	// --- Status based on messageNumber ----------------------------------------

	if messageStatus.IdRanges != nil {

		// Create a list of sorted keys.

		messageLevelKeys := make([]int, 0, len(messageStatus.IdRanges))
		for key := range messageStatus.IdRanges {
			messageLevelKeys = append(messageLevelKeys, key)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(messageLevelKeys)))

		// Using the sorted message number, find the level.

		for _, messageLevelKey := range messageLevelKeys {
			if messageNumber >= messageLevelKey {
				return messageStatus.IdRanges[messageLevelKey], err
			}
		}
	}

	// --- At this point, failed to find status -------------------------------

	return result, fmt.Errorf("could not determine status for message number %d", messageNumber)
}
