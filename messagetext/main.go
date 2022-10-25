/*
Package messagetext produces a string used in a "text" field of a log message.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagetext/messagetext_test.go
*/
package messagetext

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageNumber int

type MessageTextInterface interface {

	// Get the "text" value for a message id and its details.
	MessageText(messageNumber int, details ...interface{}) (string, error)

	// Set the map of message ids to format strings.
	SetTextTemplates(messages map[int]string)
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// Cast an integer to a message number.
func MsgNumber(messageNumber int) MessageNumber {
	return MessageNumber(messageNumber)
}
