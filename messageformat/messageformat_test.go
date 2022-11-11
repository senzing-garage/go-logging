package messageformat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	date            string
	time            string
	level           string
	location        string
	id              string
	status          string
	text            string
	duration        int64
	details         []interface{}
	expectedDefault string
	expectedJson    string
	expectedSenzing string
}{
	{
		name:            "Test case: #1",
		id:              "id-1",
		status:          "status-1",
		text:            "text-1",
		details:         []interface{}{123, "bob"},
		expectedDefault: `id-1: (status-1) text-1 [map[1:123 2:bob]]`,
		expectedJson:    `{"id":"id-1","status":"status-1","text":"text-1","details":{"1":123,"2":"bob"}}`,
	},
	{
		name:            "Test case: #2 - no id",
		status:          "status-2",
		text:            "text-2",
		details:         []interface{}{123, "bob"},
		expectedDefault: `(status-2) text-2 [map[1:123 2:bob]]`,
		expectedJson:    `{"status":"status-2","text":"text-2","details":{"1":123,"2":"bob"}}`,
	},
	{
		name:            "Test case: #3 - no status",
		id:              "id-3",
		text:            "text-3",
		details:         []interface{}{123, "bob"},
		expectedDefault: `id-3: text-3 [map[1:123 2:bob]]`,
		expectedJson:    `{"id":"id-3","text":"text-3","details":{"1":123,"2":"bob"}}`,
	},
	{
		name:            "Test case: #4 - no text",
		id:              "id-4",
		status:          "status-4",
		details:         []interface{}{123, "bob"},
		expectedDefault: `id-4: (status-4) [map[1:123 2:bob]]`,
		expectedJson:    `{"id":"id-4","status":"status-4","details":{"1":123,"2":"bob"}}`,
	},
	{
		name:            "Test case: #5 - no details",
		id:              "id-5",
		status:          "status-5",
		text:            "text-5",
		expectedDefault: `id-5: (status-5) text-5`,
		expectedJson:    `{"id":"id-5","status":"status-5","text":"text-5"}`,
	},
	{
		name:            "Test case: #6",
		expectedDefault: ``,
		expectedJson:    `{}`,
	},
	{
		name:            "Test case: #10 - date",
		date:            "date-10",
		time:            "time-10",
		level:           "level-10",
		location:        "location-10",
		id:              "id-10",
		status:          "status-10",
		text:            "text-10",
		duration:        int64(0),
		expectedDefault: `id-10: (status-10) text-10`,
		expectedSenzing: `{"date":"date-10","time":"time-10","level":"level-10","id":"id-10","status":"status-10","text":"text-10","location":"location-10"}`,
	},
	{
		name:            "Test case: #11 - Add duration",
		date:            "date-11",
		time:            "time-11",
		level:           "level-11",
		location:        "location-11",
		id:              "id-11",
		status:          "status-11",
		text:            "text-11",
		duration:        int64(11),
		expectedDefault: `id-11: (status-11) text-11`,
		expectedSenzing: `{"date":"date-11","time":"time-11","level":"level-11","id":"id-11","status":"status-11","text":"text-11","duration":11,"location":"location-11"}`,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageFormatInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatDefault
// ----------------------------------------------------------------------------

func TestMessageFormatDefault(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageFormatDefault{}
				actual, err := testObject.Message(testCase.date, testCase.time, testCase.level, testCase.location, testCase.id, testCase.status, testCase.text, testCase.duration, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatJson
// ----------------------------------------------------------------------------

func TestMessageFormatJson(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedJson) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageFormatJson{}
				actual, err := testObject.Message(testCase.date, testCase.time, testCase.level, testCase.location, testCase.id, testCase.status, testCase.text, testCase.duration, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedJson, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageFormatSenzing
// ----------------------------------------------------------------------------

func TestMessageFormatSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageFormatSenzing{}
				actual, err := testObject.Message(testCase.date, testCase.time, testCase.level, testCase.location, testCase.id, testCase.status, testCase.text, testCase.duration, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}
