package messageformat

import (
	"testing"
)

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatJson - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestJsonBuildError(test *testing.T) {
	testObject := &MessageFormatJson{}
	err := testObject.BuildError("id-1", "try-again", "text-1", 123, "bob")
	test.Logf("%v", err)
}

// -- BuildMessage ------------------------------------------------------------

func TestJsonBuildMessage(test *testing.T) {
	testObject := &MessageFormatJson{}
	message := testObject.BuildMessage("id-2", "try-again", "text-2", 123, "bob")
	test.Logf("%s", message)
}

func TestJsonBuildMessageNoId(test *testing.T) {
	testObject := &MessageFormatJson{}
	message := testObject.BuildMessage("", "try-again", "text-3", 123, "bob")
	test.Logf("%s", message)
}

func TestJsonBuildMessageNoStatus(test *testing.T) {
	testObject := &MessageFormatJson{}
	message := testObject.BuildMessage("id-4", "", "text-4", 123, "bob")
	test.Logf("%s", message)
}

func TestJsonBuildMessageNoText(test *testing.T) {
	testObject := &MessageFormatJson{}
	message := testObject.BuildMessage("id-5", "try-again", "", 123, "bob")
	test.Logf("%s", message)
}

func TestJsonBuildMessageNoDetails(test *testing.T) {
	testObject := &MessageFormatJson{}
	message := testObject.BuildMessage("id-6", "try-again", "text-6")
	test.Logf("%s", message)
}

func TestJsonBuildMessageNothing(test *testing.T) {
	testObject := &MessageFormatJson{}
	message := testObject.BuildMessage("", "", "")
	test.Logf("%s", message)
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatTerse - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestTerseBuildError(test *testing.T) {
	testObject := &MessageFormatTerse{}
	err := testObject.BuildError("id-1", "try-again", "text-1", 123, "bob")
	test.Logf("%v", err)
}

// -- BuildMessage ------------------------------------------------------------

func TestTerseBuildMessage(test *testing.T) {
	testObject := &MessageFormatTerse{}
	message := testObject.BuildMessage("id-1", "try-again", "text-1", 123, "bob")
	test.Logf("%s", message)
}
