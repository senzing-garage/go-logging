package messagelevel

import (
	"errors"
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	idRanges        map[int]logger.Level
	messageNumber   int
	details         []interface{}
	expectedDefault logger.Level
}{
	{
		name:            "messagelevel-01",
		idRanges:        IdRangesAsLevel,
		messageNumber:   0,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelTrace,
	},
	{
		name:            "messagelevel-02",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelDebug,
	},
	{
		name:            "messagelevel-03",
		idRanges:        IdRangesAsLevel,
		messageNumber:   2000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelInfo,
	},
	{
		name:            "messagelevel-04",
		idRanges:        IdRangesAsLevel,
		messageNumber:   3000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelWarn,
	},
	{
		name:            "messagelevel-05",
		idRanges:        IdRangesAsLevel,
		messageNumber:   4000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelError,
	},
	{
		name:            "messagelevel-06",
		idRanges:        IdRangesAsLevel,
		messageNumber:   5000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelFatal,
	},
	{
		name:            "messagelevel-07",
		idRanges:        IdRangesAsLevel,
		messageNumber:   6000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "messagelevel-08",
		idRanges:        IdRangesAsLevel,
		messageNumber:   7000,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "messagelevel-09",
		idRanges:        IdRangesAsLevel,
		messageNumber:   9999,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "messagelevel-10",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelTrace},
		expectedDefault: logger.LevelTrace,
	},
	{
		name:            "messagelevel-11",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelDebug},
		expectedDefault: logger.LevelDebug,
	},
	{
		name:            "messagelevel-12",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelInfo},
		expectedDefault: logger.LevelInfo,
	},
	{
		name:            "messagelevel-13",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelWarn},
		expectedDefault: logger.LevelWarn,
	},
	{
		name:            "messagelevel-14",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelError},
		expectedDefault: logger.LevelError,
	},
	{
		name:            "messagelevel-15",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelFatal},
		expectedDefault: logger.LevelFatal,
	},
	{
		name:            "messagelevel-16",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelPanic},
		expectedDefault: logger.LevelPanic,
	},
	{
		name:            "messagelevel-17",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelInfo, logger.LevelDebug},
		expectedDefault: logger.LevelDebug,
	},
	{
		name:            "messagelevel-18",
		idRanges:        IdRangesAsLevel,
		messageNumber:   1001,
		details:         []interface{}{123, "bob", logger.LevelDebug, logger.LevelInfo},
		expectedDefault: logger.LevelInfo,
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
		test.Run(testCase.name+"-ByIdRange", func(test *testing.T) {
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
		test.Run(testCase.name+"-Default", func(test *testing.T) {
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
		test.Run(testCase.name+"-Senzing", func(test *testing.T) {
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
	for key, value := range IdRangesAsLevel {
		idRangesStrings[key] = logger.LevelToTextMap[value]
	}

	for _, testCase := range testCases {
		test.Run(testCase.name+"-SenzingApi", func(test *testing.T) {
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
	for key, value := range IdRangesAsLevel {
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
		test.Run(testCase.name+"-Static", func(test *testing.T) {
			testObject := &MessageLevelStatic{
				LogLevel: logger.LevelWarn,
			}
			actual, err := testObject.MessageLevel(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, logger.LevelWarn, actual, testCase.name)
		})
	}
}
