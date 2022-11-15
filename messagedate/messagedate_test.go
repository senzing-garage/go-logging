package messagedate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name             string
	messageNumber    int
	messageTimestamp time.Time
	details          []interface{}
	expectedDefault  string
	expectedSenzing  string
}{
	{
		name:             "messagedate-01",
		messageNumber:    1001,
		messageTimestamp: time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
		expectedDefault:  "2000-01-01",
		expectedSenzing:  "2000-01-01",
	},
	{
		name:             "messagedate-02",
		messageNumber:    1002,
		messageTimestamp: time.Date(2999, time.December, 31, 0, 0, 0, 0, time.UTC),
		expectedDefault:  "2999-12-31",
		expectedSenzing:  "2999-12-31",
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
// Test interface functions for MessageDateDefault
// ----------------------------------------------------------------------------

func TestMessageDateDefault(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name+"-Default", func(test *testing.T) {
				testObject := &MessageDateDefault{}
				actual, err := testObject.MessageDate(testCase.messageNumber, testCase.messageTimestamp, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDateNull
// ----------------------------------------------------------------------------

func TestMessageDateNull(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-Null", func(test *testing.T) {
			testObject := &MessageDateNull{}
			actual, err := testObject.MessageDate(testCase.messageNumber, testCase.messageTimestamp, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, "", actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDateSenzing
// ----------------------------------------------------------------------------

func TestMessageDateSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name+"-Senzing", func(test *testing.T) {
				testObject := &MessageDateSenzing{}
				actual, err := testObject.MessageDate(testCase.messageNumber, testCase.messageTimestamp, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}
