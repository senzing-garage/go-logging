/*
The MessageTimeSenzing implementation returns a time string in the format xxxx.
*/
package messagetime

import (
	"fmt"
	"time"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeSenzing type is for returning an empty string for date value.
type MessageTimeSenzing struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageTime method returns an empty string for a date value.
func (messageTime *MessageTimeSenzing) MessageTime(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	return fmt.Sprintf("%02d:%02d:%02d.%09d", messageTimestamp.UTC().Hour(), messageTimestamp.UTC().Minute(), messageTimestamp.Second(), messageTimestamp.Nanosecond()), nil
}
