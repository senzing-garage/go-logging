/*
The messageduration package produces a value for the "duration" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messageduration/messageduration_test.go
*/
package messageduration

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDurationInterface type defines methods for determining log level.
type MessageDurationInterface interface {
	MessageDuration(messageNumber int, details ...interface{}) (int64, error) // Get the "duration" value from the details.
}
