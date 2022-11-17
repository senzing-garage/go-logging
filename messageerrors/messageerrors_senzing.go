/*
The MessageErrorsSenzing implementation returns a []interface{} containing error representations.
*/
package messageerrors

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsSenzing type is for returning a []interface{} containing error representations.
type MessageErrorsSenzing struct{}

type messageErrorsSenzing struct {
	Text interface{} `json:"text,omitempty"`
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageErrors method returns a []interface{} containing error representations.
func (messageErrors *MessageErrorsSenzing) MessageErrors(messageNumber int, details ...interface{}) (interface{}, error) {
	var err error = nil
	var result []interface{} = nil

	for _, value := range details {
		switch typedValue := value.(type) {

		case error:
			errorMessage := typedValue.Error()
			var priorError interface{}
			if isJson(errorMessage) {
				priorError = &messageErrorsSenzing{
					Text: jsonAsInterface(errorMessage),
				}
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
