/*
The MessageTextNull implementation always returns an empty string.
*/
package messagetext

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

/*
MessageTextNull always returns an empty string.
*/
type MessageTextNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// MessageText always returns an empty string.
func (messageText *MessageTextNull) MessageText(messageNumber int, details ...interface{}) (string, error) {
	return "", nil
}
