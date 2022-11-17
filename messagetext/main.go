/*
The messagetext package produces a value for the "text" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagetext/messagetext_test.go
*/
package messagetext

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageNumber type is used to identify the integer is a "message number" in detail parameters.
type MessageNumber int

// The MessageTextInterface type defines methods for creating message text.
type MessageTextInterface interface {
	MessageText(messageNumber int, details ...interface{}) (string, error) // Get the "text" value for a message id and its details.
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// The MsgNumber function casts an integer to a message number.
func MsgNumber(messageNumber int) MessageNumber {
	return MessageNumber(messageNumber)
}
