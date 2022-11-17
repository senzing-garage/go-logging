package messagelevel

import (
	"errors"
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

// Errors used in testing.
var (
	errorUnknown       = errors.New("0000E|Mock error: Unknown")
	errorTrace         = errors.New("9990E|Mock error: TRACE")
	errorDebug         = errors.New("9991E|Mock error: DEBUG")
	errorInfo          = errors.New("9992E|Mock error: INFO")
	errorWarn          = errors.New("9993E|Mock error: WARN")
	errorError         = errors.New("9994E|Mock error: ERROR")
	errorRetryable     = errors.New("9995E|Mock error: Retryable")
	errorBadUserInput  = errors.New("9996E|Mock error: Bad user input")
	errorUnrecoverable = errors.New("9997E|Mock error: Unrecoverable")
	errorFatal         = errors.New("9998E|Mock error: FATAL")
	errorPanic         = errors.New("9999E|Mock error: PANIC")
)

var testCases = []struct {
	name               string
	idLevelRanges      map[int]logger.Level
	messageNumber      int
	details            []interface{}
	expectedDefault    logger.Level
	expectedSenzingApi logger.Level
}{
	{
		name:               "messagelevel-01-Trace",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      0,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelTrace,
		expectedSenzingApi: logger.LevelTrace,
	},
	{
		name:               "messagelevel-02-Debug",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1000,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelDebug,
	},
	{
		name:               "messagelevel-03-Info",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      2000,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelInfo,
		expectedSenzingApi: logger.LevelInfo,
	},
	{
		name:               "messagelevel-04-Warn",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      3000,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelWarn,
		expectedSenzingApi: logger.LevelWarn,
	},
	{
		name:               "messagelevel-05-Error",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      4000,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelError,
		expectedSenzingApi: logger.LevelError,
	},
	{
		name:               "messagelevel-06-Fatal",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      5000,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelFatal,
		expectedSenzingApi: logger.LevelFatal,
	},
	{
		name:               "messagelevel-07-Panic",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      6000,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelPanic,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-08-Panic-7000",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      7000,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelPanic,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-09-Panic-9999",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      9999,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelPanic,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-11-MessageId-Trace",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelTrace,
		expectedSenzingApi: logger.LevelTrace,
	},
	{
		name:               "messagelevel-12-MessageId-Debug",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelDebug,
	},
	{
		name:               "messagelevel-13-MessageId-Info",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      2001,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelInfo,
		expectedSenzingApi: logger.LevelInfo,
	},
	{
		name:               "messagelevel-14-MessageId-Warn",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      3001,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelWarn,
		expectedSenzingApi: logger.LevelWarn,
	},
	{
		name:               "messagelevel-15-MessageId-Error",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      4001,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelError,
		expectedSenzingApi: logger.LevelError,
	},
	{
		name:               "messagelevel-16-MessageId-Fatal",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      5001,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelFatal,
		expectedSenzingApi: logger.LevelFatal,
	},
	{
		name:               "messagelevel-17-MessageId-Panic",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      6001,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelPanic,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-18-MessageId-Panic",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      7001,
		details:            []interface{}{123, "bob"},
		expectedDefault:    logger.LevelPanic,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-20-LevelTrace",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelTrace},
		expectedDefault:    logger.LevelTrace,
		expectedSenzingApi: logger.LevelTrace,
	},
	{
		name:               "messagelevel-21-LevelDebug",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelDebug},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelDebug,
	},
	{
		name:               "messagelevel-22-LevelInfo",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelInfo},
		expectedDefault:    logger.LevelInfo,
		expectedSenzingApi: logger.LevelInfo,
	},
	{
		name:               "messagelevel-23-LevelWarn",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelWarn},
		expectedDefault:    logger.LevelWarn,
		expectedSenzingApi: logger.LevelWarn,
	},
	{
		name:               "messagelevel-24-LevelError",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelError},
		expectedDefault:    logger.LevelError,
		expectedSenzingApi: logger.LevelError,
	},
	{
		name:               "messagelevel-25-LevelFatal",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelFatal},
		expectedDefault:    logger.LevelFatal,
		expectedSenzingApi: logger.LevelFatal,
	},
	{
		name:               "messagelevel-26-LevelPanic",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelPanic},
		expectedDefault:    logger.LevelPanic,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-27-LevelInfo-LevelDebug",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelInfo, logger.LevelDebug},
		expectedDefault:    logger.LevelInfo,
		expectedSenzingApi: logger.LevelInfo,
	},
	{
		name:               "messagelevel-28-LevelDebug-LevelInfo",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", logger.LevelDebug, logger.LevelInfo},
		expectedDefault:    logger.LevelInfo,
		expectedSenzingApi: logger.LevelInfo,
	},
	{
		name:               "messagelevel-30-Error-unknown",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorUnknown},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelDebug,
	},
	{
		name:               "messagelevel-31-Error-trace",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorTrace},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelTrace,
	},
	{
		name:               "messagelevel-32-Error-debug",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorDebug},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelDebug,
	},
	{
		name:               "messagelevel-33-Error-info",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorInfo},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelInfo,
	},
	{
		name:               "messagelevel-34-Error-warn",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorWarn},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelWarn,
	},
	{
		name:               "messagelevel-35-Error-error",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorError},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelError,
	},
	{
		name:               "messagelevel-36-Error-retryable",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorRetryable},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelError,
	},
	{
		name:               "messagelevel-37-Error-baduserinput",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorBadUserInput},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelError,
	},
	{
		name:               "messagelevel-38-Error-unrecoverable",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorUnrecoverable},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelError,
	},
	{
		name:               "messagelevel-39-Error-fatal",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorFatal},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelFatal,
	},
	{
		name:               "messagelevel-40-Error-Panic",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorPanic},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-41-Error-Info-Panic",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorInfo, errorPanic},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelPanic,
	},
	{
		name:               "messagelevel-42-Error-Error-Warn",
		idLevelRanges:      IdLevelRanges,
		messageNumber:      1001,
		details:            []interface{}{123, "bob", errorError, errorWarn},
		expectedDefault:    logger.LevelDebug,
		expectedSenzingApi: logger.LevelError,
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
				IdLevelRanges: testCase.idLevelRanges,
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
				IdLevelRanges: testCase.idLevelRanges,
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
				IdLevelRanges: testCase.idLevelRanges,
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
	for key, value := range IdLevelRanges {
		idRangesStrings[key] = logger.LevelToTextMap[value]
	}

	for _, testCase := range testCases {
		test.Run(testCase.name+"-SenzingApi", func(test *testing.T) {
			testObject := &MessageLevelSenzingApi{
				IdLevelRanges: IdLevelRanges,
				IdStatuses:    idRangesStrings,
			}
			actual, err := testObject.MessageLevel(testCase.messageNumber, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, testCase.expectedSenzingApi, actual, testCase.name)
		})
	}
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
