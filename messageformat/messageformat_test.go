package messageformat

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

func testError(test *testing.T, testObject MessageFormatInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatJson - names begin with "Test"
// ----------------------------------------------------------------------------

// -- Message -----------------------------------------------------------------

func TestJsonMessage(test *testing.T) {
	testObject := &MessageFormatJson{}
	actual, err := testObject.Message("id-2", "try-again", "text-2", 123, "bob")
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestJsonMessageNoId(test *testing.T) {
	testObject := &MessageFormatJson{}
	actual, err := testObject.Message("", "try-again", "text-3", 123, "bob")
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestJsonMessageNoStatus(test *testing.T) {
	testObject := &MessageFormatJson{}
	actual, err := testObject.Message("id-4", "", "text-4", 123, "bob")
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestJsonMessageNoText(test *testing.T) {
	testObject := &MessageFormatJson{}
	actual, err := testObject.Message("id-5", "try-again", "", 123, "bob")
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestJsonMessageNoDetails(test *testing.T) {
	testObject := &MessageFormatJson{}
	actual, err := testObject.Message("id-6", "try-again", "text-6")
	testError(test, testObject, err)
	printActual(test, actual)
}

func TestJsonMessageNothing(test *testing.T) {
	testObject := &MessageFormatJson{}
	actual, err := testObject.Message("", "", "")
	testError(test, testObject, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatTerse - names begin with "Test"
// ----------------------------------------------------------------------------

// -- Message -----------------------------------------------------------------

func TestTerseMessage(test *testing.T) {
	testObject := &MessageFormatTerse{}
	actual, err := testObject.Message("id-1", "try-again", "text-1", 123, "bob")
	testError(test, testObject, err)
	printActual(test, actual)
}
