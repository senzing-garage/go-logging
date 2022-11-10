/*
The messagetime package produces a time string.
*/
package messagetime

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeInterface type defines methods for determining the time value.
type MessageTimeInterface interface {
	MessageTime(messageNumber int, details ...interface{}) (string, error)
}
