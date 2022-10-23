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
	Messages map[int]string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagetext *MessageTextDefault) MessageText(errorNumber int, details ...interface{}) (string, error) {
	var err error = nil

	result := ""
	textTemplate, ok := messagetext.Messages[errorNumber]
	if ok {
		textRaw := fmt.Sprintf(textTemplate, details...)
		result = strings.Split(textRaw, "%!(")[0]
	}

	return result, err
}
