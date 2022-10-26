/*
The MessageStatusById implementation returns a status based on the message id.
*/
package messagestatus

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusById struct {
	StatusTemplates map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Get the "status" value given the message id.
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
