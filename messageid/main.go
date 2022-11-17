/*
The messageid package produces a value for the "id" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messageid/messageid_test.go
*/
package messageid

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageIdInterface type defines methods for producing a message identifier.
type MessageIdInterface interface {
	MessageId(messageNumber int, details ...interface{}) (string, error)
}
