/*
Package message ...
*/
package messageformat

import (
	"encoding/json"
	// "errors"
	"fmt"
	"reflect"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageFormatJson struct {
	Id      string        `json:"id,omitempty"`
	Status  string        `json:"status,omitempty"`
	Text    interface{}   `json:"text,omitempty"`
	Details interface{}   `json:"details,omitempty"`
	Errors  []interface{} `json:"errors,omitempty"`
}

// ----------------------------------------------------------------------------
// Internal methods
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
		xType := reflect.TypeOf(unknown)
		xValue := reflect.ValueOf(unknown)
		result = fmt.Sprintf("(%s)%+v", xType, xValue)
	}
	return result
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build a message given details as strings.
func (messageFormat *MessageFormatJson) Message(id string, status string, text string, details ...interface{}) (string, error) {

	var err error = nil

	if len(id) > 0 {
		messageFormat.Id = id
	}

	if len(status) > 0 {
		messageFormat.Status = status
	}

	if len(text) > 0 {
		if isJson(text) {
			messageFormat.Text = jsonAsInterface(text)
		} else {
			messageFormat.Text = text
		}
	}

	if len(details) > 0 {
		var errorsList []interface{}
		detailMap := make(map[string]interface{})
		for index, value := range details {
			switch typedValue := value.(type) {
			case error:
				errorMessage := typedValue.Error()
				var priorError interface{}
				if isJson(errorMessage) {
					priorError = jsonAsInterface(errorMessage)
				} else {
					priorError = MessageFormatJson{
						Text: errorMessage,
					}
				}
				errorsList = append(errorsList, priorError)

			case map[string]string:
				for mapIndex, mapValue := range typedValue {
					mapValueAsString := stringify(mapValue)
					if isJson(mapValueAsString) {
						detailMap[mapIndex] = jsonAsInterface(mapValueAsString)
					} else {
						detailMap[mapIndex] = mapValueAsString
					}
				}

			default:
				valueAsString := stringify(typedValue)
				if isJson(valueAsString) {
					detailMap[strconv.Itoa(index+1)] = jsonAsInterface(valueAsString)
				} else {
					detailMap[strconv.Itoa(index+1)] = valueAsString
				}
			}
		}
		messageFormat.Errors = errorsList
		messageFormat.Details = detailMap
	}

	// Convert to JSON.

	result, _ := json.Marshal(messageFormat)
	return string(result), err
}

// Parse JSON message.
// func ParseMessage(jsonString string) MessageFormatJson {
// 	var message MessageFormatJson
// 	json.Unmarshal([]byte(jsonString), &message)
// 	return message
// }
