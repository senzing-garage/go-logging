package messageformat

import (
	"errors"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/stretchr/testify/assert"
)

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

// -- BuildMessage ------------------------------------------------------------

func TestBuildMessage(test *testing.T) {
	actual := BuildMessage("A", "B", "C", "D")
	printActual(test, actual)
	assert.NotEmpty(test, actual)
}

func TestBuildMessageNoId(test *testing.T) {
	actual := BuildMessage("", "B", "C", "D")
	printActual(test, actual)
	assert.NotEmpty(test, actual)
}

func TestBuildMessageNoDetails(test *testing.T) {
	actual := BuildMessage("A", "B", "C")
	printActual(test, actual)
}

func TestBuildMessageNoMessage(test *testing.T) {
	actual := BuildMessage("A", "B", "")
	printActual(test, actual)
}

func TestBuildMessageNoLevel(test *testing.T) {
	actual := BuildMessage("A", "", "")
	printActual(test, actual)
}

func TestBuildMessageWithNonStrings(test *testing.T) {
	myError := errors.New("Bob's error")
	actual := BuildMessage("A", "myLevel", "myText", "aString", 1, 2.3, myError)
	printActual(test, actual)
}

func TestBuildMessageWithTest(test *testing.T) {
	actual := BuildMessage("A", "myLevel", "myText", test)
	printActual(test, actual)
}

func TestBuildMessageWithJsonText(test *testing.T) {
	jsonText := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA"}`
	actual := BuildMessage("A", "myLevel", jsonText, test)
	printActual(test, actual)
}

// -- BuildMessageFromError ---------------------------------------------------

func TestBuildMessageFromError(test *testing.T) {

	errorMessage1 := BuildMessage("A", "B", "C", "D")
	error1 := errors.New(errorMessage1)

	errorMessage2 := BuildMessageFromError("E", "F", "G", error1, "H")
	error2 := errors.New(errorMessage2)

	errorMessage3 := BuildMessageFromError("I", "J", "K", error2, "L")
	actual := errors.New(errorMessage3)

	printActual(test, actual)
	assert.NotEmpty(test, actual)
}

func TestBuildMessageFromErrorWithSimpleMessage(test *testing.T) {

	error1 := errors.New("Simple error")

	actual := BuildMessageFromError("A", "B", "C", error1, "D")
	printActual(test, actual)
	assert.NotEmpty(test, actual)
}

// -- BuildMessageFromErrorUsingMap -------------------------------------------

func TestBuildMessageFromErrorUsingMap(test *testing.T) {

	errorMessage := BuildMessage("A", "B", "C", "D")
	err := errors.New(errorMessage)

	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
		"TestClass":      test,
	}
	actual := BuildMessageFromErrorUsingMap("A", "B", "C", err, detailsMap)
	printActual(test, actual)
}

// -- BuildMessageUsingMap ----------------------------------------------------

func TestBuildMessageUsingMap(test *testing.T) {
	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
		"TestClass":      test,
	}
	actual := BuildMessageUsingMap("A", "B", "C", detailsMap)
	printActual(test, actual)
}

func TestBuildMessageUsingMapWithJson(test *testing.T) {
	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA"}`
	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
		"TestClass":      test,
		"JSON":           jsonData,
	}
	actual := BuildMessageUsingMap("A", "B", "C", detailsMap)
	printActual(test, actual)
}

// -- ParseMessage ------------------------------------------------------------

func TestParseMessage(test *testing.T) {
	message := BuildMessage("A", "B", "C", "D")
	parsedMessage := ParseMessage(message)

	printResult(test, "Level", parsedMessage.Level)
}

func TestParseMessageUsingMap(test *testing.T) {
	detailsMap := map[string]interface{}{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
	}
	message := BuildMessageUsingMap("A", "B", "C", detailsMap)
	parsedMessage := ParseMessage(message)
	details, ok := parsedMessage.Details.(map[string]interface{})
	if !ok {
		fmt.Printf("Error: %t", ok)
	}

	printResult(test, "First Variable", details["FirstVariable"])
}
