/*
Package messagestatus produces a string used in a "status" field of a log message.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusInterface interface {
	MessageStatus(messageNumber int, details ...interface{}) (string, error)
}
