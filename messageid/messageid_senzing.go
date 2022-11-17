/*
The MessageIdSenzing implementation returns a message id based on a format template string.
*/
package messageid

import "fmt"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageIdSenzing type is for creating message identifier based on a template.
type MessageIdSenzing struct {
	MessageIdTemplate string // Golang template string for formatting message identifier.
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageId method returns a string based on the MessageIdSenzing.IdTemplate.
func (messageId *MessageIdSenzing) MessageId(messageNumber int, details ...interface{}) (string, error) {
	return fmt.Sprintf(messageId.MessageIdTemplate, messageNumber), nil
}
