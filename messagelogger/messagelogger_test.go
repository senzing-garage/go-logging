package messagelogger

import (
	"errors"
	"testing"
	"time"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagedate"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messagelocation"
	"github.com/senzing/go-logging/messagetext"
	"github.com/senzing/go-logging/messagetime"
	"github.com/stretchr/testify/assert"
)

var messageFormat = &messageformat.MessageFormatJson{}

var messageDate = &messagedate.MessageDateStatic{
	Timestamp: getTimestamp(),
}
var messageTime = &messagetime.MessageTimeStatic{
	Timestamp: getTimestamp(),
}

var idMessages = map[int]string{
	2001: "%s knows %s",
	3001: "%s knows %s",
	4001: "%s knows %s",
	2:    "%s does not know %s",
}

var messageText = &messagetext.MessageTextTemplated{
	IdMessages: idMessages,
}
var messageLocation = &messagelocation.MessageLocationStatic{
	Location: "In AFunction() at somewhere.go:1234",
}

var testCasesForMessage = []struct {
	name                string
	componentIdentifier int
	idMessages          map[int]string
	interfacesDefault   []interface{}
	interfacesSenzing   []interface{}
	messageNumber       int
	details             []interface{}
	expectedDefault     string
	expectedJson        string
	expectedSenzing     string
}{
	{
		name:                "messagelogger-01-Trace",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       1,
		details:             []interface{}{"A", 1},
		expectedDefault:     `INFO 1: map[1:A 2:1]`,
		expectedJson:        `{"id":"senzing-99990001","status":"TRACE","details":{"1":"A","2":1}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"TRACE","id":"senzing-99990001","location":"In AFunction() at somewhere.go:1234","details":{"1":"A","2":1}}`,
	},
	{
		name:                "messagelogger-02-Debug",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       1001,
		details:             []interface{}{"A", 1},
		expectedDefault:     `INFO 1001: map[1:A 2:1]`,
		expectedJson:        `{"id":"senzing-99991001","status":"DEBUG","details":{"1":"A","2":1}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"DEBUG","id":"senzing-99991001","location":"In AFunction() at somewhere.go:1234","details":{"1":"A","2":1}}`,
	},
	{
		name:                "messagelogger-03-Info",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesDefault:   []interface{}{messageText},
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       2001,
		details:             []interface{}{"Bob", "Jane"},
		expectedDefault:     `INFO 2001: Bob knows Jane map[1:Bob 2:Jane]`,
		expectedJson:        `{"id":"senzing-99992001","status":"INFO","text":"Bob knows Jane","details":{"1":"Bob","2":"Jane"}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"INFO","id":"senzing-99992001","text":"Bob knows Jane","location":"In AFunction() at somewhere.go:1234","details":{"1":"Bob","2":"Jane"}}`,
	},
	{
		name:                "messagelogger-04-Warn",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       3001,
		details:             []interface{}{"Bob", "Jane"},
		expectedDefault:     `INFO 3001: map[1:Bob 2:Jane]`,
		expectedJson:        `{"id":"senzing-99993001","status":"WARN","text":"Bob knows Jane","details":{"1":"Bob","2":"Jane"}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"WARN","id":"senzing-99993001","text":"Bob knows Jane","location":"In AFunction() at somewhere.go:1234","details":{"1":"Bob","2":"Jane"}}`,
	},
	{
		name:                "messagelogger-05-Error",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       4001,
		details:             []interface{}{"Bob", "Jane"},
		expectedDefault:     `INFO 4001: map[1:Bob 2:Jane]`,
		expectedJson:        `{"id":"senzing-99940001","status":"ERROR","text":"Bob knows Jane","details":{"1":"Bob","2":"Jane"}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"ERROR","id":"senzing-99994001","text":"Bob knows Jane","location":"In AFunction() at somewhere.go:1234","details":{"1":"Bob","2":"Jane"}}`,
	},
	{
		name:                "messagelogger-20-Change-message-format",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesDefault:   []interface{}{messageFormat},
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       2002,
		details:             []interface{}{"A", 1},
		expectedDefault:     `{"level":"INFO","id":"2002","details":{"1":"A","2":1}}`,
		expectedJson:        `{"id":"senzing-99992002","status":"WARN","details":{"1":"A","2":1}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"INFO","id":"senzing-99992002","location":"In AFunction() at somewhere.go:1234","details":{"1":"A","2":1}}`,
	},
	{
		name:                "messagelogger-21-Include-error",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesDefault:   []interface{}{messageFormat},
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       2002,
		details:             []interface{}{"A", 1, errors.New("test error")},
		expectedDefault:     `{"level":"INFO","id":"2002","errors":[{"text":"test error"}],"details":{"1":"A","2":1}}`,
		expectedJson:        `{"id":"senzing-99992002","status":"WARN","details":{"1":"A","2":1}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"INFO","id":"senzing-99992002","location":"In AFunction() at somewhere.go:1234","errors":[{"text":"test error"}],"details":{"1":"A","2":1}}`,
	},
	{
		name:                "messagelogger-22-Include-error-JSON",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		interfacesDefault:   []interface{}{messageFormat},
		interfacesSenzing:   []interface{}{messageDate, messageTime, messageLocation},
		messageNumber:       2002,
		details:             []interface{}{"A", 1, errors.New(`{"error": "Bad error", "number": 1}`)},
		expectedDefault:     `{"level":"INFO","id":"2002","errors":[{"text":{"error":"Bad error","number":1}}],"details":{"1":"A","2":1}}`,
		expectedJson:        `{"id":"senzing-99992002","status":"WARN","details":{"1":"A","2":1}}`,
		expectedSenzing:     `{"date":"2000-01-01","time":"00:00:00.000000000","level":"INFO","id":"senzing-99992002","location":"In AFunction() at somewhere.go:1234","errors":[{"text":{"error":"Bad error","number":1}}],"details":{"1":"A","2":1}}`,
	},
}

var testCasesForIsMethods = []struct {
	name                string
	componentIdentifier int
	idMessages          map[int]string
	newLogLevel         logger.Level
	expectedTrace       bool
	expectedDebug       bool
	expectedInfo        bool
	expectedWarn        bool
	expectedError       bool
	expectedFatal       bool
	expectedPanic       bool
}{

	{
		name:                "Test case: #1 - Trace",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		newLogLevel:         logger.LevelTrace,
		expectedTrace:       true,
		expectedDebug:       true,
		expectedInfo:        true,
		expectedWarn:        true,
		expectedError:       true,
		expectedFatal:       true,
		expectedPanic:       true,
	},
	{
		name:                "Test case: #2 - Debug",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		newLogLevel:         logger.LevelDebug,
		expectedTrace:       false,
		expectedDebug:       true,
		expectedInfo:        true,
		expectedWarn:        true,
		expectedError:       true,
		expectedFatal:       true,
		expectedPanic:       true,
	},
	{
		name:                "Test case: #3 - Info",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		newLogLevel:         logger.LevelInfo,
		expectedTrace:       false,
		expectedDebug:       false,
		expectedInfo:        true,
		expectedWarn:        true,
		expectedError:       true,
		expectedFatal:       true,
		expectedPanic:       true,
	},
	{
		name:                "Test case: #4 - Warn",
		componentIdentifier: 9999,
		idMessages:          idMessages,
		newLogLevel:         logger.LevelWarn,
		expectedTrace:       false,
		expectedDebug:       false,
		expectedInfo:        false,
		expectedWarn:        true,
		expectedError:       true,
		expectedFatal:       true,
		expectedPanic:       true,
	},
	{
		name:                "Test case: #5 - Error",
		newLogLevel:         logger.LevelError,
		componentIdentifier: 9999,
		idMessages:          idMessages,
		expectedTrace:       false,
		expectedDebug:       false,
		expectedInfo:        false,
		expectedWarn:        false,
		expectedError:       true,
		expectedFatal:       true,
		expectedPanic:       true,
	},
	{
		name:                "Test case: #5 - Fatal",
		newLogLevel:         logger.LevelFatal,
		componentIdentifier: 9999,
		idMessages:          idMessages,
		expectedTrace:       false,
		expectedDebug:       false,
		expectedInfo:        false,
		expectedWarn:        false,
		expectedError:       false,
		expectedFatal:       true,
		expectedPanic:       true,
	},
	{
		name:                "Test case: #6 - Panic",
		newLogLevel:         logger.LevelPanic,
		componentIdentifier: 9999,
		idMessages:          idMessages,
		expectedTrace:       false,
		expectedDebug:       false,
		expectedInfo:        false,
		expectedWarn:        false,
		expectedError:       false,
		expectedFatal:       false,
		expectedPanic:       true,
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

func getTimestamp() time.Time {
	return time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
}

// ----------------------------------------------------------------------------
// Test interface functions using New()
// ----------------------------------------------------------------------------

// -- Test Message() method ---------------------------------------------------

func TestMessageLoggerNew(test *testing.T) {
	for _, testCase := range testCasesForMessage {
		if len(testCase.expectedDefault) > 0 {
			test.Run(testCase.name+"-Default", func(test *testing.T) {
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

func TestMessageLoggerNewLogMessage(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(1, "Bob was here.", "So was Jane.")
}

func TestMessageLoggerNewErrorLevels(test *testing.T) {
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

func TestMessageLoggerNewLogMessageWithMap(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(1001, "A couple", idMessages)
}

func TestMessageLoggerNewLogMessageWithObjects(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(2000, "An object", testObject, nil)
}

// -- Test Log() method using New(...) ----------------------------------------

func TestMessageLoggerNewLogMessageWithWarningLevel(test *testing.T) {
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

func TestMessageLoggerNewJsonFormatting(test *testing.T) {
	messageFormat := &messageformat.MessageFormatJson{}
	testObject, err := New(messageFormat)
	testError(test, testObject, err)
	testObject.Log(1, "Bob was here.", "So was Jane.")
}

func TestMessageLoggerNewMessageTemplates(test *testing.T) {
	messageFormat := &messageformat.MessageFormatJson{}
	messageText := &messagetext.MessageTextTemplated{
		IdMessages: idMessages,
	}
	testObject, err := New(messageText, messageFormat)
	testError(test, testObject, err)
	testObject.Log(1, "Bob", "Harry")
	testObject.Log(2, "Bob", "Jane", testObject, nil)
}

func TestMessageLoggerNewBadInterfaces(test *testing.T) {
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

func TestMessageLoggerNewIsMethods(test *testing.T) {
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

func TestMessageLoggerNewIsMethodDefault(test *testing.T) {
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
// Test interface functions using NewSenzingLogger()
// ----------------------------------------------------------------------------

// -- Test Message() method ---------------------------------------------------

func TestMessageLoggerNewSenzingLogger(test *testing.T) {
	for _, testCase := range testCasesForMessage {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name+"-Senzing", func(test *testing.T) {
				testObject, err := NewSenzingLogger(testCase.componentIdentifier, testCase.idMessages, testCase.interfacesSenzing...)
				testError(test, testObject, err)
				actual, err := testObject.Message(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}

// -- Test IsXxxx method ------------------------------------------------------

func TestMessageLoggerNewSenzingLoggerIsMethods(test *testing.T) {
	for _, testCase := range testCasesForIsMethods {
		test.Run(testCase.name, func(test *testing.T) {
			testObject, err := NewSenzingLogger(testCase.componentIdentifier, testCase.idMessages, testCase.newLogLevel)
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

func TestMessageLoggerNewSenzingLoggerIsMethodDefault(test *testing.T) {
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

// ----------------------------------------------------------------------------
// Test interface functions using NewSenzingApiLogger()
// ----------------------------------------------------------------------------

// -- Test Message() method ---------------------------------------------------

func TestMessageLoggerNewSenzingApiLogger(test *testing.T) {
	for _, testCase := range testCasesForMessage {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name+"-Senzing", func(test *testing.T) {
				testObject, err := NewSenzingApiLogger(testCase.componentIdentifier, testCase.idMessages, nil, testCase.interfacesSenzing...)
				testError(test, testObject, err)
				actual, err := testObject.Message(testCase.messageNumber, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}
