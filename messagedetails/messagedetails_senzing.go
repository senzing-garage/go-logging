/*
The MessageDetailsSenzing implementation returns map[string]interface{}.
*/
package messagedetails

import (
	"fmt"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDetailsSenzing type is for returning a map[string]interface{}.
type MessageDetailsSenzing struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDetails method returns a map[string]interface{} with un-indexed instances receiving an ordinal index.
func (messageDetails *MessageDetailsSenzing) MessageDetails(messageNumber int, details ...interface{}) (interface{}, error) {
	var err error = nil

	result := make(map[string]interface{})

	// Process different types of details.

	for index, value := range details {
		switch typedValue := value.(type) {
		case nil:
			result[strconv.Itoa(index+1)] = "<nil>"

		case int, float64:
			result[strconv.Itoa(index+1)] = typedValue

		case string:
			if isJson(typedValue) {
				result[strconv.Itoa(index+1)] = jsonAsInterface(typedValue)
			} else {
				result[strconv.Itoa(index+1)] = typedValue
			}

		case bool:
			result[strconv.Itoa(index+1)] = fmt.Sprintf("%t", typedValue)

		case error:
			// do nothing.

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

	if len(result) == 0 {
		result = nil
	}

	return result, err
}
