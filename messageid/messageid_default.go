/*
The MessageIdDefault implementation returns a message id based on Sprintf("%d").
*/
package messageid

import "fmt"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageIdDefault type is for a simple integer message identifier.
type MessageIdDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageId method simply returns a string of the integer passed in.
func (messageId *MessageIdDefault) MessageId(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	result := fmt.Sprintf("%d", messageNumber)
	return result, err
}
