/*
The messagestatus package produces a string used in a "status" field of a log message.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusInterface type defines methods for determining status.
type MessageStatusInterface interface {
	MessageStatus(messageNumber int, details ...interface{}) (string, error)
}
