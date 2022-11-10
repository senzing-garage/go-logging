/*
The messagedate package produces a date string.
*/
package messagedate

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDateInterface type defines methods for determining the date value.
type MessageDateInterface interface {
	MessageDate(messageNumber int, details ...interface{}) (string, error)
}
