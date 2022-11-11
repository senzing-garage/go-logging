/*
The MessageDateDefault implementation returns a date in the format YYYY-MM-DD in UTC.
*/
package messagedate

import (
	"fmt"
	"time"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDateDefault type is for returning a date in the format YYYY-MM-DD in UTC.
type MessageDateDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDate method returns string for a date value in the format YYYY-MM-DD in UTC.
func (messageDate *MessageDateDefault) MessageDate(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	return fmt.Sprintf("%04d-%02d-%02d", messageTimestamp.UTC().Year(), messageTimestamp.UTC().Month(), messageTimestamp.UTC().Day()), nil
}
