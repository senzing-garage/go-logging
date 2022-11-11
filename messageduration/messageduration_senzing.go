/*
The MessageDurationSenzing implementation returns the number of nanoseconds in the first
time.Duration seen in details.
*/
package messageduration

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDurationSenzing type is for returning nanoseconds from a time.Duration in details.
type MessageDurationSenzing struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDuration method returns number of nanoseconds as a 64-bit integer.
func (messageDuration *MessageDurationSenzing) MessageDuration(messageNumber int, details ...interface{}) (int64, error) {
	var err error = nil
	result := int64(0)

	if len(details) > 0 {
		for _, value := range details {
			switch typedValue := value.(type) {
			case time.Duration:
				return typedValue.Nanoseconds(), err
			}
		}
	}

	return result, err
}
