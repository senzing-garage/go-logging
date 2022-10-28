/*
The messageid package produces customized message identifiers.
*/
package messageid

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageIdInterface type defines methods for producing a message identifier.
type MessageIdInterface interface {
	MessageId(messageNumber int, details ...interface{}) (string, error)
}
