/*
Package messagetext produces a string used in a "status" field of a log message.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageStatusInterface interface {
	MessageStatus(messageNumber int, details ...interface{}) (string, error)
}
