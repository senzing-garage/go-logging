/*
The MessageErrorsNull implementation returns an empty value.
*/
package messageerrors

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsNull type is for returning an empty value.
type MessageErrorsNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageErrors method returns an empty value.
func (messageErrors *MessageErrorsNull) MessageErrors(messageNumber int, details ...interface{}) ([]interface{}, error) {
	return nil, nil
}
