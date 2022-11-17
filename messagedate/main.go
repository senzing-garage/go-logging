/*
The messagedate package produces a value for the "date" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagedate/messagedate_test.go
*/
package messagedate

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDateInterface type defines methods for determining the date value.
type MessageDateInterface interface {
	MessageDate(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) // Get the "date" value from the id, messageTimestamp, and details.
}
