/*
The MessageTimeDefault implementation returns a time string in the format HH-MM-SS.mmmmmm.
*/
package messagetime

import (
	"fmt"
	"time"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeDefault type is for returning a time string in the format HH-MM-SS.mmmmmm.
type MessageTimeDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageTime method returns a time string in the format HH-MM-SS.mmmmmm.
func (messageTime *MessageTimeDefault) MessageTime(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	return fmt.Sprintf("%02d:%02d:%02d.%06d", messageTimestamp.UTC().Hour(), messageTimestamp.UTC().Minute(), messageTimestamp.Second(), messageTimestamp.Nanosecond()/1000), nil
}
