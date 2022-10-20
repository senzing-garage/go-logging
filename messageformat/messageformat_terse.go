/*
Package message ...
*/
package messageformat

import (
	"errors"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageFormatTerse struct {
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (messageFormat *MessageFormatTerse) BuildError(id string, status string, text string, details ...interface{}) error {
	return errors.New(messageFormat.BuildMessage(id, status, text, details...))
}

func (messageFormat *MessageFormatTerse) BuildMessage(id string, status string, text string, details ...interface{}) string {
	return fmt.Sprintf("%s: %s - %s; %v", id, status, text, details)
}
