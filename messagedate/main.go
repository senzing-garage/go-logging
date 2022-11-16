/*
The messagedate package produces a value for the "date" field.
*/
package messagedate

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDateInterface type defines methods for determining the date value.
type MessageDateInterface interface {
	MessageDate(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error)
}
