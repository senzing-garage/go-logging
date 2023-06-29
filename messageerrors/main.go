/*
The messageerrors package produces a value for the "errors" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messageerrors/messageerrors_test.go
*/
package messageerrors

import (
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsInterface type defines methods for aggregating errors from details.
type MessageErrorsInterface interface {
	MessageErrors(messageNumber int, details ...interface{}) (interface{}, error) // Get the "errors" value from the details.
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
