/*
The MessageStatusSenzing implementation returns a status based on
which range the message id falls in.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusSenzing type is for determining a status based on what range a message number resides in.
type MessageStatusSenzing struct {
	IdStatuses map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageStatus method gets the "status" value from a range in MessageStatusSenzing.IdStatusRanges given the message id.
func (messageStatus *MessageStatusSenzing) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
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

	// --- At this point, failed to find status -------------------------------

	return result, err
}
