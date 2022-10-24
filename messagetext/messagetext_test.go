package messagetext

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

func testError(test *testing.T, testObject MessageTextInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelSenzingApi - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestMessageText(test *testing.T) {
	testObject := &MessageTextDefault{
		TextTemplates: map[int]string{
			1: "Bob's middle initial is \"%s\" and his favorite number is %d.",
		},
	}
	actual, err := testObject.MessageText(1, "A", 1)
	testError(test, testObject, err)
	printActual(test, actual)
}
