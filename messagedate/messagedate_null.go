/*
The MessageDateNull implementation returns an empty string for a date value.
*/
package messagedate

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDateNull type is for returning an empty string for date value.
type MessageDateNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDate method returns an empty string for a date value.
func (messageDate *MessageDateNull) MessageDate(messageNumber int, details ...interface{}) (string, error) {
	return "", nil
}
