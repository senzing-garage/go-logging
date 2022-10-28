/*
The MessageStatusById implementation returns a status based on the message number.
*/
package messagestatus

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusById type is for returning a status based on the message number.
type MessageStatusById struct {
	StatusTemplates map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageStatus method gets the "status" value from MessageStatusById.StatusTemplates for the given the message number.
func (messageStatus *MessageStatusById) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	if messageStatus.StatusTemplates != nil {
		result, ok := messageStatus.StatusTemplates[messageNumber]
		if ok {
			return result, err
		}
	}

	// --- At this point, failed to find status -------------------------------

	return result, fmt.Errorf("could not determine status for message number %d", messageNumber)
}
