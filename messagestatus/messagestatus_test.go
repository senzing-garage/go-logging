package messagestatus

import (
	"errors"
	"testing"

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
	IdStatuses         map[int]string
	idStatusRanges     map[int]string
	messageNumber      int
	details            []interface{}
	expectedById       string
	expectedByIdRange  string
	expectedDefault    string
	expectedSenzing    string
	expectedSenzingApi string
}{
	{
		name:               "messagestatus-01-Trace",
		messageNumber:      0,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1},
		expectedById:       Trace,
		expectedByIdRange:  Trace,
		expectedDefault:    Trace,
		expectedSenzingApi: Trace,
	},
	{
		name:               "messagestatus-02-Debug",
		messageNumber:      1000,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1},
		expectedById:       Debug,
		expectedByIdRange:  Debug,
		expectedDefault:    Debug,
		expectedSenzingApi: Debug,
	},
	{
		name:               "messagestatus-03-Info",
		messageNumber:      2000,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1},
		expectedById:       Info,
		expectedByIdRange:  Info,
		expectedDefault:    Info,
		expectedSenzingApi: Info,
	},
	{
		name:               "messagestatus-04-Warn",
		messageNumber:      3000,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1},
		expectedById:       Warn,
		expectedByIdRange:  Warn,
		expectedDefault:    Warn,
		expectedSenzingApi: Warn,
	},
	{
		name:               "messagestatus-05-Error",
		messageNumber:      4000,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1},
		expectedById:       Error,
		expectedByIdRange:  Error,
		expectedDefault:    Error,
		expectedSenzingApi: Error,
	},
	{
		name:               "messagestatus-06-Fatal",
		messageNumber:      5000,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1},
		expectedById:       Fatal,
		expectedByIdRange:  Fatal,
		expectedDefault:    Fatal,
		expectedSenzingApi: Fatal,
	},
	{
		name:               "messagestatus-07-Panic",
		messageNumber:      6000,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1},
		expectedById:       Panic,
		expectedByIdRange:  Panic,
		expectedDefault:    Panic,
		expectedSenzingApi: Panic,
	},
	{
		name:              "messagestatus-11-Trace",
		messageNumber:     1,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1},
		expectedByIdRange: Trace,
	},
	{
		name:              "messagestatus-12-Debug",
		messageNumber:     1001,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1},
		expectedByIdRange: Debug,
	},
	{
		name:              "messagestatus-13-Info",
		messageNumber:     2001,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1},
		expectedByIdRange: Info,
	},
	{
		name:              "messagestatus-14-Warn",
		messageNumber:     3001,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1},
		expectedByIdRange: Warn,
	},
	{
		name:              "messagestatus-15-Error",
		messageNumber:     4001,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1},
		expectedByIdRange: Error,
	},
	{
		name:              "messagestatus-16-Fatal",
		messageNumber:     5001,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1},
		expectedByIdRange: Fatal,
	},
	{
		name:              "messagestatus-17-Panic",
		messageNumber:     6001,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1},
		expectedByIdRange: Panic,
	},
	{
		name:              "messagestatus-20-Error-unknown",
		messageNumber:     1001,
		IdStatuses:        IdLevelRangesAsString,
		idStatusRanges:    IdLevelRangesAsString,
		details:           []interface{}{"A", 1, errorUnknown},
		expectedByIdRange: Debug,
	},
	{
		name:               "messagestatus-21-Error-Trace",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorTrace},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Trace,
	},
	{
		name:               "messagestatus-22-Error-Debug",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorDebug},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Debug,
	},
	{
		name:               "messagestatus-22-Error-Info",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorInfo},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Info,
	},
	{
		name:               "messagestatus-23-Error-Warn",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorWarn},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Warn,
	},
	{
		name:               "messagestatus-23-Error-Error",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorError},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Error,
	},
	{
		name:               "messagestatus-24-Error-Retryable",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorRetryable},
		expectedByIdRange:  Debug,
		expectedSenzingApi: ErrorRetryable,
	},
	{
		name:               "messagestatus-25-Error-BadUserInput",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorBadUserInput},
		expectedByIdRange:  Debug,
		expectedSenzingApi: ErrorBadUserInput,
	},
	{
		name:               "messagestatus-26-Error-Unrecoverable",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorUnrecoverable},
		expectedByIdRange:  Debug,
		expectedSenzingApi: ErrorUnrecoverable,
	},
	{
		name:               "messagestatus-27-Error-Fatal",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorFatal},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Fatal,
	},
	{
		name:               "messagestatus-28-Error-Panic",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorPanic},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Panic,
	},
	{
		name:               "messagestatus-30-Error-Retryable-and-Unrecoverable",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorRetryable, errorUnrecoverable},
		expectedByIdRange:  Debug,
		expectedSenzingApi: ErrorUnrecoverable,
	},
	{
		name:               "messagestatus-31-Error-Unrecoverable-and-Retryable",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorUnrecoverable, errorRetryable},
		expectedByIdRange:  Debug,
		expectedSenzingApi: ErrorUnrecoverable,
	},
	{
		name:               "messagestatus-32-Error-Panic-Unrecoverable-and-Retryable",
		messageNumber:      1001,
		IdStatuses:         IdLevelRangesAsString,
		idStatusRanges:     IdLevelRangesAsString,
		details:            []interface{}{"A", 1, errorPanic, errorUnrecoverable, errorRetryable},
		expectedByIdRange:  Debug,
		expectedSenzingApi: Panic,
	},
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusById
// ----------------------------------------------------------------------------

func TestMessageStatusById(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-ById", func(test *testing.T) {
			testObject := &MessageStatusById{
				IdStatuses: testCase.idStatusRanges,
			}
			actual, _ := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedById, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusByIdRange
// ----------------------------------------------------------------------------

func TestMessageStatusByIdRange(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-ByIdRange", func(test *testing.T) {
			testObject := &MessageStatusByIdRange{
				IdStatusRanges: testCase.idStatusRanges,
			}
			actual, _ := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedByIdRange, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusSenzing
// ----------------------------------------------------------------------------

func TestMessageStatusSenzing(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-Senzing", func(test *testing.T) {
			testObject := &MessageStatusSenzing{}
			actual, _ := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusSenzingApi
// ----------------------------------------------------------------------------

func TestMessageStatusSenzingApi(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-SenzingApi", func(test *testing.T) {
			testObject := &MessageStatusSenzingApi{
				IdStatuses: testCase.IdStatuses,
			}
			actual, _ := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedSenzingApi, actual, testCase.name)
		})
	}
}
