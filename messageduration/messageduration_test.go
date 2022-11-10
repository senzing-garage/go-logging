package messageduration

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

func testError(test *testing.T, testObject MessageDurationInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDurationNull - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageDurationNull(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageDurationNull{}
			actual, err := testObject.MessageDuration(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, int64(0), actual, testCase.name)
		})
	}
}
