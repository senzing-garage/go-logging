/*
The MessageStatusByIdRange implementation returns a status based on
which range the message id falls in.
*/
package messagestatus

import (
	"sort"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusByIdRange type is for determining a status based on what range a message number resides in.
type MessageStatusByIdRange struct {
	IdStatusRanges map[int]string // A map of "low-bound" of a range and the corresponding status message.
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageStatus method gets the "status" value from a range in MessageStatusByIdRange.IdRanges given the message id.
func (messageStatus *MessageStatusByIdRange) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	if messageStatus.IdStatusRanges != nil {
		// Create a list of reverse-sorted keys.

		messageLevelKeys := make([]int, 0, len(messageStatus.IdStatusRanges))
		for key := range messageStatus.IdStatusRanges {
			messageLevelKeys = append(messageLevelKeys, key)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(messageLevelKeys)))

		// Using the sorted message number, find the level.

		for _, messageLevelKey := range messageLevelKeys {
			if messageNumber >= messageLevelKey {
				return messageStatus.IdStatusRanges[messageLevelKey], err
			}
		}
	}

	// --- At this point, failed to find status -------------------------------

	return result, err

}
