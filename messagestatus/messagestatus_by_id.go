/*
The MessageStatusById implementation returns a status based on the message number.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusById type is for returning a status based on the message number.
type MessageStatusById struct {
	IdStatuses map[int]string // A map of message ids to the corresponding status message.
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageStatus method gets the "status" value from MessageStatusById.IdStatus for the given the message number.
func (messageStatus *MessageStatusById) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	if messageStatus.IdStatuses != nil {
		result, ok := messageStatus.IdStatuses[messageNumber]
		if ok {
			return result, err
		}
	}

	// --- At this point, failed to find status -------------------------------

	return result, err
}
