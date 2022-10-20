// The logger package is a set of method to help with common tasks.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagebuilder

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageBuilderInterface interface {
	BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error
	BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string
	BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string
	BuildMessageId(idTemplate string, errorNumber int) string
	BuildMessageLevel(errorNumber int, message string) string
}
