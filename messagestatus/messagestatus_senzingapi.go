/*
The MessageStatusSenzingApi implementation calculates a status value based on message id and Senzing return code.
*/
package messagestatus

import (
	"strings"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

/*
The MessageStatusSenzingApi type is for constructing status values for the SenzingAPI.
First, it looks in the details for an explicit messagestatus.Status value.
Next, it sees if the message number is in a lookup table.
Finally, it looks at the Senzing error code.
If in none of those places, an empty string is returned.
*/
type MessageStatusSenzingApi struct {
	IdStatuses map[int]string // A map of message ids to the corresponding status message.
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

/*
Types of Senzing errors.
These are the strings that may be returned from MessageStatus()
*/
const (
	Debug              = logger.LevelDebugName
	Error              = logger.LevelErrorName
	ErrorBadUserInput  = logger.LevelErrorName + "_bad_user_input"
	ErrorRetryable     = logger.LevelErrorName + "_retryable"
	ErrorUnrecoverable = logger.LevelErrorName + "_unrecoverable"
	Fatal              = logger.LevelFatalName
	Info               = logger.LevelInfoName
	Panic              = logger.LevelPanicName
	Trace              = logger.LevelTraceName
	Warn               = logger.LevelWarnName
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A map of Senzing errors to the corresponding error level.
var senzingApiErrorsMap = map[string]string{
	"0002E":  Info,
	"0019E":  ErrorUnrecoverable,
	"0037E":  ErrorBadUserInput,  // Unknown resolved entity value
	"0052E":  ErrorBadUserInput,  // Unknown relationship ID value
	"0063E":  ErrorUnrecoverable, // G2ConfigMgr is not initialized
	"7221E":  ErrorUnrecoverable, // No engine configuration registered
	"30121E": ErrorBadUserInput,  // JSON parsing Failure
}

// The order of severity/verbosity from most severe to most verbose.
var messagePrecedence = []string{
	Panic,
	Fatal,
	ErrorUnrecoverable,
	ErrorBadUserInput,
	ErrorRetryable,
	Error,
	Warn,
	Info,
	Debug,
	Trace,
}

func (messageStatus *MessageStatusSenzingApi) messageStatusBySenzingError(messageNumber int, details ...interface{}) string {

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
					return senzingError
				}
			}
		}
	}
	return ""
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageStatus method returns a status based on a message number indexed into senzingApiErrorsMap.
func (messageStatus *MessageStatusSenzingApi) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	// First priority:  Status explicitly given in details parameter.
	// Last occurance of messagestatus.Status is used.

	foundInDetails := false
	for _, value := range details {
		switch typedValue := value.(type) {
		case Status:
			foundInDetails = true
			result = string(typedValue)
		}
	}
	if foundInDetails {
		return result, err
	}

	// Second priority: Status based on message number lookup.

	if messageStatus.IdStatuses != nil {
		result, ok := messageStatus.IdStatuses[messageNumber]
		if ok {
			return result, err
		}
	}

	// Third priority: Status based on Senzing error passed in via details.

	result = messageStatus.messageStatusBySenzingError(messageNumber, details...)
	if len(result) > 0 {
		return result, err
	}

	// --- At this point, failed to find status -------------------------------

	return result, err
}
