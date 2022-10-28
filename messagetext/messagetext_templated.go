/*
The MessageTextTemplated implementation maps the message number to a format string.
The format string is populated with values submitted.
*/
package messagetext

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTextTemplated type uses format string replacement to produce a "text" string.
type MessageTextTemplated struct {

	// A map from message numbers to format string.
	TextTemplates map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The MessageText method chooses a format string based on the message number and populates it from the details.
To override the message number, submit a detail of type MessageNumber.
The MessageNumber value will be used to choose the template from MessageTextTemplated.TextTemplates.
*/
func (messageText *MessageTextTemplated) MessageText(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	result := ""

	// Determine if a message number was passed in via "details" parameter.

	if len(details) > 0 {
		for index := len(details) - 1; index >= 0; index-- {
			detail := details[index]
			switch typedDetail := detail.(type) {
			case MessageNumber:
				textTemplate, ok := messageText.TextTemplates[int(typedDetail)]
				if ok {
					textRaw := fmt.Sprintf(textTemplate, details...)
					result = strings.Split(textRaw, "%!(")[0]
					break
				}
			}
		}
	}

	// The normal case is that the message number is passed in as the "messageNumber" parameter.

	if result == "" {
		textTemplate, ok := messageText.TextTemplates[messageNumber]
		if ok {
			textRaw := fmt.Sprintf(textTemplate, details...)
			result = strings.Split(textRaw, "%!(")[0]
		}
	}

	return result, err
}
