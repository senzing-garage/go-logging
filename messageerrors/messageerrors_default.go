/*
The MessageErrorsDefault implementation returns a []interface{} containing error representations.
*/
package messageerrors

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsDefault type is for returning a []interface{} containing error representations.
type MessageErrorsDefault struct{}

type messageErrorsDefault struct {
	Text interface{} `json:"text,omitempty"`
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageErrors method returns a []interface{} containing error representations.
func (messageErrors *MessageErrorsDefault) MessageErrors(messageNumber int, details ...interface{}) (interface{}, error) {
	var err error = nil
	var result []interface{} = nil

	for _, value := range details {
		switch typedValue := value.(type) {

		case error:
			errorMessage := typedValue.Error()
			var priorError interface{}
			if isJson(errorMessage) {
				priorError = &messageErrorsDefault{
					Text: jsonAsInterface(errorMessage),
				}
			} else {
				priorError = &messageErrorsDefault{
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
