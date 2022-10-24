/*
Package messagetext produces a string used in a "text"
field of a log message.
*/
package messagetext

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageTextInterface interface {
	MessageText(messageNumber int, details ...interface{}) (string, error)
	SetTextTemplates(messages map[int]string)
}
