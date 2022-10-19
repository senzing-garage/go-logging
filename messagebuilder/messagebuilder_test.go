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

func TestBuildErrorWithDetails(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message", "A", 1)
	test.Log("Actual:", err)
}

// -- BuildMessage ------------------------------------------------------------

func TestBuildMessage(test *testing.T) {
	message := BuildMessage("unique-id-%04d", 5, "Message")
	test.Log("Actual:", message)
}

func TestBuildMessageWithDetails(test *testing.T) {
	message := BuildMessage("unique-id-%04d", 5, "Message", "A", 1)
	test.Log("Actual:", message)
}

// -- BuildMessageFromError ---------------------------------------------------

func TestBuildMessageFromError(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	test.Log("Actual:", err)

	message := BuildMessageFromError("unique-id-%04d", 5, "Message", err)
	test.Log("Actual:", message)
}

func TestBuildMessageFromErrorWithDetails(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	test.Log("Actual:", err)

	message := BuildMessageFromError("unique-id-%04d", 5, "Message", err, "A", 1)
	test.Log("Actual:", message)
}

// -- BuildMessageFromErrorUsingMap -------------------------------------------

func TestBuildMessageFromErrorUsingMap(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	test.Log("Actual:", err)

	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
		"TestClass":      test,
	}

	message := BuildMessageFromErrorUsingMap("unique-id-%04d", 5, "Message", err, detailsMap)
	test.Log("Actual:", message)
}

// -- BuildMessageUsingMap ----------------------------------------------------

func TestBuildMessageUsingMap(test *testing.T) {

	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
		"TestClass":      test,
	}

	message := BuildMessageUsingMap("unique-id-%04d", 5, "Message", detailsMap)
	test.Log("Actual:", message)
}

// -- BuildMessageId ----------------------------------------------------------

func TestBuildMessageId(test *testing.T) {

	message := BuildMessageId("unique-id-%04d", 5)
	test.Log("Actual:", message)
}

// -- BuildMessageLevel -------------------------------------------------------

func TestBuildMessageLevel(test *testing.T) {
	message := BuildMessageLevel(5, "Message")
	test.Log("Actual:", message)
}

// -- Miscellaneous -----------------------------------------------------------
