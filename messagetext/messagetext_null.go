/*
The MessageTextNull implementation always returns an empty string.
*/
package messagetext

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTextNull type always returns an empty string.
type MessageTextNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageText method always returns an empty string.
func (messageText *MessageTextNull) MessageText(messageNumber int, details ...interface{}) (string, error) {
	return "", nil
}
