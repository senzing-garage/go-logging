/*
Package helper ...
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

// TODO:
func (messageId *MessageIdDefault) MessageId(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	idTemplate := "%04d"
	if len(messageId.IdTemplate) > 0 {
		idTemplate = messageId.IdTemplate
	}
	result := fmt.Sprintf(idTemplate, messageNumber)
	return result, err
}

// TODO:
func (messagetext *MessageIdDefault) SetIdTemplate(idTemplate string) {
	messagetext.IdTemplate = idTemplate
}
