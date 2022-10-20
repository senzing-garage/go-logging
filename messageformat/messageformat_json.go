/*
Package message ...
*/
package messageformat

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// type JsonDetail struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

type MessageFormatJson struct {
	Id      string        `json:"id,omitempty"`
	Status  string        `json:"status,omitempty"`
	Text    interface{}   `json:"text,omitempty"`
	Details interface{}   `json:"details,omitempty"`
	Error   []interface{} `json:"error,omitempty"`
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

func (messageFormat *MessageFormatJson) BuildError(id string, status string, text string, details ...interface{}) error {
	return errors.New(messageFormat.BuildMessage(id, status, text, details...))
}

// Build a message given details as strings.
func (messageFormat *MessageFormatJson) BuildMessage(id string, status string, text string, details ...interface{}) string {

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
				messageFormat.Error = append(messageFormat.Error, priorError)
			default:
				valueAsString := stringify(typedValue)
				if isJson(valueAsString) {
					detailMap[strconv.Itoa(index+1)] = jsonAsInterface(valueAsString)
				} else {
					detailMap[strconv.Itoa(index+1)] = valueAsString
				}
			}
		}
		messageFormat.Details = detailMap
	}

	// Convert to JSON.

	result, _ := json.Marshal(messageFormat)
	return string(result)
}

// Build a message from an error
// func BuildMessageFromErrorXX(id string, level string, text string, err error, details ...interface{}) string {

// 	messageFormat := MessageFormatJson{}

// 	if len(id) > 0 {
// 		messageFormat.Id = id
// 	}

// 	if len(level) > 0 {
// 		messageFormat.Status = level
// 	}

// 	if len(text) > 0 {
// 		if isJson(text) {
// 			messageFormat.Text = jsonAsInterface(text)
// 		} else {
// 			messageFormat.Text = text
// 		}
// 	}

// 	if len(details) > 0 {
// 		detailMap := make(map[string]interface{})
// 		for index, value := range details {
// 			valueAsString := stringify(value)
// 			if isJson(valueAsString) {
// 				detailMap[strconv.Itoa(index+1)] = jsonAsInterface(valueAsString)
// 			} else {
// 				detailMap[strconv.Itoa(index+1)] = valueAsString
// 			}
// 		}
// 		messageFormat.Details = detailMap
// 	}

// 	// Nest prior Error message.

// 	if err != nil {
// 		errorMessage := err.Error()

// 		var priorError interface{}
// 		if isJson(errorMessage) {
// 			priorError = jsonAsInterface(errorMessage)
// 		} else {
// 			priorError = MessageFormatJson{
// 				Text: errorMessage,
// 			}
// 		}
// 		messageFormat.Error = priorError
// 	}

// 	// Convert to JSON.

// 	result, _ := json.Marshal(messageFormat)
// 	return string(result)
// }

// Build a message given details as a map of strings.
// func BuildMessageFromErrorUsingMapXX(id string, level string, text string, err error, details map[string]interface{}) string {

// 	messageFormat := MessageFormatJson{}

// 	// Fill optional fields.

// 	if len(id) > 0 {
// 		messageFormat.Id = id
// 	}

// 	if len(level) > 0 {
// 		messageFormat.Status = level
// 	}

// 	if len(text) > 0 {
// 		if isJson(text) {
// 			messageFormat.Text = jsonAsInterface(text)
// 		} else {
// 			messageFormat.Text = text
// 		}
// 	}

// 	if len(details) > 0 {
// 		detailMap := make(map[string]interface{})
// 		for index, value := range details {
// 			valueAsString := stringify(value)
// 			if isJson(valueAsString) {
// 				detailMap[index] = jsonAsInterface(valueAsString)
// 			} else {
// 				detailMap[index] = valueAsString
// 			}
// 		}
// 		messageFormat.Details = detailMap
// 	}

// 	// Nest prior Error message.

// 	if err != nil {
// 		errorMessage := err.Error()
// 		var priorError interface{}
// 		if isJson(errorMessage) {
// 			priorError = jsonAsInterface(errorMessage)
// 		} else {
// 			priorError = MessageFormatJson{
// 				Text: errorMessage,
// 			}
// 		}
// 		messageFormat.Error = priorError
// 	}

// 	// Convert to JSON.

// 	result, _ := json.Marshal(messageFormat)
// 	return string(result)
// }

// Build a message given details as a map of strings.
// func (messageFormat *MessageFormatJson) BuildMessageUsingMap(id string, level string, text string, details map[string]interface{}) string {

// 	// Fill optional fields.

// 	if len(id) > 0 {
// 		messageFormat.Id = id
// 	}

// 	if len(level) > 0 {
// 		messageFormat.Status = level
// 	}

// 	if len(text) > 0 {
// 		if isJson(text) {
// 			messageFormat.Text = jsonAsInterface(text)
// 		} else {
// 			messageFormat.Text = text
// 		}
// 	}

// 	if len(details) > 0 {
// 		detailMap := make(map[string]interface{})
// 		for index, value := range details {
// 			switch typedValue := value.(type) {
// 			case error:
// 				errorMessage := typedValue.Error()
// 				var priorError interface{}
// 				if isJson(errorMessage) {
// 					priorError = jsonAsInterface(errorMessage)
// 				} else {
// 					priorError = MessageFormatJson{
// 						Text: errorMessage,
// 					}
// 				}
// 				messageFormat.Error = append(messageFormat.Error, priorError)
// 			default:
// 				valueAsString := stringify(typedValue)
// 				if isJson(valueAsString) {
// 					detailMap[index] = jsonAsInterface(valueAsString)
// 				} else {
// 					detailMap[index] = valueAsString
// 				}
// 			}
// 		}
// 		messageFormat.Details = detailMap
// 	}

// 	// Convert to JSON.

// 	result, _ := json.Marshal(messageFormat)
// 	return string(result)
// }

// Parse JSON message.
// func ParseMessage(jsonString string) MessageFormatJson {
// 	var message MessageFormatJson
// 	json.Unmarshal([]byte(jsonString), &message)
// 	return message
// }
