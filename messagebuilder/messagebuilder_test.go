package messagebuilder

import (
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
)

const MessageIdFormat = "senzing-9999%04d"

const (
	EnablePrinting = 1
	TruncateLength = 300
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func truncate(aString string) string {
	return truncator.Truncate(aString, TruncateLength, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if EnablePrinting == 1 {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result)))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestBuildError(test *testing.T) {
	actual := BuildError("unique-id-%04d", 5, "Error message")
	printActual(test, actual)
}

func TestBuildErrorWithDetails(test *testing.T) {
	actual := BuildError("unique-id-%04d", 5, "Error message", "A", 1)
	printActual(test, actual)
}

// -- BuildMessage ------------------------------------------------------------

func TestBuildMessage(test *testing.T) {
	actual := BuildMessage("unique-id-%04d", 5, "Message")
	printActual(test, actual)
}

func TestBuildMessageWithDetails(test *testing.T) {
	actual := BuildMessage("unique-id-%04d", 5, "Message", "A", 1)
	printActual(test, actual)
}

// -- BuildMessageFromError ---------------------------------------------------

func TestBuildMessageFromError(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	printResult(test, "Error", err)
	actual := BuildMessageFromError("unique-id-%04d", 5, "Message", err)
	printActual(test, actual)
}

func TestBuildMessageFromErrorWithDetails(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	printResult(test, "Error", err)
	actual := BuildMessageFromError("unique-id-%04d", 5, "Message", err, "A", 1)
	printActual(test, actual)
}

// -- BuildMessageFromErrorUsingMap -------------------------------------------

func TestBuildMessageFromErrorUsingMap(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	printResult(test, "Error", err)

	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
		"TestClass":      test,
	}

	actual := BuildMessageFromErrorUsingMap("unique-id-%04d", 5, "Message", err, detailsMap)
	printActual(test, actual)
}

// -- BuildMessageUsingMap ----------------------------------------------------

func TestBuildMessageUsingMap(test *testing.T) {

	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
		"TestClass":      test,
	}

	actual := BuildMessageUsingMap("unique-id-%04d", 5, "Message", detailsMap)
	printActual(test, actual)
}

// -- BuildMessageId ----------------------------------------------------------

func TestBuildMessageId(test *testing.T) {
	actual := BuildMessageId("unique-id-%04d", 5)
	printActual(test, actual)
}

// -- BuildMessageLevel -------------------------------------------------------

func TestBuildMessageLevel(test *testing.T) {
	actual := BuildMessageLevel(5, "Message")
	printActual(test, actual)
}

// -- Miscellaneous -----------------------------------------------------------
