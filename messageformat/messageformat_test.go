package messageformat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	id              string
	status          string
	text            string
	details         []interface{}
	expectedJSON    string
	expectedDefault string
}{
	{
		name:            "Test case: #1",
		id:              "id-1",
		status:          "status-1",
		text:            "text-1",
		details:         []interface{}{123, "bob"},
		expectedJSON:    `{"id":"id-1","status":"status-1","text":"text-1","details":{"1":123,"2":"bob"}}`,
		expectedDefault: `id-1: (status-1) text-1 [map[1:123 2:"bob"]]`,
	},
	{
		name:            "Test case: #2 - no id",
		status:          "status-2",
		text:            "text-2",
		details:         []interface{}{123, "bob"},
		expectedJSON:    `{"status":"status-2","text":"text-2","details":{"1":123,"2":"bob"}}`,
		expectedDefault: `(status-2) text-2 [map[1:123 2:"bob"]]`,
	},
	{
		name:            "Test case: #3 - no status",
		id:              "id-3",
		text:            "text-3",
		details:         []interface{}{123, "bob"},
		expectedJSON:    `{"id":"id-3","text":"text-3","details":{"1":123,"2":"bob"}}`,
		expectedDefault: `id-3: text-3 [map[1:123 2:"bob"]]`,
	},
	{
		name:            "Test case: #4 - no text",
		id:              "id-4",
		status:          "status-4",
		details:         []interface{}{123, "bob"},
		expectedJSON:    `{"id":"id-4","status":"status-4","details":{"1":123,"2":"bob"}}`,
		expectedDefault: `id-4: (status-4) [map[1:123 2:"bob"]]`,
	},
	{
		name:            "Test case: #5 - no details",
		id:              "id-5",
		status:          "status-5",
		text:            "text-5",
		expectedJSON:    `{"id":"id-5","status":"status-5","text":"text-5"}`,
		expectedDefault: `id-5: (status-5) text-5`,
	},
	{
		name:            "Test case: #6",
		expectedJSON:    `{}`,
		expectedDefault: ``,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageFormatInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatDefault - names begin with "Test"
// ----------------------------------------------------------------------------

func TestDefaultMessages(test *testing.T) {
	testObject := &MessageFormatDefault{}
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				actual, err := testObject.Message(testCase.id, testCase.status, testCase.text, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatJson - names begin with "Test"
// ----------------------------------------------------------------------------

func TestJsonMessages(test *testing.T) {
	testObject := &MessageFormatJson{}
	for _, testCase := range testCases {
		if len(testCase.expectedJSON) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				actual, err := testObject.Message(testCase.id, testCase.status, testCase.text, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedJSON, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatJson - names begin with "Test"
// ----------------------------------------------------------------------------

func TestSenzingMessages(test *testing.T) {
	testObject := &MessageFormatSenzing{}
	for _, testCase := range testCases {
		if len(testCase.expectedJSON) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				actual, err := testObject.Message(testCase.id, testCase.status, testCase.text, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedJSON, actual, testCase.name)
			})
		}
	}
}
