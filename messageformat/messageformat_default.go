/*
The MessageFormatDefault implementation returns a simple string.
*/
package messageformat

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageFormatDefault type is for creating terse, default formatted messages.
type MessageFormatDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The Message method creates a terse, default formatted message.
func (messageFormat *MessageFormatDefault) Message(id string, status string, text string, details ...interface{}) (string, error) {
	var err error = nil

	result := ""
	if len(id) > 0 {
		result = result + fmt.Sprintf("%s: ", id)
	}
	if len(status) > 0 {
		result = result + fmt.Sprintf("(%s) ", status)
	}
	if len(text) > 0 {
		result = result + fmt.Sprintf("%s ", text)
	}
	if len(details) > 0 {
		detailMap := make(map[int]interface{})
		for index, value := range details {
			detailMap[index+1] = fmt.Sprintf("%#v", value)
		}
		result = result + fmt.Sprintf("[%v] ", detailMap)
	}

	return result, err
}
