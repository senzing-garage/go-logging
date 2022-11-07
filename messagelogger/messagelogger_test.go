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

var testCasesForMessage = []struct {
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
		expectedDefault:   `0: [map[1:A 2:1]]`,
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
		name:              "Test case: #3 - Warn",
		productIdentifier: 9999,
		idMessages:        idMessages,
		interfacesDefault: []interface{}{messageText},
		messageNumber:     1,
		details:           []interface{}{"Bob", "Jane"},
		expectedDefault:   `1: Bob knows Jane [map[1:Bob 2:Jane]]`,
		expectedJson:      `{"id":"senzing-99990001","status":"INFO","text":"Bob knows Jane","details":{"1":"Bob","2":"Jane"}}`,
	},
}

var testCasesForIsMethods = []struct {
	name              string
	productIdentifier int
	idMessages        map[int]string
	newLogLevel       logger.Level
	expectedTrace     bool
	expectedDebug     bool
	expectedInfo      bool
	expectedWarn      bool
	expectedError     bool
	expectedFatal     bool
	expectedPanic     bool
}{

	{
		name:              "Test case: #1 - Trace",
		productIdentifier: 9999,
		idMessages:        idMessages,
		newLogLevel:       logger.LevelTrace,
		expectedTrace:     true,
		expectedDebug:     true,
		expectedInfo:      true,
		expectedWarn:      true,
		expectedError:     true,
		expectedFatal:     true,
		expectedPanic:     true,
	},
	{
		name:              "Test case: #2 - Debug",
		productIdentifier: 9999,
		idMessages:        idMessages,
		newLogLevel:       logger.LevelDebug,
		expectedTrace:     false,
		expectedDebug:     true,
		expectedInfo:      true,
		expectedWarn:      true,
		expectedError:     true,
		expectedFatal:     true,
		expectedPanic:     true,
	},
	{
		name:              "Test case: #3 - Info",
		productIdentifier: 9999,
		idMessages:        idMessages,
		newLogLevel:       logger.LevelInfo,
		expectedTrace:     false,
		expectedDebug:     false,
		expectedInfo:      true,
		expectedWarn:      true,
		expectedError:     true,
		expectedFatal:     true,
		expectedPanic:     true,
	},
	{
		name:              "Test case: #4 - Warn",
		productIdentifier: 9999,
		idMessages:        idMessages,
		newLogLevel:       logger.LevelWarn,
		expectedTrace:     false,
		expectedDebug:     false,
		expectedInfo:      false,
		expectedWarn:      true,
		expectedError:     true,
		expectedFatal:     true,
		expectedPanic:     true,
	},
	{
		name:              "Test case: #5 - Error",
		newLogLevel:       logger.LevelError,
		productIdentifier: 9999,
		idMessages:        idMessages,
		expectedTrace:     false,
		expectedDebug:     false,
		expectedInfo:      false,
		expectedWarn:      false,
		expectedError:     true,
		expectedFatal:     true,
		expectedPanic:     true,
	},
	{
		name:              "Test case: #5 - Fatal",
		newLogLevel:       logger.LevelFatal,
		productIdentifier: 9999,
		idMessages:        idMessages,
		expectedTrace:     false,
		expectedDebug:     false,
		expectedInfo:      false,
		expectedWarn:      false,
		expectedError:     false,
		expectedFatal:     true,
		expectedPanic:     true,
	},
	{
		name:              "Test case: #6 - Panic",
		newLogLevel:       logger.LevelPanic,
		productIdentifier: 9999,
		idMessages:        idMessages,
		expectedTrace:     false,
		expectedDebug:     false,
		expectedInfo:      false,
		expectedWarn:      false,
		expectedError:     false,
		expectedFatal:     false,
		expectedPanic:     true,
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

func TestDefault(test *testing.T) {
	for _, testCase := range testCasesForMessage {
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

func TestDefaultErrorLevels(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(1, logger.LevelTrace)
	testObject.Log(2, logger.LevelDebug)
	testObject.Log(3, logger.LevelInfo)
	testObject.Log(4, logger.LevelWarn)
	testObject.Log(5, logger.LevelError)
	// testObject.Log(6, logger.LevelFatal)
	// testObject.Log(7, logger.LevelPanic)
}

func TestDefaultLogMessageWithMap(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(1001, "A couple", idMessages)
}

func TestDefaultLogMessageWithObjects(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(2000, "An object", testObject, nil)
}

// -- Test Log() method using New(...) ----------------------------------------

func TestNewLogMessageWithWarningLevel(test *testing.T) {
	testObject, err := New(logger.LevelWarn)
	testError(test, testObject, err)
	testObject.Log(1, logger.LevelTrace)
	testObject.Log(2, logger.LevelDebug)
	testObject.Log(3, logger.LevelInfo)
	testObject.Log(4, logger.LevelWarn)
	testObject.Log(5, logger.LevelError)
	// testObject.Log(6, logger.LevelFatal)
	// testObject.Log(7, logger.LevelPanic)
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
	testObject.Log(1, "Bob", "Jane", testObject, nil)
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

func TestNewIsMethods(test *testing.T) {
	for _, testCase := range testCasesForIsMethods {
		test.Run(testCase.name, func(test *testing.T) {
			testObject, err := New(testCase.newLogLevel)
			testError(test, testObject, err)
			assert.Equal(test, testCase.expectedTrace, testObject.IsTrace(), "Trace")
			assert.Equal(test, testCase.expectedDebug, testObject.IsDebug(), "Debug")
			assert.Equal(test, testCase.expectedInfo, testObject.IsInfo(), "Info")
			assert.Equal(test, testCase.expectedWarn, testObject.IsWarn(), "Warn")
			assert.Equal(test, testCase.expectedError, testObject.IsError(), "Error")
			assert.Equal(test, testCase.expectedFatal, testObject.IsFatal(), "Fatal")
			assert.Equal(test, testCase.expectedPanic, testObject.IsPanic(), "Panic")
		})
	}
}

func TestNewIsMethodDefault(test *testing.T) {
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

// ----------------------------------------------------------------------------
// Test interface functions using NewSenzingLogger() - names begin with "Test"
// ----------------------------------------------------------------------------

// -- Test Message() method ---------------------------------------------------

func TestNewSenzingLogger(test *testing.T) {
	for _, testCase := range testCasesForMessage {
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

// -- Test IsXxxx method ------------------------------------------------------

func TestNewSenzingLoggerIsMethods(test *testing.T) {
	for _, testCase := range testCasesForIsMethods {
		test.Run(testCase.name, func(test *testing.T) {
			testObject, err := NewSenzingLogger(testCase.productIdentifier, testCase.idMessages, testCase.newLogLevel)
			testError(test, testObject, err)
			assert.Equal(test, testCase.expectedTrace, testObject.IsTrace(), "Trace")
			assert.Equal(test, testCase.expectedDebug, testObject.IsDebug(), "Debug")
			assert.Equal(test, testCase.expectedInfo, testObject.IsInfo(), "Info")
			assert.Equal(test, testCase.expectedWarn, testObject.IsWarn(), "Warn")
			assert.Equal(test, testCase.expectedError, testObject.IsError(), "Error")
			assert.Equal(test, testCase.expectedFatal, testObject.IsFatal(), "Fatal")
			assert.Equal(test, testCase.expectedPanic, testObject.IsPanic(), "Panic")
		})
	}
}

func TestNewSenzingLoggerIsMethodDefault(test *testing.T) {
	// Should be same as logger.InfoLevel.
	testObject, err := NewSenzingLogger(9999, idMessages)
	testError(test, testObject, err)
	assert.False(test, testObject.IsTrace(), "Trace")
	assert.False(test, testObject.IsDebug(), "Debug")
	assert.True(test, testObject.IsInfo(), "Info")
	assert.True(test, testObject.IsWarn(), "Warn")
	assert.True(test, testObject.IsError(), "Error")
	assert.True(test, testObject.IsFatal(), "Fatal")
	assert.True(test, testObject.IsPanic(), "Panic")
}
