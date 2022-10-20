// The message package formats messages into a JSON string.
package messageformat

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageFormatInterface interface {
	BuildError(id string, status string, text string, details ...interface{}) error
	BuildMessage(id string, status string, text string, details ...interface{}) string
}
