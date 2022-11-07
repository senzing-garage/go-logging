package messagelogger

import (
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messagetext"
	"github.com/stretchr/testify/assert"
)

var idMessages = map[int]string{
	1: "%s knows %s",
	2: "%s does not know %s",
}

var messageFormat = &messageformat.MessageFormatJson{}
var messageText = &messagetext.MessageTextTemplated{
	IdMessages: idMessages,
}

var testCases = []struct {
	name              string
	productIdentifier int
	idMessages        map[int]string
	interfacesDefault []interface{}
	interfacesSenzing []interface{}
	messageNumber     int
	details           []interface{}
	expectedDefault   string
	expectedJson      string
}{
	{
		name:              "Test case: #1 - Info",
		productIdentifier: 9999,
		idMessages:        idMessages,
		messageNumber:     0,
		details:           []interface{}{"A", 1},
		expectedDefault:   `0: [map[1:"A" 2:1]]`,
		expectedJson:      `{"id":"senzing-99990000","status":"INFO","details":{"1":"A","2":1}}`,
	},
	{
		name:              "Test case: #2 - Warn",
		productIdentifier: 9999,
		idMessages:        idMessages,
		interfacesDefault: []interface{}{messageFormat},
		messageNumber:     1000,
		details:           []interface{}{"A", 1},
		expectedDefault:   `{"id":"1000","details":{"1":"A","2":1}}`,
		expectedJson:      `{"id":"senzing-99991000","status":"WARN","details":{"1":"A","2":1}}`,
	},
	{
		name:              "Test case: #10 - Warn",
		productIdentifier: 9999,
		idMessages:        idMessages,
		interfacesDefault: []interface{}{messageText},
		messageNumber:     1,
		details:           []interface{}{"Bob", "Jane"},
		expectedDefault:   `1: Bob knows Jane [map[1:"Bob" 2:"Jane"]]`,
		expectedJson:      `{"id":"senzing-99990001","status":"INFO","text":"Bob knows Jane","details":{"1":"Bob","2":"Jane"}}`,
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageLoggerInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions using New() - names begin with "Test"
// ----------------------------------------------------------------------------

// -- Test Message() method ---------------------------------------------------

func TestNew(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject, err := New(testCase.interfacesDefault...)
				testError(test, testObject, err)
				actual, err := testObject.Message(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedDefault, actual, testCase.name)
			})
		}
	}
}

// -- Test Log() method using New() -------------------------------------------

func TestDefaultLogMessage(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(1, "Bob was here.", "So was Jane.")
}

func TestDefaultLogMessageErrorLevels(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(1)
	testObject.Log(2, logger.LevelTrace)
	testObject.Log(3, logger.LevelDebug)
	testObject.Log(4, logger.LevelInfo)
	testObject.Log(5, logger.LevelWarn)
	testObject.Log(6, logger.LevelError)
	// testObject.Log(7, logger.LevelFatal)
	// testObject.Log(8, logger.LevelPanic)
}

func TestDefaultLogMessageWithMap(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(1001, "A couple", idMessages)
}

func TestDefaultLogMessageWithObject(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(2000, "An object", testObject)
}

// -- Test Log() method using New(...) ----------------------------------------

func TestNewLogMessageErrorLevels(test *testing.T) {
	testObject, err := New(logger.LevelWarn)
	testError(test, testObject, err)
	testObject.Log(1)
	testObject.Log(2, logger.LevelTrace)
	testObject.Log(3, logger.LevelDebug)
	testObject.Log(4, logger.LevelInfo)
	testObject.Log(5, logger.LevelWarn)
	testObject.Log(6, logger.LevelError)
	// testObject.Log(7, logger.LevelFatal)
	// testObject.Log(8, logger.LevelPanic)
}

func TestNewJsonFormatting(test *testing.T) {
	messageFormat := &messageformat.MessageFormatJson{}
	testObject, err := New(messageFormat)
	testError(test, testObject, err)
	testObject.Log(1, "Bob was here.", "So was Jane.")
}

func TestNewMessageTemplates(test *testing.T) {
	messageFormat := &messageformat.MessageFormatJson{}
	messageText := &messagetext.MessageTextTemplated{
		IdMessages: idMessages,
	}
	testObject, err := New(messageText, messageFormat)
	testError(test, testObject, err)
	testObject.Log(1, "Bob", "Jane", testObject)
	testObject.Log(2, "Bob", "Harry")
}

func TestNewBadInterfaces(test *testing.T) {
	expectedErrContains := "unsupported interfaces"
	messageFormat := &messageformat.MessageFormatJson{}
	messageText := &messagetext.MessageTextTemplated{
		IdMessages: idMessages,
	}
	messageLogger := &MessageLoggerDefault{}
	_, err := New(messageText, messageLogger, messageFormat, "ABC", 123)
	if assert.Error(test, err) {
		assert.ErrorContains(test, err, expectedErrContains)
	}
}

// -- Test IsXxxx method ------------------------------------------------------

func TestDefaultLogMessageIsXxxx(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	assert.False(test, testObject.IsTrace(), "Trace")
	assert.False(test, testObject.IsDebug(), "Debug")
	assert.True(test, testObject.IsInfo(), "Info")
	assert.True(test, testObject.IsWarn(), "Warn")
	assert.True(test, testObject.IsError(), "Error")
	assert.True(test, testObject.IsFatal(), "Fatal")
	assert.True(test, testObject.IsPanic(), "Panic")
}

func TestDefaultLogMessageIsXxxxForTraceLevel(test *testing.T) {
	testObject, err := New(logger.LevelTrace)
	testError(test, testObject, err)
	assert.True(test, testObject.IsTrace(), "Trace")
	assert.True(test, testObject.IsDebug(), "Debug")
	assert.True(test, testObject.IsInfo(), "Info")
	assert.True(test, testObject.IsWarn(), "Warn")
	assert.True(test, testObject.IsError(), "Error")
	assert.True(test, testObject.IsFatal(), "Fatal")
	assert.True(test, testObject.IsPanic(), "Panic")
}

func TestDefaultLogMessageIsXxxxForErrorLevel(test *testing.T) {
	testObject, err := New(logger.LevelError)
	testError(test, testObject, err)
	assert.False(test, testObject.IsTrace(), "Trace")
	assert.False(test, testObject.IsDebug(), "Debug")
	assert.False(test, testObject.IsInfo(), "Info")
	assert.False(test, testObject.IsWarn(), "Warn")
	assert.True(test, testObject.IsError(), "Error")
	assert.True(test, testObject.IsFatal(), "Fatal")
	assert.True(test, testObject.IsPanic(), "Panic")
}

// ----------------------------------------------------------------------------
// Test interface functions using NewSenzingLogger() - names begin with "Test"
// ----------------------------------------------------------------------------

// -- Test Message() method ---------------------------------------------------

func TestNewSenzingLogger(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedJson) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject, err := NewSenzingLogger(testCase.productIdentifier, testCase.idMessages, testCase.interfacesSenzing...)
				testError(test, testObject, err)
				actual, err := testObject.Message(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedJson, actual, testCase.name)
			})
		}
	}
}
