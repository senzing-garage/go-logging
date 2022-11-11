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
	expectedSenzing string
}{
	{
		name:            "messagelocation-01",
		callerSkip:      0,
		messageNumber:   1000,
		expectedSenzing: "In MessageLocation() at messagelocation_senzing.go:31",
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
// Test interface functions for MessageDateNull
// ----------------------------------------------------------------------------

func TestMessageLocationNull(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-Null", func(test *testing.T) {
			testObject := &MessageLocationNull{}
			actual, err := testObject.MessageLocation(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, "", actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdSenzing
// ----------------------------------------------------------------------------

func TestMessageLocationSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name+"-Senzing", func(test *testing.T) {
				testObject := &MessageLocationSenzing{
					CallerSkip: testCase.callerSkip,
				}
				actual, err := testObject.MessageLocation(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}
