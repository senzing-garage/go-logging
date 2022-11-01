package messagelogger

import (
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messagetext"
	"github.com/stretchr/testify/assert"
)

// const MessageIdFormat = "senzing-9999%04d"

// const printResults = 1

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

// func printResult(test *testing.T, title string, result interface{}) {
// 	if printResults == 1 {
// 		test.Logf("%s: %v", title, fmt.Sprintf("%v", result))
// 	}
// }

// func printActual(test *testing.T, actual interface{}) {
// 	printResult(test, "Actual", actual)
// }

func testError(test *testing.T, testObject MessageLoggerInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

// -- Log using New() with defaults ---------------------------------------------

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
	stringMap := map[string]string{
		"Husband": "Bob",
		"Wife":    "Jane",
	}
	testObject.Log(1001, "A couple", stringMap)
}

func TestDefaultLogMessageWithObject(test *testing.T) {
	testObject, err := New()
	testError(test, testObject, err)
	testObject.Log(2000, "An object", testObject)
}

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

// -- Log using New() with defaults ---------------------------------------------

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
		IdMessages: map[int]string{
			1: "%s knows %s",
			2: "%s does not know %s",
		},
	}
	testObject, err := New(messageText, messageFormat)
	testError(test, testObject, err)
	testObject.Log(1, "Bob", "Jane", testObject)
	testObject.Log(2, "Bob", "Harry")
}

func TestNewBadInterfaces(test *testing.T) {
	messageFormat := &messageformat.MessageFormatJson{}
	messageText := &messagetext.MessageTextTemplated{
		IdMessages: map[int]string{
			1: "%s knows %s",
			2: "%s does not know %s",
		},
	}
	messageLogger := &MessageLoggerDefault{}
	testObject, err := New(messageText, messageLogger, messageFormat, "ABC", 123)
	testError(test, testObject, err)
	testObject.Log(1, "Bob", "Jane", testObject)
	testObject.Log(2, "Bob", "Harry", err)
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
