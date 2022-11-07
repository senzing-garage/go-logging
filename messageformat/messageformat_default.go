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
		for index, unknown := range details {
			switch value := unknown.(type) {
			case string:
				detailMap[index+1] = value
			case int:
				detailMap[index+1] = value
			case float64:
				detailMap[index+1] = value
			case bool:
				detailMap[index+1] = fmt.Sprintf("%t", value)
			case error:
				detailMap[index+1] = value.Error()
			case nil:
				detailMap[index+1] = "<nil>"
			default:
				xType := reflect.TypeOf(unknown)
				xValue := reflect.ValueOf(unknown)
				detailMap[index+1] = fmt.Sprintf("(%s)%#v", xType, xValue)
			}
		}
		result = result + fmt.Sprintf("[%v] ", detailMap)
	}
	result = strings.TrimSpace(result)

	return result, err
}
