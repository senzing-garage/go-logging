/*
The MessageDetailsNull implementation returns an empty value.
*/
package messagedetails

import (
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsNull type is for returning an empty value.
type MessageDetailsSenzing struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDetails method returns an empty value.
func (messageDetails *MessageDetailsSenzing) MessageErrors(messageNumber int, details ...interface{}) (interface{}, error) {
	var err error = nil

	result := make(map[string]interface{})

	// Process different types of details.

	for index, value := range details {
		switch typedValue := value.(type) {
		case nil:
			result[strconv.Itoa(index+1)] = "<nil>"

		case error:
			// do nothing

		case map[string]string:
			for mapIndex, mapValue := range typedValue {
				mapValueAsString := stringify(mapValue)
				if isJson(mapValueAsString) {
					result[mapIndex] = jsonAsInterface(mapValueAsString)
				} else {
					result[mapIndex] = mapValueAsString
				}
			}

		default:
			valueAsString := stringify(typedValue)
			if isJson(valueAsString) {
				result[strconv.Itoa(index+1)] = jsonAsInterface(valueAsString)
			} else {
				result[strconv.Itoa(index+1)] = valueAsString
			}
		}
	}

	return result, err
}
