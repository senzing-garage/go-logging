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
func (mesageId *MessageIdDefault) GetMessageId(errorNumber int) (string, error) {
	var err error = nil
	result := fmt.Sprintf(mesageId.IdTemplate, errorNumber)
	return result, err
}
