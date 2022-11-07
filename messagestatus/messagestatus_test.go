package messagestatus

import (
	"errors"
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

var IdRanges = map[int]string{
	0000: logger.LevelInfoName,
	1000: logger.LevelWarnName,
	2000: logger.LevelErrorName,
	3000: logger.LevelDebugName,
	4000: logger.LevelTraceName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}

var testCases = []struct {
	name            string
	idRanges        map[int]string
	IdStatuses      map[int]string
	messageNumber   int
	details         []interface{}
	expectedDefault string
}{
	{
		name:            "Test case: #1 - Info",
		messageNumber:   0,
		idRanges:        IdRanges,
		IdStatuses:      IdRanges,
		details:         []interface{}{"A", 1},
		expectedDefault: logger.LevelInfoName,
	},
	{
		name:            "Test case: #2 - Warn",
		messageNumber:   1000,
		idRanges:        IdRanges,
		IdStatuses:      IdRanges,
		details:         []interface{}{"A", 1},
		expectedDefault: logger.LevelWarnName,
	},
	{
		name:            "Test case: #3 - Error",
		messageNumber:   2000,
		idRanges:        IdRanges,
		IdStatuses:      IdRanges,
		details:         []interface{}{"A", 1},
		expectedDefault: logger.LevelErrorName,
	},
	{
		name:            "Test case: #4 - Debug",
		messageNumber:   3000,
		idRanges:        IdRanges,
		IdStatuses:      IdRanges,
		details:         []interface{}{"A", 1},
		expectedDefault: logger.LevelDebugName,
	},
	{
		name:            "Test case: #5 - Trace",
		messageNumber:   4000,
		idRanges:        IdRanges,
		IdStatuses:      IdRanges,
		details:         []interface{}{"A", 1},
		expectedDefault: logger.LevelTraceName,
	},
	{
		name:            "Test case: #6 - Fatal",
		messageNumber:   5000,
		idRanges:        IdRanges,
		IdStatuses:      IdRanges,
		details:         []interface{}{"A", 1},
		expectedDefault: logger.LevelFatalName,
	},
	{
		name:            "Test case: #7 - Panic",
		messageNumber:   6000,
		idRanges:        IdRanges,
		IdStatuses:      IdRanges,
		details:         []interface{}{"A", 1},
		expectedDefault: logger.LevelPanicName,
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
// Test interface functions for MessageStatusById - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageStatusById(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageStatusById{
					IdStatuses: testCase.idRanges,
				}
				actual, err := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusByIdRange - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageStatusByIdRange(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageStatusByIdRange{
					IdRanges: testCase.idRanges,
				}
				actual, err := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusSenzing - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageStatusSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageStatusSenzing{
					IdRanges: testCase.idRanges,
				}
				actual, err := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageStatusSenzingApi - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageStatusSenzingApi(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageStatusSenzingApi{
					IdRanges:   testCase.idRanges,
					IdStatuses: testCase.idRanges,
				}
				actual, err := testObject.MessageStatus(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
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
