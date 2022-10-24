/*
Package messagetext produces a string used in a "text" field of a log message.
*/
package messagetext

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageTextInterface interface {

	// Get the "text" value for a message id and its details.
	MessageText(messageNumber int, details ...interface{}) (string, error)

	// Set the map of message ids to format strings.
	SetTextTemplates(messages map[int]string)
}
