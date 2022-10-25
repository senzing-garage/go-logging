/*
The MessageTextDefault implementation maps the message number to a format string.
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
MessageTextDefault uses simple format string replacement to produce a "text" string.
*/
type MessageTextDefault struct {

	// A map from message numbers to format string. See SetTextTemplates.
	TextTemplates map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// MessageText gets the "text" value given the message number and it's details.
func (messagetext *MessageTextDefault) MessageText(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	result := ""

	if len(details) > 0 {
		for index := len(details) - 1; index >= 0; index-- {
			detail := details[index]
			switch typedDetail := detail.(type) {
			case MessageNumber:
				textTemplate, ok := messagetext.TextTemplates[int(typedDetail)]
				if ok {
					textRaw := fmt.Sprintf(textTemplate, details...)
					result = strings.Split(textRaw, "%!(")[0]
					break
				}
			}
		}
	}

	if result == "" {
		textTemplate, ok := messagetext.TextTemplates[messageNumber]
		if ok {
			textRaw := fmt.Sprintf(textTemplate, details...)
			result = strings.Split(textRaw, "%!(")[0]
		}
	}

	return result, err
}

/*
SetTextTemplates sets the map of message numbers to format strings.
Example map:

	var textTemplates = map[int]string{
		5:    "The favorite number for %s is %d",
		999:  "A test of INFO",
		1000: "A test of WARN",
		2000: "A test of ERROR",
	}
	messagelogger.GetMessageLogger().SetTextTemplates(textTemplates)
*/
func (messagetext *MessageTextDefault) SetTextTemplates(messages map[int]string) {
	messagetext.TextTemplates = messages
}
