/*
The MessageDurationDefault implementation returns the number of milliseconds in the first
time.Duration seen in details.
*/
package messageduration

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDurationDefault type is for returning milliseconds from a time.Duration in details.
type MessageDurationDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDuration method returns number of milliseconds as a 64-bit integer.
func (messageDuration *MessageDurationDefault) MessageDuration(messageNumber int, details ...interface{}) (int64, error) {
	var err error = nil
	result := int64(0)

	if len(details) > 0 {
		for _, value := range details {
			switch typedValue := value.(type) {
			case time.Duration:
				return typedValue.Milliseconds(), err
			}
		}
	}

	return result, err
}
