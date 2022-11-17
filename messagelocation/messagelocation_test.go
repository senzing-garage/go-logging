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
		name:            "messagelocation-01-skip-0",
		callerSkip:      0,
		messageNumber:   1000,
		expectedDefault: "In MessageLocation() at messagelocation_default.go:",
		expectedSenzing: "In MessageLocation() at messagelocation_senzing.go:",
	},
	{
		name:            "messagelocation-02-skip-1",
		callerSkip:      1,
		messageNumber:   1000,
		expectedDefault: "In func1() at messagelocation_test.go:",
		expectedSenzing: "In func1() at messagelocation_test.go:",
	},
	{
		name:            "messagelocation-03-skip-2",
		callerSkip:      2,
		messageNumber:   1000,
		expectedDefault: "In tRunner() at testing.go:",
		expectedSenzing: "In tRunner() at testing.go:",
	},
	{
		name:            "messagelocation-04-skip-3",
		callerSkip:      3,
		messageNumber:   1000,
		expectedDefault: "In goexit() at asm_amd64.s:",
		expectedSenzing: "In goexit() at asm_amd64.s:",
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
// Test interface functions for MessageIdDefault
// ----------------------------------------------------------------------------

func TestMessageLocationDefault(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name+"-Default", func(test *testing.T) {
				testObject := &MessageLocationDefault{
					CallerSkip: testCase.callerSkip,
				}
				actual, err := testObject.MessageLocation(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Contains(test, actual, testCase.expectedDefault, testCase.name)
			})
		}
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
				assert.Contains(test, actual, testCase.expectedSenzing, testCase.name)
			})
		}
	}
}
