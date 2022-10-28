/*
The MessageStatusNull implementation returns an empty string for a status value.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusNull type is for returning an empty string for status value.
type MessageStatusNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageStatus method returns an empty string for a status value.
func (messageStatus *MessageStatusNull) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	return "", nil
}
