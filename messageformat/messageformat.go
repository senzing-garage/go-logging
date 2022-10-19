/*
Package message ...
*/
package messageformat

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

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
func BuildMessage(id string, level string, text string, details ...interface{}) string {

	resultStruct := MessageFormat{}

	if len(id) > 0 {
		resultStruct.Id = id
	}

	if len(level) > 0 {
		resultStruct.Level = level
	}

	if len(text) > 0 {
		if isJson(text) {
			resultStruct.Text = jsonAsInterface(text)
		} else {
			resultStruct.Text = text
		}
	}

	if len(details) > 0 {
		detailMap := make(map[string]interface{})
		for index, value := range details {
			valueAsString := stringify(value)
			if isJson(valueAsString) {
				detailMap[strconv.Itoa(index+1)] = jsonAsInterface(valueAsString)
			} else {
				detailMap[strconv.Itoa(index+1)] = valueAsString
			}
		}
		resultStruct.Details = detailMap
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Build a message from an error
func BuildMessageFromError(id string, level string, text string, err error, details ...interface{}) string {

	resultStruct := MessageFormat{}

	if len(id) > 0 {
		resultStruct.Id = id
	}

	if len(level) > 0 {
		resultStruct.Level = level
	}

	if len(text) > 0 {
		if isJson(text) {
			resultStruct.Text = jsonAsInterface(text)
		} else {
			resultStruct.Text = text
		}
	}

	if len(details) > 0 {
		detailMap := make(map[string]interface{})
		for index, value := range details {
			valueAsString := stringify(value)
			if isJson(valueAsString) {
				detailMap[strconv.Itoa(index+1)] = jsonAsInterface(valueAsString)
			} else {
				detailMap[strconv.Itoa(index+1)] = valueAsString
			}
		}
		resultStruct.Details = detailMap
	}

	// Nest prior Error message.

	if err != nil {
		errorMessage := err.Error()

		var priorError interface{}
		if isJson(errorMessage) {
			priorError = jsonAsInterface(errorMessage)
		} else {
			priorError = MessageFormat{
				Text: errorMessage,
			}
		}
		resultStruct.Error = priorError
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Build a message given details as a map of strings.
func BuildMessageFromErrorUsingMap(id string, level string, text string, err error, details map[string]interface{}) string {

	resultStruct := MessageFormat{}

	// Fill optional fields.

	if len(id) > 0 {
		resultStruct.Id = id
	}

	if len(level) > 0 {
		resultStruct.Level = level
	}

	if len(text) > 0 {
		if isJson(text) {
			resultStruct.Text = jsonAsInterface(text)
		} else {
			resultStruct.Text = text
		}
	}

	if len(details) > 0 {
		detailMap := make(map[string]interface{})
		for index, value := range details {
			valueAsString := stringify(value)
			if isJson(valueAsString) {
				detailMap[index] = jsonAsInterface(valueAsString)
			} else {
				detailMap[index] = valueAsString
			}
		}
		resultStruct.Details = detailMap
	}

	// Nest prior Error message.

	if err != nil {
		errorMessage := err.Error()
		var priorError interface{}
		if isJson(errorMessage) {
			priorError = jsonAsInterface(errorMessage)
		} else {
			priorError = MessageFormat{
				Text: errorMessage,
			}
		}
		resultStruct.Error = priorError
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Build a message given details as a map of strings.
func BuildMessageUsingMap(id string, level string, text string, details map[string]interface{}) string {

	resultStruct := MessageFormat{}

	// Fill optional fields.

	if len(id) > 0 {
		resultStruct.Id = id
	}

	if len(level) > 0 {
		resultStruct.Level = level
	}

	if len(text) > 0 {
		if isJson(text) {
			resultStruct.Text = jsonAsInterface(text)
		} else {
			resultStruct.Text = text
		}
	}

	if len(details) > 0 {
		detailMap := make(map[string]interface{})
		for index, value := range details {
			valueAsString := stringify(value)
			if isJson(valueAsString) {
				detailMap[index] = jsonAsInterface(valueAsString)
			} else {
				detailMap[index] = valueAsString
			}
		}
		resultStruct.Details = detailMap
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Parse JSON message.
func ParseMessage(jsonString string) MessageFormat {
	var message MessageFormat
	json.Unmarshal([]byte(jsonString), &message)
	return message
}
