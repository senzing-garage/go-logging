/*
The messagelocation package produces a value for the location string.
*/
package messagelocation

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type CallerSkip int

// The MessageLogLevelInterface type defines methods for producing the value of the location field.
type MessageLocationInterface interface {
	MessageLocation(messageNumber int, details ...interface{}) (string, error)
}
