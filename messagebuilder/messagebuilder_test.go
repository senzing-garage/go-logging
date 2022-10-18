package messagebuilder

import (
	"errors"
	"testing"
)

const MessageIdFormat = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestBuildError(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	test.Log("Actual:", err)
}

func TestLogMessage(test *testing.T) {
	LogMessage(MessageIdFormat, 2000, "Test message", "Variable1", "Variable2")
}

func TestLogMessageFromError(test *testing.T) {
	anError := errors.New("This is a new error")
	LogMessageFromError(MessageIdFormat, 2002, "Test message", anError, "Variable1", "Variable2")
}
