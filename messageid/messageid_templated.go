/*
The MessageIdDefault implementation returns a message id based on Sprintf("%v").
*/
package messageid

import "fmt"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageIdTemplated struct {
	IdTemplate string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (messageId *MessageIdTemplated) MessageId(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	idTemplate := "%04d"
	if len(messageId.IdTemplate) > 0 {
		idTemplate = messageId.IdTemplate
	}
	result := fmt.Sprintf(idTemplate, messageNumber)
	return result, err
}
