/*
For more information, see [test].

[test]: messagetext/messagetext_test.go
*/
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

// -- MessageText -------------------------------------------------------------

func TestMessageText(test *testing.T) {
	testObject := &MessageTextTemplated{
		IdMessages: map[int]string{
			1: "Bob's middle initial is \"%s\" and his favorite number is %d.",
			2: "Sally got an \"%s\" on the paper.",
		},
	}
	actual, err := testObject.MessageText(1, "A", 1)
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestMessageTextWithMessageNumber(test *testing.T) {
	testObject := &MessageTextTemplated{
		IdMessages: map[int]string{
			1: "Bob's middle initial is \"%s\" and his favorite number is %d.",
			2: "Sally got an \"%s\" on the paper.",
		},
	}
	actual, err := testObject.MessageText(1, "A", 1, MsgNumber(2))
	testError(test, testObject, err)
	printActual(test, actual)
}
