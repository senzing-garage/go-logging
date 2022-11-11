/*
The MessageTimeDefault implementation returns a time string in the format xxxx.
*/
package messagetime

import (
	"fmt"
	"time"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeDefault type is for returning an empty string for date value.
type MessageTimeDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageTime method returns an empty string for a date value.
func (messageTime *MessageTimeDefault) MessageTime(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	return fmt.Sprintf("%02d:%02d:%02d.%06d", messageTimestamp.UTC().Hour(), messageTimestamp.UTC().Minute(), messageTimestamp.Second(), messageTimestamp.UnixMilli()), nil
}
