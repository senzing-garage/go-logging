package messagedetails

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	messageNumber   int
	details         []interface{}
	expectedDefault map[string]interface{}
	expectedSenzing map[string]interface{}
}{
	{
		name:            "messagedetails-01",
		messageNumber:   1001,
		expectedDefault: map[string]interface{}{},
		expectedSenzing: map[string]interface{}{},
	},
	{
		name:            "messagedetails-02",
		messageNumber:   1002,
		details:         []interface{}{"A", 1, 1.1, true},
		expectedDefault: map[string]interface{}{"1": "A", "2": 1, "3": 1.1, "4": "true"},
		expectedSenzing: map[string]interface{}{"1": "A", "2": 1, "3": 1.1, "4": "true"},
	},
	{
		name:            "messagedetails-03",
		messageNumber:   1003,
		details:         []interface{}{"A", errors.New("test error")},
		expectedDefault: map[string]interface{}{"1": "A"},
		expectedSenzing: map[string]interface{}{"1": "A"},
	},
	{
		name:            "messagedetails-04",
		messageNumber:   1004,
		details:         []interface{}{"A", map[string]string{"Name": "Bob"}},
		expectedDefault: map[string]interface{}{"1": "A", "Name": "Bob"},
		expectedSenzing: map[string]interface{}{"1": "A", "Name": "Bob"},
	},
	{
		name:            "messagedetails-05",
		messageNumber:   1005,
		details:         []interface{}{"{\"A\": \"A JSON example\"}"},
		expectedDefault: map[string]interface{}{"1": json.RawMessage(`{"A": "A JSON example"}`)},
		expectedSenzing: map[string]interface{}{"1": json.RawMessage(`{"A": "A JSON example"}`)},
	},
	{
		name:            "messagedetails-06",
		messageNumber:   1006,
		details:         []interface{}{`{"A": "A JSON example"}`},
		expectedDefault: map[string]interface{}{"1": json.RawMessage(`{"A": "A JSON example"}`)},
		expectedSenzing: map[string]interface{}{"1": json.RawMessage(`{"A": "A JSON example"}`)},
	},
	{
		name:            "messagedetails-07",
		messageNumber:   1007,
		details:         []interface{}{`{"A": {"B": "A JSON example"}}`},
		expectedDefault: map[string]interface{}{"1": json.RawMessage(`{"A": {"B": "A JSON example"}}`)},
		expectedSenzing: map[string]interface{}{"1": json.RawMessage(`{"A": {"B": "A JSON example"}}`)},
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageDetailsInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDetailsDefault
// ----------------------------------------------------------------------------

func TestMessageDetailsDefault(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name+"-Default", func(test *testing.T) {
				testObject := &MessageDetailsDefault{}
				actual, err := testObject.MessageDetails(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDetailsNull
// ----------------------------------------------------------------------------

func TestMessageDetailsNull(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-Null", func(test *testing.T) {
			testObject := &MessageDetailsNull{}
			actual, err := testObject.MessageDetails(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, nil, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageDetailsSenzing
// ----------------------------------------------------------------------------

func TestMessageDetailsSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name+"-Senzing", func(test *testing.T) {
				testObject := &MessageDetailsSenzing{}
				actual, err := testObject.MessageDetails(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}
