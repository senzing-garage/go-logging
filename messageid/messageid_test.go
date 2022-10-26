package messageid

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

func testError(test *testing.T, testObject MessageIdInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdTemplate - names begin with "Test"
// ----------------------------------------------------------------------------

// -- MessageId ---------------------------------------------------------------

func TestMessageId(test *testing.T) {
	testObject := &MessageIdTemplated{
		IdTemplate: "senzing-9999%04d",
	}
	actual, err := testObject.MessageId(1)
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestMessageIdWithDetails(test *testing.T) {
	testObject := &MessageIdTemplated{
		IdTemplate: "senzing-9999%04d",
	}
	actual, err := testObject.MessageId(1, "A", 1, testObject)
	testError(test, testObject, err)
	printActual(test, actual)
}
