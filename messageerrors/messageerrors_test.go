package messageerrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	messageNumber   int
	details         []interface{}
	expectedDefault []interface{}
	expectedSenzing []interface{}
}{
	{
		name:            "messageerrors-01",
		messageNumber:   1,
		expectedDefault: []interface{}([]interface{}(nil)),
		expectedSenzing: []interface{}([]interface{}(nil)),
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageErrorsInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdDefault
// ----------------------------------------------------------------------------

func TestMessageErrorsDefault(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name+"-Default", func(test *testing.T) {
				testObject := &MessageErrorsDefault{}
				actual, err := testObject.MessageErrors(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDetailsNull
// ----------------------------------------------------------------------------

func TestMessageErrorsNull(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-Null", func(test *testing.T) {
			testObject := &MessageErrorsNull{}
			actual, err := testObject.MessageErrors(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, nil, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdSenzing
// ----------------------------------------------------------------------------

func TestMessageErrorsSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name+"-Senzing", func(test *testing.T) {
				testObject := &MessageErrorsSenzing{}
				actual, err := testObject.MessageErrors(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}
