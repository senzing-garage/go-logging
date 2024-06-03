package logging

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var idMessagesTest = map[int]string{
	0001: "TRACE: %s works with %s",
	1001: "DEBUG: %s works with %s",
	2001: "INFO: %s works with %s",
	3001: "WARN: %s works with %s",
	4001: "ERROR: %s works with %s",
	5001: "FATAL: %s works with %s",
	6001: "PANIC: %s works with %s",
}

var testCasesForMessage = []struct {
	name                          string
	messageNumber                 int
	options                       []interface{}
	details                       []interface{}
	expectedNew                   string
	expectedNewSenzingSdkLogger   string
	expectedNewSenzingToolsLogger string
}{
	{
		name:                          "logging-0001",
		messageNumber:                 1,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   ``,
		expectedNewSenzingSdkLogger:   ``,
		expectedNewSenzingToolsLogger: ``,
	},
	{
		name:                          "logging-1001",
		messageNumber:                 1001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   ``,
		expectedNewSenzingSdkLogger:   ``,
		expectedNewSenzingToolsLogger: ``,
	},
	{
		name:                          "logging-2001",
		messageNumber:                 2001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"INFO","text":"INFO: Bob works with Jane","id":"2001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingSdkLogger:   `{"level":"INFO","text":"INFO: Bob works with Jane","id":"senzing-99972001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"INFO","text":"INFO: Bob works with Jane","id":"senzing-99972001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                          "logging-2002",
		messageNumber:                 2002,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden(), getOptionLogLevel("WARN")},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   ``,
		expectedNewSenzingSdkLogger:   ``,
		expectedNewSenzingToolsLogger: ``,
	},
	{
		name:                          "logging-3001",
		messageNumber:                 3001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"WARN","text":"WARN: Bob works with Jane","id":"3001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingSdkLogger:   `{"level":"WARN","text":"WARN: Bob works with Jane","id":"senzing-99973001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"WARN","text":"WARN: Bob works with Jane","id":"senzing-99973001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                          "logging-3002",
		messageNumber:                 3002,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden(), getOptionLogLevel("ERROR")},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   ``,
		expectedNewSenzingSdkLogger:   ``,
		expectedNewSenzingToolsLogger: ``,
	},
	{
		name:                          "logging-4001",
		messageNumber:                 4001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"4001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingSdkLogger:   `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"senzing-99974001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"senzing-99974001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                          "logging-5001",
		messageNumber:                 5001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"5001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingSdkLogger:   `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"senzing-99975001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"senzing-99975001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                          "logging-6001",
		messageNumber:                 6001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"6001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingSdkLogger:   `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"senzing-99976001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"senzing-99976001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
}

var testCasesForIsValidLogLevelName = []struct {
	name         string
	logLevelName string
	expected     bool
}{
	{
		name:         "valid-" + LevelTraceName,
		logLevelName: LevelTraceName,
		expected:     true,
	},
	{
		name:         "valid-" + LevelDebugName,
		logLevelName: LevelDebugName,
		expected:     true,
	},
	{
		name:         "valid-" + LevelInfoName,
		logLevelName: LevelInfoName,
		expected:     true,
	},
	{
		name:         "valid-" + LevelWarnName,
		logLevelName: LevelWarnName,
		expected:     true,
	},
	{
		name:         "valid-" + LevelErrorName,
		logLevelName: LevelErrorName,
		expected:     true,
	},
	{
		name:         "valid-" + LevelFatalName,
		logLevelName: LevelFatalName,
		expected:     true,
	},
	{
		name:         "valid-" + LevelPanicName,
		logLevelName: LevelPanicName,
		expected:     true,
	},
	{
		name:         "valid-bad-XYZZY",
		logLevelName: "XYZZY",
		expected:     false,
	},
}

var testCasesForIsXxxx = []struct {
	name          string
	logLevelName  string
	expectedTrace bool
	expectedDebug bool
	expectedInfo  bool
	expectedWarn  bool
	expectedError bool
	expectedFatal bool
	expectedPanic bool
}{
	{
		name:          "is-" + LevelTraceName,
		logLevelName:  LevelTraceName,
		expectedTrace: true,
		expectedDebug: true,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + LevelDebugName,
		logLevelName:  LevelDebugName,
		expectedTrace: false,
		expectedDebug: true,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + LevelInfoName,
		logLevelName:  LevelInfoName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + LevelWarnName,
		logLevelName:  LevelWarnName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + LevelErrorName,
		logLevelName:  LevelErrorName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + LevelFatalName,
		logLevelName:  LevelFatalName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: false,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + LevelPanicName,
		logLevelName:  LevelPanicName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: false,
		expectedFatal: false,
		expectedPanic: true,
	},
}

var (
	componentId  int           = 9997
	outputString *bytes.Buffer = new(bytes.Buffer)
)

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getOptionCallerSkip() *OptionCallerSkip {
	return &OptionCallerSkip{
		Value: 0,
	}
}

func getOptionIdMessages() *OptionIDMessages {
	return &OptionIDMessages{
		Value: idMessagesTest,
	}
}

func getOptionLogLevel(logLevelName string) *OptionLogLevel {
	return &OptionLogLevel{
		Value: logLevelName,
	}
}

func getOptionOutput() *OptionOutput {
	return &OptionOutput{
		Value: outputString,
	}
}

func getOptionTimeHidden() *OptionTimeHidden {
	return &OptionTimeHidden{
		Value: true,
	}
}

func testError(test *testing.T, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestLoggingImpl_IsValidLogLevelName(test *testing.T) {
	for _, testCase := range testCasesForIsValidLogLevelName {
		test.Run(testCase.name, func(test *testing.T) {
			actual := IsValidLogLevelName(testCase.logLevelName)
			assert.Equal(test, testCase.expected, actual, testCase.name)
		})
	}
}

func TestLoggingImpl_IsXxxx(test *testing.T) {
	for _, testCase := range testCasesForIsXxxx {
		test.Run(testCase.name, func(test *testing.T) {
			testObject, _ := New()
			testObject.SetLogLevel(testCase.logLevelName)
			assert.Equal(test, testCase.expectedTrace, testObject.IsTrace(), testCase.name)
			assert.Equal(test, testCase.expectedDebug, testObject.IsDebug(), testCase.name)
			assert.Equal(test, testCase.expectedInfo, testObject.IsInfo(), testCase.name)
			assert.Equal(test, testCase.expectedWarn, testObject.IsWarn(), testCase.name)
			assert.Equal(test, testCase.expectedError, testObject.IsError(), testCase.name)
			assert.Equal(test, testCase.expectedFatal, testObject.IsFatal(), testCase.name)
			assert.Equal(test, testCase.expectedPanic, testObject.IsPanic(), testCase.name)
		})
	}
}

func TestLoggingImpl_New(test *testing.T) {
	outputString.Reset()
	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-New", func(test *testing.T) {
			testObject, err := New(testCase.options...)
			testError(test, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNew, outputString.String(), testCase.name)
			outputString.Reset()
		})
	}
}

func TestLoggingImpl_NewSenzingToolsLogger(test *testing.T) {
	outputString.Reset()
	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-NewSenzingToolsLogger", func(test *testing.T) {
			testObject, err := NewSenzingToolsLogger(componentId, idMessagesTest, testCase.options...)
			testError(test, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNewSenzingToolsLogger, outputString.String(), testCase.name)
			outputString.Reset()
		})
	}
}

func TestLoggingImpl_NewSenzingSdkLogger(test *testing.T) {
	outputString.Reset()
	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-NewSenzingSdkLogger", func(test *testing.T) {
			testObject, err := NewSenzingSdkLogger(componentId, idMessagesTest, testCase.options...)
			testError(test, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNewSenzingSdkLogger, outputString.String(), testCase.name)
			outputString.Reset()
		})
	}
}
