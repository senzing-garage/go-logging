package messagebuilder

import (
	"testing"
)

const MessageIdFormat = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestBuildError(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	test.Log("Actual:", err)
}

// -- BuildMessage ------------------------------------------------------------
// -- BuildMessageFromError ---------------------------------------------------
// -- BuildMessageFromErrorUsingMap -------------------------------------------
// -- BuildMessageUsingMap ----------------------------------------------------
// -- BuildMessageId ----------------------------------------------------------
// -- BuildMessageLevel -------------------------------------------------------
