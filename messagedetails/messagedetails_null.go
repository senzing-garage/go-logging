/*
The MessageDetailsNull implementation returns an empty value.
*/
package messagedetails

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsNull type is for returning an empty value.
type MessageDetailsNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDetails method returns an empty value.
func (messageDetails *MessageDetailsNull) MessageDetails(messageNumber int, details ...interface{}) (interface{}, error) {
	return nil, nil
}
