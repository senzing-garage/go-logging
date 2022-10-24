package messageloglevel

import (
	"errors"
	"fmt"
	"testing"
)

const printResults = 1

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func printResult(test *testing.T, title string, result interface{}) {
	if printResults == 1 {
		test.Logf("%s: %v", title, fmt.Sprintf("%v", result))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, testObject MessageLogLevelInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelSenzingApi - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestSenzingApiMessageLogLevel(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{}
	actual, err := testObject.MessageLogLevel(1, "This is message text")
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestSenzingApiMessageLogLevelWithErrors(test *testing.T) {
	anError1 := errors.New("0019E|Configuration not found")
	anError2 := errors.New("0099E|Made up error")

	testObject := &MessageLogLevelSenzingApi{}
	actual, err := testObject.MessageLogLevel(1, "A", 1, testObject, anError1, anError2)
	testError(test, testObject, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelInfo - names begin with "Test"
// ----------------------------------------------------------------------------

// -- MessageLogLevel ---------------------------------------------------------

func TestMessageLogLevel(test *testing.T) {
	testObject := &MessageLogLevelInfo{}
	actual, err := testObject.MessageLogLevel(1, "This is message text")
	testError(test, testObject, err)
	printActual(test, actual)
}
