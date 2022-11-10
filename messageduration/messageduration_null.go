/*
The MessageDurationNull implementation returns an empty value for a duration value.
*/
package messageduration

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDurationNull type is for returning an empty string for duration value.
type MessageDurationNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDuration method returns an empty string for a duration value.
func (messageDuration *MessageDurationNull) MessageDuration(messageNumber int, details ...interface{}) (int64, error) {
	return 0, nil
}
