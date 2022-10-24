package messagestatus

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

func TestMessageStatusWithSenzingApiError(test *testing.T) {
	anError := errors.New("0099E|Configuration not found")
	testObject := &MessageStatusSenzingApi{}
	actual, err := testObject.MessageStatus(1, "A", 1, testObject, anError)
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestMessageStatusWith2SenzingApiError2(test *testing.T) {
	anError1 := errors.New("0019E|Configuration not found")
	anError2 := errors.New("0099E|Made up error")

	testObject := &MessageStatusSenzingApi{}
	actual, err := testObject.MessageStatus(1, "A", 1, testObject, anError1, anError2)
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestMessageStatusWithUnknownError(test *testing.T) {
	anError := errors.New("1234E|Made up error")

	testObject := &MessageStatusSenzingApi{}
	actual, err := testObject.MessageStatus(1000, "A", 1, testObject, anError)
	testError(test, testObject, err)
	printActual(test, actual)
}
