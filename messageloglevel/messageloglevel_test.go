package messageloglevel

import (
	"errors"
	"fmt"
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name            string
	messageNumber   int
	details         []interface{}
	expectedDefault logger.Level
}{
	{
		name:            "Test case: #1",
		messageNumber:   1,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelInfo,
	},
	{
		name:            "Test case: #1",
		messageNumber:   1001,
		details:         []interface{}{123, "bob"},
		expectedDefault: logger.LevelInfo,
	},
}

const printResults = 1

var idLevels = map[int]string{
	0001: logger.LevelInfoName,
	0002: logger.LevelWarnName,
	0003: logger.LevelErrorName,
	0004: logger.LevelDebugName,
	0005: logger.LevelTraceName,
	0006: logger.LevelFatalName,
	0007: logger.LevelPanicName,
}

var idRanges = map[int]string{
	0000: logger.LevelInfoName,
	1000: logger.LevelWarnName,
	2000: logger.LevelErrorName,
	3000: logger.LevelDebugName,
	4000: logger.LevelTraceName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func printResult(test *testing.T, title string, result interface{}) {
	if printResults == 1 {
		test.Logf("%s: %v", title, fmt.Sprintf("%v", result))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, testObject MessageLogLevelInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelSenzingApi - names begin with "Test"
// ----------------------------------------------------------------------------

func TestSenzingApiMessageLogLevels(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}

	actual, err := testObject.MessageLogLevel(0)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelInfo)
}

func TestSenzingApiMessageLogLevelInfo(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(0)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelInfo)
}

func TestSenzingApiMessageLogLevelWarn(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(1000)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelWarn)
}

func TestSenzingApiMessageLogLevelError(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(2000)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelError)
}

func TestSenzingApiMessageLogLevelDebug(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(3000)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelDebug)
}

func TestSenzingApiMessageLogLevelTrace(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(4000)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelTrace)
}

func TestSenzingApiMessageLogLevelFatal(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(5000)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelFatal)
}

func TestSenzingApiMessageLogLevelPanic(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(6000)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelPanic)
}

func TestSenzingApiMessageLogLevelUnknown(test *testing.T) {
	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(7000)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelPanic)
}

func TestSenzingApiMessageLogLevelWithErrors(test *testing.T) {
	anError1 := errors.New("0019E|Configuration not found")
	anError2 := errors.New("0099E|Made up error")

	testObject := &MessageLogLevelSenzingApi{
		IdRanges: idRanges,
	}
	actual, err := testObject.MessageLogLevel(1, "A", 1, testObject, anError1, anError2)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelError)

}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelInfo - names begin with "Test"
// ----------------------------------------------------------------------------

// -- MessageLogLevel ---------------------------------------------------------

func TestMessageLogLevelDefault(test *testing.T) {
	testObject := &MessageLogLevelDefault{}
	actual, err := testObject.MessageLogLevel(1, "This is message text")
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelInfo)
}

func TestMessageLogLevelDefaultWarn(test *testing.T) {
	testObject := &MessageLogLevelDefault{}
	actual, err := testObject.MessageLogLevel(1, "This is message text", logger.LevelWarn)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelWarn)
}

func TestMessageLogLevelDefaultError(test *testing.T) {
	testObject := &MessageLogLevelDefault{}
	actual, err := testObject.MessageLogLevel(1, "This is message text", logger.LevelWarn, logger.LevelError)
	testError(test, testObject, err)
	printActual(test, actual)
	assert.True(test, actual == logger.LevelError)
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdSenzing - names begin with "Test"
// ----------------------------------------------------------------------------

func TestDefaultMessageLogLevel(test *testing.T) {
	for _, testCase := range testCases {
		if testCase.expectedDefault > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageLogLevelDefault{}
				actual, err := testObject.MessageLogLevel(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageIdSenzing - names begin with "Test"
// ----------------------------------------------------------------------------

// func TestDefaultMessageId(test *testing.T) {
// 	for _, testCase := range testCases {
// 		if len(testCase.expectedTemplated) > 0 {
// 			test.Run(testCase.name, func(test *testing.T) {
// 				testObject := &MessageIdSenzing{
// 					MessageIdTemplate: testCase.template,
// 				}
// 				actual, err := testObject.MessageId(testCase.messageNumber, testCase.details...)
// 				testError(test, testObject, err)
// 				assert.Equal(test, testCase.expectedTemplated, actual, testCase.name)
// 			})
// 		}
// 	}
// }
