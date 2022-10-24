package messagestatus

import (
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

func testError(test *testing.T, testObject MessageStatusInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelSenzingApi - names begin with "Test"
// ----------------------------------------------------------------------------

// -- MessageStatus -----------------------------------------------------------

func TestMessageStatus(test *testing.T) {
	testObject := &MessageStatusSenzingApi{}
	actual, err := testObject.MessageStatus(1)
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestMessageStatusWithDetails(test *testing.T) {
	testObject := &MessageStatusSenzingApi{}
	actual, err := testObject.MessageStatus(1, "A", 1, testObject)
	testError(test, testObject, err)
	printActual(test, actual)
}
