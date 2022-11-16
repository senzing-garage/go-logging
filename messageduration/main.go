/*
The messageduration package produces a value for the "duration" field.
*/
package messageduration

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDurationInterface type defines methods for determining log level.
type MessageDurationInterface interface {
	MessageDuration(messageNumber int, details ...interface{}) (int64, error)
}
