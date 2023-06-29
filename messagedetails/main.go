/*
The messagedetails package produces a value for the "details" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagedetails/messagedetails_test.go
*/
package messagedetails

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDetailsInterface type defines methods for determining the details value.
type MessageDetailsInterface interface {
	MessageDetails(messageNumber int, details ...interface{}) (interface{}, error) // Get the "details" value from the messageNumber and details.
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func isJson(unknownString string) bool {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownStringUnescaped), &jsonString) == nil
}

func jsonAsInterface(unknownString string) interface{} {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	err = json.Unmarshal([]byte(unknownStringUnescaped), &jsonString)
	if err != nil {
		panic(err)
	}
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
