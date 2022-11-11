/*
The messageerrors package produces a date string.
*/
package messageerrors

import "encoding/json"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsInterface type defines methods for aggregating errors from details.
type MessageErrorsInterface interface {
	MessageErrors(messageNumber int, details ...interface{}) ([]interface{}, error)
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
