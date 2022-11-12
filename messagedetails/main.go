/*
The messagedetails package produces a date string.
*/
package messagedetails

import (
	"encoding/json"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDetailsInterface type defines methods for determining the date value.
type MessageDetailsInterface interface {
	MessageDetails(messageNumber int, details ...interface{}) (interface{}, error)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func isJson(unknownString string) bool {
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownString), &jsonString) == nil
}

func jsonAsInterface(unknownString string) interface{} {
	var jsonString json.RawMessage
	json.Unmarshal([]byte(unknownString), &jsonString)
	return jsonString
}

func stringify(unknown interface{}) string {
	// See https://pkg.go.dev/fmt for format strings.
	var result string
	switch value := unknown.(type) {
	case nil:
		result = "<nil>"
	case string:
		result = value
	case int:
		result = fmt.Sprintf("%d", value)
	case float64:
		result = fmt.Sprintf("%g", value)
	case bool:
		result = fmt.Sprintf("%t", value)
	case error:
		result = value.Error()
	default:
		// xType := reflect.TypeOf(unknown)
		// xValue := reflect.ValueOf(unknown)
		// result = fmt.Sprintf("(%s)%#v", xType, xValue)
		result = fmt.Sprintf("%#v", unknown)
	}
	return result
}