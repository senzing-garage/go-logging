/*
The MessageFormatDefault implementation returns a simple string.
*/
package messageformat

import (
	"fmt"
	"reflect"
	"strings"
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
func (messageFormat *MessageFormatDefault) Message(date string, time string, level string, location string, id string, status string, text string, duration int64, errors interface{}, details interface{}) (string, error) {
	var err error = nil

	result := ""

	if len(level) > 0 {
		result = result + fmt.Sprintf("%s ", level)
	}

	if len(id) > 0 {
		result = result + fmt.Sprintf("%s: ", id)
	}
	if len(status) > 0 {
		result = result + fmt.Sprintf("(%s) ", status)
	}
	if len(text) > 0 {
		result = result + fmt.Sprintf("%s ", text)
	}

	if errors != nil {
		if !reflect.ValueOf(errors).IsNil() {
			result = result + fmt.Sprintf("%#v ", errors)
		}
	}

	if details != nil {
		if !reflect.ValueOf(details).IsNil() {
			result = result + fmt.Sprintf("%v ", details)
		}
	}

	result = strings.TrimSpace(result)

	return result, err
}
