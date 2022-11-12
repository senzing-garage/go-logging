/*
The MessageErrorsSenzing implementation returns an empty value.
*/
package messageerrors

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsSenzing type is for returning an empty value.
type MessageErrorsSenzing struct{}

type messageErrorsSenzing struct {
	Text interface{} `json:"text,omitempty"`
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageErrors method returns an empty value.
func (messageErrors *MessageErrorsSenzing) MessageErrors(messageNumber int, details ...interface{}) (interface{}, error) {
	var err error = nil
	var result []interface{} = nil

	// Work with details.

	// Process different types of details.

	for _, value := range details {
		switch typedValue := value.(type) {

		case error:
			errorMessage := typedValue.Error()
			var priorError interface{}
			if isJson(errorMessage) {
				priorError = jsonAsInterface(errorMessage)
			} else {
				priorError = &messageErrorsSenzing{
					Text: errorMessage,
				}
			}
			result = append(result, priorError)
		}
	}

	if len(result) == 0 {
		result = nil
	}

	return result, err
}
