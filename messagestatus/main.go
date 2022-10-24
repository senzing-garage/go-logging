/*
Package messagestatus provides "status" values.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageStatusInterface interface {
	MessageStatus(messageNumber int, details ...interface{}) (string, error)
}
