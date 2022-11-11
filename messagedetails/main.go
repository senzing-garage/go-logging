/*
The messagedetails package produces a date string.
*/
package messagedetails

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDetailsInterface type defines methods for determining the date value.
type MessageDetailsInterface interface {
	MessageDetails(messageNumber int, details ...interface{}) ([]interface{}, error)
}
