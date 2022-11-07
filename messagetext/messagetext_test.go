/*
For more information, see [test].

[test]: messagetext/messagetext_test.go
*/
package messagetext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	idMessages      map[int]string
	messageNumber   int
	details         []interface{}
	expectedDefault string
}{
	{
		name:          "Test case: #1",
		messageNumber: 1,
		idMessages: map[int]string{
			1: `Bob's middle initial is "%s" and his favorite number is %d.`,
		},
		details:         []interface{}{"A", 1},
		expectedDefault: `Bob's middle initial is "A" and his favorite number is 1.`,
	},
	{
		name:          "Test case: #2 - Specific Message Number",
		messageNumber: 1,
		idMessages: map[int]string{
			2: "Sally got an \"%s\" on the paper.",
		},
		details:         []interface{}{"A", 1, MsgNumber(2)},
		expectedDefault: `Sally got an "A" on the paper.`,
	},
	{
		name:          "Test case: #3 - GT/LT",
		messageNumber: 1,
		idMessages: map[int]string{
			1: `>>>  Try this: "%s"  <<<`,
		},
		details:         []interface{}{"Test GT/LT"},
		expectedDefault: `>>>  Try this: "Test GT/LT"  <<<`,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageTextInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageTextSenzing - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageTextSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageTextSenzing{
					IdMessages: testCase.idMessages,
				}
				actual, err := testObject.MessageText(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageTextTemplated - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageTextTemplated(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageTextTemplated{
					IdMessages: testCase.idMessages,
				}
				actual, err := testObject.MessageText(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}
