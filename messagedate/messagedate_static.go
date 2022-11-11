/*
The MessageDateStatic implementation returns a specified date.
Used mostly for repeatable test cases.
*/
package messagedate

import (
	"fmt"
	"time"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageDateStatic type is for returning a specific date.
type MessageDateStatic struct {
	Format    string
	Timestamp time.Time
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDate method returns string for a date value in the format YYYY-MM-DD in UTC.
func (messageDate *MessageDateStatic) MessageDate(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	if messageDate.Format == "" {
		messageDate.Format = "%04d-%02d-%02d"
	}
	return fmt.Sprintf(messageDate.Format, messageDate.Timestamp.UTC().Year(), messageDate.Timestamp.UTC().Month(), messageDate.Timestamp.UTC().Day()), nil
}
