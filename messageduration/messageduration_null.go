/*
The MessageDurationNull implementation returns a zero alue for a duration value.
*/
package messageduration

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDurationNull type is for returning a zero alue for a duration value.
type MessageDurationNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDuration method returns a zero alue for a duration value.
func (messageDuration *MessageDurationNull) MessageDuration(messageNumber int, details ...interface{}) (int64, error) {
	return 0, nil
}
