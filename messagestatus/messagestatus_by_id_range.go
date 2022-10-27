/*
The MessageStatusByIdRange implementation returns a status based on
which range the message id falls in.
*/
package messagestatus

import (
	"fmt"
	"sort"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusByIdRange struct {
	IdRanges map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Get the "status" value given the message id.
func (messageStatus *MessageStatusByIdRange) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	// Create a list of reverse-sorted keys.

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

	// --- At this point, failed to find status -------------------------------

	return result, fmt.Errorf("could not determine status for message number %d", messageNumber)

}
