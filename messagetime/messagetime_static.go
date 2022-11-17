/*
The MessageTimeStatic implementation returns a specified date.
Used mostly for repeatable test cases.
*/
package messagetime

import (
	"fmt"
	"time"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageTimeStatic type is for returning a specific date.
type MessageTimeStatic struct {
	Format    string    // A golang format string for formatting time. Default: "%02d:%02d:%02d.%09d"
	Timestamp time.Time // User specified time.
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageDate method returns string for a date value in the format YYYY-MM-DD in UTC.
func (messageDate *MessageTimeStatic) MessageTime(messageNumber int, messageTimestamp time.Time, details ...interface{}) (string, error) {
	if messageDate.Format == "" {
		messageDate.Format = "%02d:%02d:%02d.%09d"
	}
	return fmt.Sprintf(messageDate.Format, messageDate.Timestamp.UTC().Hour(), messageDate.Timestamp.UTC().Minute(), messageDate.Timestamp.Second(), messageDate.Timestamp.Nanosecond()), nil
}
