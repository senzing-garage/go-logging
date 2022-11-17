/*
The messagetime package produces a value for the "time" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagetime/messagetime_test.go
*/
package messagetime

import "time"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeInterface type defines methods for determining the time value.
type MessageTimeInterface interface {
	MessageTime(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) // Get the "time" value from the id, messageTimestamp, and details.
}
