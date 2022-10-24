/*
Package helper ...
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageStatusNull) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""
	return result, err
}
