/*
The messageformat package renders messages.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messageformat/messageformat_test.go
*/
package messageformat

import (
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageFormatInterface type defines methods for producing formatting messages.
type MessageFormatInterface interface {
	Message(date string, time string, level string, location string, id string, status string, text string, duration int64, errors interface{}, details interface{}) (string, error) // Create a message.
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
	json.Unmarshal([]byte(unknownStringUnescaped), &jsonString)
	return jsonString
}
