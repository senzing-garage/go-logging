// The message package formats messages into a JSON string.
package messageformat

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Detail struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MessageFormat struct {
	Id      string      `json:"id,omitempty"`
	Level   string      `json:"level,omitempty"`
	Text    interface{} `json:"text,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

var messageformat *MessageFormat
