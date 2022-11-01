/*
The MessageIdTemplated implementation returns a message id based on a format template string.
*/
package messageid

import "fmt"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageIdTemplated type is for creating message identifier based on a template.
type MessageIdTemplated struct {
	MessageIdTemplate string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageId method returns a string based on the MessageIdTemplated.IdTemplate.
// If IdTemplate is not set, the value "%04d" is used.
func (messageId *MessageIdTemplated) MessageId(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	messageIdTemplate := "%04d"
	if len(messageId.MessageIdTemplate) > 0 {
		messageIdTemplate = messageId.MessageIdTemplate
	}
	result := fmt.Sprintf(messageIdTemplate, messageNumber)
	return result, err
}
