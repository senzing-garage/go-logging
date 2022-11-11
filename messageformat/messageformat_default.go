/*
The MessageFormatDefault implementation returns a simple string.
*/
package messageformat

import (
	"fmt"
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
	if len(id) > 0 {
		result = result + fmt.Sprintf("%s: ", id)
	}
	if len(status) > 0 {
		result = result + fmt.Sprintf("(%s) ", status)
	}
	if len(text) > 0 {
		result = result + fmt.Sprintf("%s ", text)
	}

	if details != nil {
		result = result + fmt.Sprintf("%v ", details)
	}

	// detailMap := make(map[int]interface{})
	// for index, unknown := range details {
	// 	switch value := unknown.(type) {
	// 	case nil:
	// 		detailMap[index+1] = "<nil>"
	// 	case string, int, float64:
	// 		detailMap[index+1] = value
	// 	case bool:
	// 		detailMap[index+1] = fmt.Sprintf("%t", value)
	// 	case error:
	// 		detailMap[index+1] = value.Error()
	// 	default:
	// 		// xType := reflect.TypeOf(unknown)
	// 		// xValue := reflect.ValueOf(unknown)
	// 		// detailMap[index+1] = fmt.Sprintf("(%s)%#v", xType, xValue)
	// 		detailMap[index+1] = fmt.Sprintf("%#v", unknown)
	// 	}
	// }
	// result = result + fmt.Sprintf("[%v] ", detailMap)
	result = strings.TrimSpace(result)

	return result, err
}
