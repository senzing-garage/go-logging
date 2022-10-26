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

type MessageFormatDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

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
