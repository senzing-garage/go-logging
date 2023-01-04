/*
The MessageStatusSenzingApi implementation calculates a status value based on message id and Senzing return code.
*/
package messagestatus

import (
	"strings"
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
// Interface methods
// ----------------------------------------------------------------------------

func (messageStatus *MessageStatusSenzingApi) messageStatusBySenzingError(messageNumber int, details ...interface{}) string {

	// Create a list of Senzing errors by looking at details.

	var senzingErrors []string
	for _, detail := range details {
		switch typedDetail := detail.(type) {
		case error:
			errorMessage := typedDetail.Error()
			messageSplits := strings.Split(errorMessage, "|")
			for key, value := range SenzingApiErrorsMap {
				if messageSplits[0] == key {
					senzingErrors = append(senzingErrors, value)
				}
			}
		}
	}

	// In the list of Senzing errors, determine the highest priority error.

	if len(senzingErrors) > 0 {
		for _, messagePrecedenceLevel := range MessagePrecedence {
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
	// Last occurrence of messagestatus.Status is used.

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
