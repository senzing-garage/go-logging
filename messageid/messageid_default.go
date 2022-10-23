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
func (messageId *MessageIdDefault) MessageId(errorNumber int) (string, error) {
	var err error = nil
	idTemplate := "%04d"
	if len(messageId.IdTemplate) > 0 {
		idTemplate = messageId.IdTemplate
	}
	result := fmt.Sprintf(idTemplate, errorNumber)
	return result, err
}

// TODO:
func (messagetext *MessageIdDefault) SetMessageIdTemplate(idTemplate string) {
	messagetext.IdTemplate = idTemplate
}
