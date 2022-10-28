/*
The MessageFormatJson implementation returns a message in the JSON format.
*/
package messageformat

import (
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageFormatJson type is for creating formatted messages in JSON.
type MessageFormatJson struct {
}

// Fields in the formatted message.
// Order is important.
// It should be id, status, text, errors, details.
type messageFormatJson struct {
	Id      string        `json:"id,omitempty"`
	Status  string        `json:"status,omitempty"`
	Text    interface{}   `json:"text,omitempty"`
	Errors  []interface{} `json:"errors,omitempty"`
	Details interface{}   `json:"details,omitempty"`
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The Message method creates a JSON formatted message.
func (messageFormat *MessageFormatJson) Message(id string, status string, text string, details ...interface{}) (string, error) {
	var err error = nil
	messageBuilder := &messageFormatJson{}

	// Set output Id, Status, and Text fields.

	if len(id) > 0 {
		messageBuilder.Id = id
	}

	if len(status) > 0 {
		messageBuilder.Status = status
	}

	if len(text) > 0 {
		if isJson(text) {
			messageBuilder.Text = jsonAsInterface(text)
		} else {
			messageBuilder.Text = text
		}
	}

	// Work with details.

	if len(details) > 0 {
		var errorsList []interface{}
		detailMap := make(map[string]interface{})

		// Process different types of details.

		for index, value := range details {
			switch typedValue := value.(type) {
			case error:
				errorMessage := typedValue.Error()
				var priorError interface{}
				if isJson(errorMessage) {
					priorError = jsonAsInterface(errorMessage)
				} else {
					priorError = &messageFormatJson{
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

		// Set output Errors and Details fields.

		messageBuilder.Errors = errorsList
		messageBuilder.Details = detailMap
	}

	// Convert to JSON.

	result, _ := json.Marshal(messageBuilder)
	return string(result), err
}
