package messageid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name              string
	template          string
	messageNumber     int
	details           []interface{}
	expectedDefault   string
	expectedTemplated string
}{
	{
		name:              "Test case: #1",
		template:          "senzing-9999%04d",
		messageNumber:     1,
		expectedDefault:   `1`,
		expectedTemplated: `senzing-99990001`,
	},
	{
		name:              "Test case: #2",
		template:          "senzing-9999%04d",
		messageNumber:     2,
		details:           []interface{}{123, "bob"},
		expectedDefault:   `2`,
		expectedTemplated: `senzing-99990002`,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageIdInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdDefault - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageIdDefault(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageIdDefault{}
				actual, err := testObject.MessageId(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdSenzing - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageIdSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedTemplated) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageIdSenzing{
					MessageIdTemplate: testCase.template,
				}
				actual, err := testObject.MessageId(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedTemplated, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdTemplated - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageIdTemplated(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedTemplated) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageIdTemplated{
					MessageIdTemplate: testCase.template,
				}
				actual, err := testObject.MessageId(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedTemplated, actual, testCase.name)
			})
		}
	}
}
