package logging

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/senzing-garage/go-messaging/messenger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
)

var (
	badComponentIdentifier = 10000
	badIDMessages          map[int]string
	badIDStatuses          map[int]string
	badLogLevelName                      = "BadLogLevelName"
	messageIDTemplate                    = "test-%04d"
	testDuration           time.Duration = 10000
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

var idStatusesTest = map[int]string{
	2000: "SUCCESS",
	4000: "FAILURE",
	6000: "DISASTER",
}

var testCasesForMessage = []struct {
	name                     string
	messageNumber            int
	options                  []interface{}
	details                  []interface{}
	expectedNew              string
	expectedNewSenzingLogger string
}{
	{
		name:                     "logging-0001",
		messageNumber:            1,
		options:                  []interface{}{getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:                     "logging-1001",
		messageNumber:            1001,
		options:                  []interface{}{getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:                     "logging-2001",
		messageNumber:            2001,
		options:                  []interface{}{getMessageFields(), getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"INFO","text":"INFO: Bob works with Jane","id":"2001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"INFO","text":"INFO: Bob works with Jane","id":"SZTL99972001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                     "logging-2002",
		messageNumber:            2002,
		options:                  []interface{}{getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden(), getOptionLogLevel("WARN")},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:                     "logging-3001",
		messageNumber:            3001,
		options:                  []interface{}{getMessageFields(), getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"WARN","text":"WARN: Bob works with Jane","id":"3001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","text":"WARN: Bob works with Jane","id":"SZTL99973001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                     "logging-3002",
		messageNumber:            3002,
		options:                  []interface{}{getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden(), getOptionLogLevel("ERROR")},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:                     "logging-3003",
		messageNumber:            3003,
		options:                  []interface{}{getOptionIDStatuses(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{},
		expectedNew:              `{"level":"WARN","id":"3003"}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","id":"SZTL99973003"}` + "\n",
	},
	{
		name:                     "logging-3004",
		messageNumber:            3004,
		options:                  []interface{}{getOptionIDTemplate(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{},
		expectedNew:              `{"level":"WARN","id":"test-3004"}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","id":"test-3004"}` + "\n",
	},
	{
		name:                     "logging-3005",
		messageNumber:            3005,
		options:                  []interface{}{getMessageFields(), getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane", &MessageCode{Value: "Test Code"}, &MessageDuration{Value: 1234567}, &MessageID{Value: "3005-Test"}, &MessageLevel{Value: "ERROR"}, &MessageLocation{Value: "Test Location"}, &MessageReason{Value: "Test reason ..."}, &MessageStatus{Value: "Test-Status"}, &MessageText{Value: "Test text"}, &MessageTime{Value: time.Now()}, &OptionCallerSkip{Value: 3}, testDuration},
		expectedNew:              `{"level":"WARN","text":"Test text","id":"ERROR","code":"Test Code","reason":"Test reason ...","status":"Test-Status","duration":10000,"location":"In func1() at logging_test.go:371","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","text":"Test text","id":"ERROR","code":"Test Code","reason":"Test reason ...","status":"Test-Status","duration":10000,"location":"In func1() at logging_test.go:424","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                     "logging-3006",
		messageNumber:            3006,
		options:                  []interface{}{&OptionMessageField{Value: "id"}, &OptionCallerSkip{Value: 3}, getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane", &MessageCode{Value: "Test Code"}, &MessageDuration{Value: 1234567}, &MessageID{Value: "3005-Test"}, &MessageLevel{Value: "ERROR"}, &MessageLocation{Value: "Test Location"}, &MessageReason{Value: "Test reason ..."}, &MessageStatus{Value: "Test-Status"}, &MessageText{Value: "Test text"}, &MessageTime{Value: time.Now()}, &OptionCallerSkip{Value: 3}, testDuration},
		expectedNew:              `{"level":"WARN","text":"Test text","id":"ERROR"}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","id":"ERROR","reason":"Test reason ..."}` + "\n",
	},
	{
		name:                     "logging-4001",
		messageNumber:            4001,
		options:                  []interface{}{getMessageFields(), getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"4001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"SZTL99974001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                     "logging-5001",
		messageNumber:            5001,
		options:                  []interface{}{getMessageFields(), getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"5001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"SZTL99975001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:                     "logging-6001",
		messageNumber:            6001,
		options:                  []interface{}{getMessageFields(), getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"6001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"SZTL99976001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
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
	componentID  = 9997
	outputString = new(bytes.Buffer) // *bytes.Buffer
)

// ----------------------------------------------------------------------------
// Test interface methods
// ----------------------------------------------------------------------------

func TestBasicLogging_GetLogLevel(test *testing.T) {
	loggerOptions := []interface{}{}
	logger, err := NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)
	actual := logger.GetLogLevel()
	assert.Equal(test, "INFO", actual)
}

func TestBasicLogging_IsXxxx(test *testing.T) {
	for _, testCase := range testCasesForIsXxxx {
		test.Run(testCase.name, func(test *testing.T) {
			testObject, _ := New()
			err := testObject.SetLogLevel(testCase.logLevelName)
			require.NoError(test, err)
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

func TestBasicLogging_JSON(test *testing.T) {
	loggerOptions := []interface{}{}
	logger, err := NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)
	actual := logger.JSON(9999, "detail")
	assert.NotEmpty(test, len(actual))
}

func TestBasicLogging_NewError(test *testing.T) {
	loggerOptions := []interface{}{getOptionIDMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()}
	logger, err := NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)
	err = logger.NewError(4000, "A bad thing")
	require.Error(test, err)
}

func TestBasicLogging_SetLogLevel_badLogLevelName(test *testing.T) {
	loggerOptions := []interface{}{}
	logger, err := NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)
	err = logger.SetLogLevel(badLogLevelName)
	require.Error(test, err)
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestLogging_IsValidLogLevelName(test *testing.T) {
	for _, testCase := range testCasesForIsValidLogLevelName {
		test.Run(testCase.name, func(test *testing.T) {
			actual := IsValidLogLevelName(testCase.logLevelName)
			assert.Equal(test, testCase.expected, actual, testCase.name)
		})
	}
}

func TestLogging_New(test *testing.T) {
	outputString.Reset()
	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-New", func(test *testing.T) {
			testObject, err := New(testCase.options...)
			require.NoError(test, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNew, outputString.String(), testCase.name)
			outputString.Reset()
		})
	}
}

func TestLogging_New_badComponentIdentifier(test *testing.T) {
	options := []interface{}{
		&OptionComponentID{
			Value: badComponentIdentifier,
		},
	}
	_, err := New(options...)
	require.Error(test, err)
}

func TestLogging_New_badIDMessages(test *testing.T) {
	options := []interface{}{
		&OptionIDMessages{
			Value: badIDMessages,
		},
	}
	_, err := New(options...)
	require.Error(test, err)
}

func TestLogging_New_badIDStatuses(test *testing.T) {
	options := []interface{}{
		&OptionIDStatuses{
			Value: badIDStatuses,
		},
	}
	_, err := New(options...)
	require.Error(test, err)
}

func TestLogging_New_badLogLevelName(test *testing.T) {
	options := []interface{}{
		&OptionLogLevel{
			Value: badLogLevelName,
		},
	}
	_, err := New(options...)
	require.Error(test, err)
}

func TestLogging_NewSenzingLogger(test *testing.T) {
	outputString.Reset()
	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-NewSenzingLogger", func(test *testing.T) {
			testObject, err := NewSenzingLogger(componentID, idMessagesTest, testCase.options...)
			require.NoError(test, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNewSenzingLogger, outputString.String(), testCase.name)
			outputString.Reset()
		})
	}
}

// ----------------------------------------------------------------------------
// Test private method functions
// ----------------------------------------------------------------------------

func TestBasicLogging_initialize_badLeveler(test *testing.T) {
	messenger, err := messenger.New()
	require.NoError(test, err)
	var output io.Writer = os.Stderr
	var slogLeveler = new(slog.LevelVar)
	slogLeveler.Set(slog.LevelInfo)
	sLogger := slog.New(slog.NewJSONHandler(output, SlogHandlerOptions(slogLeveler)))
	logger := &BasicLogging{
		Ctx:       context.TODO(),
		logger:    sLogger,
		messenger: messenger,
	}
	assert.Panics(test, func() { _ = logger.initialize() })
}

func TestBasicLogging_initialize_badLogger(test *testing.T) {
	messenger, err := messenger.New()
	require.NoError(test, err)
	logger := &BasicLogging{
		Ctx:       context.TODO(),
		messenger: messenger,
	}
	assert.Panics(test, func() { _ = logger.initialize() })
}

func TestBasicLogging_initialize_badMessenger(test *testing.T) {
	logger := &BasicLogging{
		Ctx: context.TODO(),
	}
	assert.Panics(test, func() { _ = logger.initialize() })
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getOptionCallerSkip() *OptionCallerSkip {
	return &OptionCallerSkip{
		Value: 0,
	}
}

func getOptionIDMessages() *OptionIDMessages {
	return &OptionIDMessages{
		Value: idMessagesTest,
	}
}

func getOptionIDStatuses() *OptionIDStatuses {
	return &OptionIDStatuses{
		Value: idStatusesTest,
	}
}

func getOptionIDTemplate() *OptionMessageIDTemplate {
	return &OptionMessageIDTemplate{
		Value: messageIDTemplate,
	}
}

func getOptionLogLevel(logLevelName string) *OptionLogLevel {
	return &OptionLogLevel{
		Value: logLevelName,
	}
}

func getMessageFields() *OptionMessageFields {
	return &OptionMessageFields{
		Value: AllMessageFields,
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
	var err error
	return err
}

func teardown() error {
	var err error
	return err
}
