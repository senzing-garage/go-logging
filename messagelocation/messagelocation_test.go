package messagelocation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	callerSkip      int
	messageNumber   int
	details         []interface{}
	expectedDefault string
}{
	{
		name:          "Test case: #1",
		callerSkip:    1,
		messageNumber: 1000,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageLocationInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDateNull - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageDateNull(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageLocationNull{}
			actual, err := testObject.MessageLocation(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, "", actual, testCase.name)
		})
	}
}
