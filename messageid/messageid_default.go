/*
The MessageIdDefault implementation returns a message id based on a message template.
*/
package messageid

import "fmt"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageIdDefault struct {
	IdTemplate string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (messageId *MessageIdDefault) MessageId(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	idTemplate := "%04d"
	if len(messageId.IdTemplate) > 0 {
		idTemplate = messageId.IdTemplate
	}
	result := fmt.Sprintf(idTemplate, messageNumber)
	return result, err
}

func (messagetext *MessageIdDefault) SetIdTemplate(idTemplate string) {
	messagetext.IdTemplate = idTemplate
}
