/*
Package message ...
*/
package messageformat

import (
	"errors"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageFormatTerse struct {
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (messageFormat *MessageFormatTerse) BuildError(id string, status string, text string, details ...interface{}) error {
	return errors.New(messageFormat.BuildMessage(id, status, text, details...))
}

func (messageFormat *MessageFormatTerse) BuildMessage(id string, status string, text string, details ...interface{}) string {
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
			detailMap[index+1] = value
		}
		result = result + fmt.Sprintf("%v ", detailMap)
	}

	return result
}
