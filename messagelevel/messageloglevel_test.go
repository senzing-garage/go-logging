package messagelevel

import (
	"errors"
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

var idRanges = map[int]logger.Level{
	0000: logger.LevelInfo,
	1000: logger.LevelWarn,
	2000: logger.LevelError,
	3000: logger.LevelDebug,
	4000: logger.LevelTrace,
	5000: logger.LevelFatal,
	6000: logger.LevelPanic,
}

var testCases = []struct {
	name            string
	idRanges        map[int]logger.Level
	messageNumber   int
	details         []interface{}
	expectedDefault logger.Level
}{
	{
		name:            "Test case: #1",
		idRanges:        idRanges,
		messageNumber:   0,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelInfo,
	},
	{
		name:            "Test case: #2",
		idRanges:        idRanges,
		messageNumber:   1000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelWarn,
	},
	{
		name:            "Test case: #3",
		idRanges:        idRanges,
		messageNumber:   2000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelError,
	},
	{
		name:            "Test case: #4",
		idRanges:        idRanges,
		messageNumber:   3000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelDebug,
	},
	{
		name:            "Test case: #5",
		idRanges:        idRanges,
		messageNumber:   4000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelTrace,
	},
	{
		name:            "Test case: #6",
		idRanges:        idRanges,
		messageNumber:   5000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelFatal,
	},
	{
		name:            "Test case: #7",
		idRanges:        idRanges,
		messageNumber:   6000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "Test case: #8",
		idRanges:        idRanges,
		messageNumber:   7000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "Test case: #9",
		idRanges:        idRanges,
		messageNumber:   9999,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "Test case: #10",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelInfo},
		expectedDefault: logger.LevelInfo,
	},
	{
		name:            "Test case: #11",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelWarn},
		expectedDefault: logger.LevelWarn,
	},
	{
		name:            "Test case: #12",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelError},
		expectedDefault: logger.LevelError,
	},
	{
		name:            "Test case: #13",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelDebug},
		expectedDefault: logger.LevelDebug,
	},
	{
		name:            "Test case: #14",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelTrace},
		expectedDefault: logger.LevelTrace,
	},
	{
		name:            "Test case: #15",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelFatal},
		expectedDefault: logger.LevelFatal,
	},
	{
		name:            "Test case: #16",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelPanic},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "Test case: #17",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelPanic},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "Test case: #18",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelInfo, logger.LevelDebug},
		expectedDefault: logger.LevelInfo,
	},
	{
		name:            "Test case: #19",
		idRanges:        idRanges,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelDebug, logger.LevelInfo},
		expectedDefault: logger.LevelDebug,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageLevelInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusById
// ----------------------------------------------------------------------------

func TestMessageLevelByIdRange(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageLevelByIdRange{
				IdRanges: testCase.idRanges,
			}
			actual, err := testObject.MessageLevel(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelDefault
// ----------------------------------------------------------------------------

func TestMessageLevelDefault(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageLevelDefault{
				IdRanges: testCase.idRanges,
			}
			actual, err := testObject.MessageLevel(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelSenzing
// ----------------------------------------------------------------------------

func TestMessageLevelSenzing(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageLevelSenzing{
				IdRanges: testCase.idRanges,
			}
			actual, err := testObject.MessageLevel(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelSenzingApi
// ----------------------------------------------------------------------------

func TestMessageLevelSenzingApi(test *testing.T) {

	idRangesStrings := make(map[int]string)
	for key, value := range idRanges {
		idRangesStrings[key] = logger.LevelToTextMap[value]
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageLevelSenzingApi{
				IdRanges:   idRangesStrings,
				IdStatuses: idRangesStrings,
			}
			actual, err := testObject.MessageLevel(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
		})
	}
}

func TestSenzingApiMessageLevelWithErrors(test *testing.T) {
	expected := logger.LevelError
	anError1 := errors.New("0019E|Configuration not found")
	anError2 := errors.New("0099E|Made up error")

	idRangesStrings := make(map[int]string)
	for key, value := range idRanges {
		idRangesStrings[key] = logger.LevelToTextMap[value]
	}

	testObject := &MessageLevelSenzingApi{
		IdRanges:   idRangesStrings,
		IdStatuses: idRangesStrings,
	}
	actual, err := testObject.MessageLevel(1, "A", 1, testObject, anError1, anError2)
	testError(test, testObject, err)
	assert.Equal(test, expected, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelStatic
// ----------------------------------------------------------------------------

func TestMessageLevelStatic(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageLevelStatic{
				LogLevel: logger.LevelWarn,
			}
			actual, err := testObject.MessageLevel(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, logger.LevelWarn, actual, testCase.name)
		})
	}
}
