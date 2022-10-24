// The message package formats messages into a JSON string.
package messageformat

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageFormatInterface interface {
	Message(id string, status string, text string, details ...interface{}) (string, error)
}
