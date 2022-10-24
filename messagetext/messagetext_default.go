/*
Package helper ...
*/
package messagetext

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageTextDefault struct {
	TextTemplates map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagetext *MessageTextDefault) MessageText(errorNumber int, details ...interface{}) (string, error) {
	var err error = nil

	result := ""
	textTemplate, ok := messagetext.TextTemplates[errorNumber]
	if ok {
		textRaw := fmt.Sprintf(textTemplate, details...)
		result = strings.Split(textRaw, "%!(")[0]
	}

	return result, err
}

// TODO:
func (messagetext *MessageTextDefault) SetTextTemplates(messages map[int]string) {
	messagetext.TextTemplates = messages
}
