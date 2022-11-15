/*
The messagetime package produces a time string.
*/
package messagetime

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeInterface type defines methods for determining the time value.
type MessageTimeInterface interface {
	MessageTime(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error)
}
