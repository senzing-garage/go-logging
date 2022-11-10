package messagedate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	messageNumber   int
	details         []interface{}
	expectedDefault string
}{
	{
		name:          "Test case: #1",
		messageNumber: 1000,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageDateInterface, err error) {
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
			testObject := &MessageDateNull{}
			actual, err := testObject.MessageDate(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, "", actual, testCase.name)
		})
	}
}
