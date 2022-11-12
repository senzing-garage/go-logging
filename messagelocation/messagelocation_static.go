/*
The MessageLocationStatic implementation returns a fixed string for a location value.
Used mostly for repeatable test cases.
*/
package messagelocation

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLocationStatic type is for returning a fixed string for a location value.
type MessageLocationStatic struct {
	Location string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLocation method returnsa fixed string for a location value.
func (messageLocation *MessageLocationStatic) MessageLocation(messageNumber int, details ...interface{}) (string, error) {
	return messageLocation.Location, nil
}
