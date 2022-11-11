/*
The MessageDateNull implementation returns an empty string for a date value.
*/
package messagedate

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDateNull type is for returning an empty string for date value.
type MessageDateNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDate method returns an empty string for a date value.
func (messageDate *MessageDateNull) MessageDate(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	return "", nil
}
