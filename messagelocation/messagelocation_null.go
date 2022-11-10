/*
The MessageLocationNull implementation returns an empty string for a location value.
*/
package messagelocation

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLocationNull type is for returning an empty string for location value.
type MessageLocationNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLocation method returns an empty string for a location value.
func (messageLocation *MessageLocationNull) MessageLocation(messageNumber int, details ...interface{}) (string, error) {
	return "", nil
}
