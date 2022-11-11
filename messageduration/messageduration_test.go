package messageduration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	messageNumber   int
	details         []interface{}
	expectedDefault int64
	expectedSenzing int64
}{
	{
		name:            "Test case: #1",
		messageNumber:   1001,
		details:         []interface{}{getDuration(1)},
		expectedDefault: int64(0),
		expectedSenzing: int64(1),
	},
	{
		name:            "Test case: #2",
		messageNumber:   1002,
		details:         []interface{}{int64(2222)},
		expectedDefault: int64(0),
		expectedSenzing: int64(0),
	},
	{
		name:            "Test case: #3",
		messageNumber:   1003,
		details:         []interface{}{getDuration(3333)},
		expectedDefault: int64(3),
		expectedSenzing: int64(3333),
	},
	{
		name:            "Test case: #4",
		messageNumber:   1000,
		details:         []interface{}{getDuration(4444)},
		expectedDefault: int64(4),
		expectedSenzing: int64(4444),
	},
	{
		name:            "Test case: #5",
		messageNumber:   1000,
		details:         []interface{}{getDuration(555555)},
		expectedDefault: int64(555),
		expectedSenzing: int64(555555),
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getDuration(durationInNanoseconds int64) time.Duration {
	durationString := fmt.Sprintf("%dns", durationInNanoseconds)
	result, err := time.ParseDuration(durationString)
	if err != nil {
		fmt.Printf(">>>> Error: %s\n", err.Error())
	}
	return result
}

func testError(test *testing.T, testObject MessageDurationInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDurationNull
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
// Test interface functions for MessageDurationDefault
// ----------------------------------------------------------------------------

func TestMessageDurationDefault(test *testing.T) {
	for _, testCase := range testCases {
		if testCase.expectedDefault > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageDurationDefault{}
				actual, err := testObject.MessageDuration(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDurationSenzing
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
