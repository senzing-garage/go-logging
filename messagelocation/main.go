/*
The messagelocation package produces a value for the "location" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagelocation/messagelocation_test.go
*/
package messagelocation

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The CallerSkip type is used to identify the integer is the detail parameters.
// Number of stacks to ascend. See https://pkg.go.dev/runtime#Caller
type CallerSkip int

// The MessageLogLevelInterface type defines methods for producing the value of the location field.
type MessageLocationInterface interface {
	MessageLocation(messageNumber int, details ...interface{}) (string, error) // Get the "location" value from the messageNumber and details.
}
