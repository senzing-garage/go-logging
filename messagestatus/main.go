/*
The messagestatus package produces a value for the "status" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagestatus/messagestatus_test.go
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusInterface type defines methods for determining status.
type MessageStatusInterface interface {
	MessageStatus(messageNumber int, details ...interface{}) (string, error)
}

// The Status type is used to identify strings as being status strings in details parameter.
type Status string
