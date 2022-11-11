package messageduration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var duration1 = Duration(1111)

var testCases = []struct {
	name            string
	messageNumber   int
	details         []interface{}
	expectedDefault string
	expectedSenzing int64
}{
	{
		name:            "Test case: #1",
		messageNumber:   1000,
		details:         []interface{}{duration1},
		expectedSenzing: int64(1111),
	},
	{
		name:            "Test case: #2",
		messageNumber:   1000,
		details:         []interface{}{int64(2222)},
		expectedSenzing: int64(0),
	},
	{
		name:            "Test case: #3",
		messageNumber:   1000,
		details:         []interface{}{Duration(3333)},
		expectedSenzing: int64(3333),
	},
	{
		name:            "Test case: #4",
		messageNumber:   1000,
		details:         []interface{}{Duration(4444), Duration(4040)},
		expectedSenzing: int64(4444),
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

// ----------------------------------------------------------------------------
// Test interface functions for MessageDurationSenzing - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageDurationSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if testCase.expectedSenzing > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageDurationSenzing{}
				actual, err := testObject.MessageDuration(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}
