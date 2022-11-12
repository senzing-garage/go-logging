package messagestatus

import (
	"errors"
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

var IdRanges = map[int]string{
	0000: logger.LevelTraceName,
	1000: logger.LevelDebugName,
	2000: logger.LevelInfoName,
	3000: logger.LevelWarnName,
	4000: logger.LevelErrorName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}

var testCases = []struct {
	name               string
	idRanges           map[int]string
	IdStatuses         map[int]string
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
		idRanges:           IdRanges,
		IdStatuses:         IdRanges,
		details:            []interface{}{"A", 1},
		expectedById:       logger.LevelTraceName,
		expectedByIdRange:  logger.LevelTraceName,
		expectedDefault:    logger.LevelTraceName,
		expectedSenzingApi: logger.LevelTraceName,
	},
	{
		name:               "messagestatus-02-Debug",
		messageNumber:      1000,
		idRanges:           IdRanges,
		IdStatuses:         IdRanges,
		details:            []interface{}{"A", 1},
		expectedById:       logger.LevelDebugName,
		expectedByIdRange:  logger.LevelDebugName,
		expectedDefault:    logger.LevelDebugName,
		expectedSenzingApi: logger.LevelDebugName,
	},
	{
		name:               "messagestatus-03-Info",
		messageNumber:      2000,
		idRanges:           IdRanges,
		IdStatuses:         IdRanges,
		details:            []interface{}{"A", 1},
		expectedById:       logger.LevelInfoName,
		expectedByIdRange:  logger.LevelInfoName,
		expectedDefault:    logger.LevelInfoName,
		expectedSenzingApi: logger.LevelInfoName,
	},
	{
		name:               "messagestatus-04-Warn",
		messageNumber:      3000,
		idRanges:           IdRanges,
		IdStatuses:         IdRanges,
		details:            []interface{}{"A", 1},
		expectedById:       logger.LevelWarnName,
		expectedByIdRange:  logger.LevelWarnName,
		expectedDefault:    logger.LevelWarnName,
		expectedSenzingApi: logger.LevelWarnName,
	},
	{
		name:               "messagestatus-05-Error",
		messageNumber:      4000,
		idRanges:           IdRanges,
		IdStatuses:         IdRanges,
		details:            []interface{}{"A", 1},
		expectedById:       logger.LevelErrorName,
		expectedByIdRange:  logger.LevelErrorName,
		expectedDefault:    logger.LevelErrorName,
		expectedSenzingApi: logger.LevelErrorName,
	},
	{
		name:               "messagestatus-06-Fatal",
		messageNumber:      5000,
		idRanges:           IdRanges,
		IdStatuses:         IdRanges,
		details:            []interface{}{"A", 1},
		expectedById:       logger.LevelFatalName,
		expectedByIdRange:  logger.LevelFatalName,
		expectedDefault:    logger.LevelFatalName,
		expectedSenzingApi: logger.LevelFatalName,
	},
	{
		name:               "messagestatus-07-Panic",
		messageNumber:      6000,
		idRanges:           IdRanges,
		IdStatuses:         IdRanges,
		details:            []interface{}{"A", 1},
		expectedById:       logger.LevelPanicName,
		expectedByIdRange:  logger.LevelPanicName,
		expectedDefault:    logger.LevelPanicName,
		expectedSenzingApi: logger.LevelPanicName,
	},
	{
		name:              "messagestatus-11-Trace",
		messageNumber:     1,
		idRanges:          IdRanges,
		IdStatuses:        IdRanges,
		details:           []interface{}{"A", 1},
		expectedByIdRange: logger.LevelTraceName,
	},
	{
		name:              "messagestatus-12-Debug",
		messageNumber:     1001,
		idRanges:          IdRanges,
		IdStatuses:        IdRanges,
		details:           []interface{}{"A", 1},
		expectedByIdRange: logger.LevelDebugName,
	},
	{
		name:              "messagestatus-13-Info",
		messageNumber:     2001,
		idRanges:          IdRanges,
		IdStatuses:        IdRanges,
		details:           []interface{}{"A", 1},
		expectedByIdRange: logger.LevelInfoName,
	},
	{
		name:              "messagestatus-14-Warn",
		messageNumber:     3001,
		idRanges:          IdRanges,
		IdStatuses:        IdRanges,
		details:           []interface{}{"A", 1},
		expectedByIdRange: logger.LevelWarnName,
	},
	{
		name:              "messagestatus-15-Error",
		messageNumber:     4001,
		idRanges:          IdRanges,
		IdStatuses:        IdRanges,
		details:           []interface{}{"A", 1},
		expectedByIdRange: logger.LevelErrorName,
	},
	{
		name:              "messagestatus-16-Fatal",
		messageNumber:     5001,
		idRanges:          IdRanges,
		IdStatuses:        IdRanges,
		details:           []interface{}{"A", 1},
		expectedByIdRange: logger.LevelFatalName,
	},
	{
		name:              "messagestatus-17-Panic",
		messageNumber:     6001,
		idRanges:          IdRanges,
		IdStatuses:        IdRanges,
		details:           []interface{}{"A", 1},
		expectedByIdRange: logger.LevelPanicName,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageStatusInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusById
// ----------------------------------------------------------------------------

func TestMessageStatusById(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name+"-ById", func(test *testing.T) {
			testObject := &MessageStatusById{
				IdStatuses: testCase.idRanges,
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
				IdRanges: testCase.idRanges,
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
				IdStatuses: testCase.idRanges,
			}
			actual, _ := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedSenzingApi, actual, testCase.name)
		})
	}
}

func TestMessageStatusSenzingApiWith0037E(test *testing.T) {
	expected := "ERROR_bad_user_input"
	anError := errors.New("0037E|Unknown resolved entity value '2'")
	testObject := &MessageStatusSenzingApi{}
	actual, err := testObject.MessageStatus(1, "A", 1, testObject, anError)
	testError(test, testObject, err)
	assert.Equal(test, expected, actual)
}

func TestMessageStatusSenzingApiWithSenzingApiError(test *testing.T) {
	expected := "ERROR_bad_user_input"
	anError := errors.New("0037E|Unknown resolved entity value")
	testObject := &MessageStatusSenzingApi{
		IdRanges: map[int]string{
			0: logger.LevelInfoName,
		},
	}
	actual, err := testObject.MessageStatus(1, "A", 1, testObject, anError)
	testError(test, testObject, err)
	assert.Equal(test, expected, actual)
}

func TestMessageStatusSenzingApiWith2SenzingApiError2(test *testing.T) {
	expected := "ERROR_unrecoverable"
	anError1 := errors.New("0019E|Configuration not found")
	anError2 := errors.New("0099E|Made up error")
	testObject := &MessageStatusSenzingApi{
		IdRanges: map[int]string{
			0: logger.LevelInfoName,
		},
	}
	actual, err := testObject.MessageStatus(1, "A", 1, testObject, anError1, anError2)
	testError(test, testObject, err)
	assert.Equal(test, expected, actual)
}

func TestMessageStatusSenzingApiWithUnknownError(test *testing.T) {
	expected := logger.LevelWarnName
	anError := errors.New("1234E|Made up error")
	testObject := &MessageStatusSenzingApi{
		IdRanges: map[int]string{
			0:    logger.LevelInfoName,
			1000: logger.LevelWarnName,
		},
	}
	actual, err := testObject.MessageStatus(1000, "A", 1, testObject, anError)
	testError(test, testObject, err)
	assert.Equal(test, expected, actual)
}
