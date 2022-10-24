/*
The MessageTextDefault implementation maps the error number
to a format string.
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

/*
MessageTextDefault uses simple format string replacement
to produce a "text" string.
*/
type MessageTextDefault struct {

	// A map from message ids to format string.
	TextTemplates map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
 */
func (messagetext *MessageTextDefault) MessageText(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil

	result := ""
	textTemplate, ok := messagetext.TextTemplates[messageNumber]
	if ok {
		textRaw := fmt.Sprintf(textTemplate, details...)
		result = strings.Split(textRaw, "%!(")[0]
	}

	return result, err
}

/*
Set the map of message ids to format strings.
*/
func (messagetext *MessageTextDefault) SetTextTemplates(messages map[int]string) {
	messagetext.TextTemplates = messages
}
