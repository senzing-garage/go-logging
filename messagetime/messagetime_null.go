/*
The MessageTimeNull implementation returns an empty string for a time value.
*/
package messagetime

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeNull type is for returning an empty string for time value.
type MessageTimeNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageTime method returns an empty string for a time value.
func (messageTime *MessageTimeNull) MessageTime(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	return "", nil
}
