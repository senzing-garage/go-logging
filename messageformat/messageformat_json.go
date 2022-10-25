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

type MessageFormatJson struct {
	Details interface{}   `json:"details,omitempty"`
	Errors  []interface{} `json:"errors,omitempty"`
	Id      string        `json:"id,omitempty"`
	Status  string        `json:"status,omitempty"`
	Text    interface{}   `json:"text,omitempty"`
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build a message given details as strings.
func (messageFormat *MessageFormatJson) Message(id string, status string, text string, details ...interface{}) (string, error) {

	// Because the structure could be shared, thread safty needs to be implemented.

	lock.Lock()
	defer lock.Unlock()

	var err error = nil

	// Clear old values.

	messageFormat.Details = nil
	messageFormat.Errors = nil
	messageFormat.Id = ""
	messageFormat.Status = ""
	messageFormat.Text = nil

	// Set new values.

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
