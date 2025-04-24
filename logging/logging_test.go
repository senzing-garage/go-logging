package logging_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/senzing-garage/go-logging/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	badComponentIdentifier               = 10000
	badLogLevelName                      = "BadLogLevelName"
	componentID                          = 9997
	messageIDTemplate                    = "test-%04d"
	testDuration           time.Duration = 10000
)

var (
	badIDMessages map[int]string //nolint
	badIDStatuses map[int]string //nolint
)

var idMessagesTest = map[int]string{ //nolint
	0001: "TRACE: %s works with %s",
	1001: "DEBUG: %s works with %s",
	2001: "INFO: %s works with %s",
	3001: "WARN: %s works with %s",
	4001: "ERROR: %s works with %s",
	5001: "FATAL: %s works with %s",
	6001: "PANIC: %s works with %s",
}

var idStatusesTest = map[int]string{ //nolint
	2000: "SUCCESS",
	4000: "FAILURE",
	6000: "DISASTER",
}

var testCasesForMessage = []struct { //nolint
	name                     string
	messageNumber            int
	options                  []interface{}
	details                  []interface{}
	expectedNew              string
	expectedNewSenzingLogger string
}{
	{
		name:          "logging-0001",
		messageNumber: 1,
		options: []interface{}{
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:          "logging-1001",
		messageNumber: 1001,
		options: []interface{}{
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:          "logging-2001",
		messageNumber: 2001,
		options: []interface{}{
			getMessageFields(),
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"INFO","text":"INFO: Bob works with Jane","id":"2001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"INFO","text":"INFO: Bob works with Jane","id":"SZTL99972001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:          "logging-2002",
		messageNumber: 2002,
		options: []interface{}{
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
			getOptionLogLevel("WARN"),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:          "logging-3001",
		messageNumber: 3001,
		options: []interface{}{
			getMessageFields(),
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"WARN","text":"WARN: Bob works with Jane","id":"3001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","text":"WARN: Bob works with Jane","id":"SZTL99973001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:          "logging-3002",
		messageNumber: 3002,
		options: []interface{}{
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
			getOptionLogLevel("ERROR"),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              ``,
		expectedNewSenzingLogger: ``,
	},
	{
		name:                     "logging-3003",
		messageNumber:            3003,
		options:                  []interface{}{getOptionIDStatuses(), getOptionTimeHidden()},
		details:                  []interface{}{},
		expectedNew:              `{"level":"WARN","id":"3003"}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","id":"SZTL99973003"}` + "\n",
	},
	{
		name:                     "logging-3004",
		messageNumber:            3004,
		options:                  []interface{}{getOptionIDTemplate(), getOptionTimeHidden()},
		details:                  []interface{}{},
		expectedNew:              `{"level":"WARN","id":"test-3004"}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","id":"test-3004"}` + "\n",
	},
	{
		name:          "logging-3005",
		messageNumber: 3005,
		options: []interface{}{
			getMessageFields(),
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details: []interface{}{
			"Bob",
			"Jane",
			logging.MessageCode{Value: "Test Code"},
			logging.MessageDuration{Value: 1234567},
			logging.MessageID{Value: "3005-Test"},
			logging.MessageLevel{Value: "ERROR"},
			logging.MessageLocation{Value: "Test Location"},
			logging.MessageReason{Value: "Test reason ..."},
			logging.MessageStatus{Value: "Test-Status"},
			logging.MessageText{Value: "Test text"},
			logging.MessageTime{Value: time.Now()},
			logging.OptionCallerSkip{Value: 3},
			testDuration,
		},
		expectedNew:              `{"level":"WARN","text":"Test text","id":"ERROR","code":"Test Code","reason":"Test reason ...","status":"Test-Status","duration":10000,"location":"In func1() at logging_test.go:473","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","text":"Test text","id":"ERROR","code":"Test Code","reason":"Test reason ...","status":"Test-Status","duration":10000,"location":"In func1() at logging_test.go:539","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:          "logging-3006",
		messageNumber: 3006,
		options: []interface{}{
			logging.OptionMessageField{Value: "id"},
			logging.OptionCallerSkip{Value: 3},
			getOptionTimeHidden(),
		},
		details: []interface{}{
			"Bob",
			"Jane",
			logging.MessageCode{Value: "Test Code"},
			logging.MessageDuration{Value: 1234567},
			logging.MessageID{Value: "3005-Test"},
			logging.MessageLevel{Value: "ERROR"},
			logging.MessageLocation{Value: "Test Location"},
			logging.MessageReason{Value: "Test reason ..."},
			logging.MessageStatus{Value: "Test-Status"},
			logging.MessageText{Value: "Test text"},
			logging.MessageTime{Value: time.Now()},
			logging.OptionCallerSkip{Value: 3},
			testDuration,
		},
		expectedNew:              `{"level":"WARN","text":"Test text","id":"ERROR"}` + "\n",
		expectedNewSenzingLogger: `{"level":"WARN","id":"ERROR","reason":"Test reason ..."}` + "\n",
	},
	{
		name:          "logging-4001",
		messageNumber: 4001,
		options: []interface{}{
			getMessageFields(),
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"4001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"SZTL99974001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:          "logging-5001",
		messageNumber: 5001,
		options: []interface{}{
			getMessageFields(),
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"5001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"SZTL99975001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
	{
		name:          "logging-6001",
		messageNumber: 6001,
		options: []interface{}{
			getMessageFields(),
			getOptionIDMessages(),
			getOptionCallerSkip(),
			getOptionTimeHidden(),
		},
		details:                  []interface{}{"Bob", "Jane"},
		expectedNew:              `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"6001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
		expectedNewSenzingLogger: `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"SZTL99976001","details":[{"position":1,"type":"string","value":"Bob"},{"position":2,"type":"string","value":"Jane"}]}` + "\n",
	},
}

var testCasesForIsValidLogLevelName = []struct { //nolint
	name         string
	logLevelName string
	expected     bool
}{
	{
		name:         "valid-" + logging.LevelTraceName,
		logLevelName: logging.LevelTraceName,
		expected:     true,
	},
	{
		name:         "valid-" + logging.LevelDebugName,
		logLevelName: logging.LevelDebugName,
		expected:     true,
	},
	{
		name:         "valid-" + logging.LevelInfoName,
		logLevelName: logging.LevelInfoName,
		expected:     true,
	},
	{
		name:         "valid-" + logging.LevelWarnName,
		logLevelName: logging.LevelWarnName,
		expected:     true,
	},
	{
		name:         "valid-" + logging.LevelErrorName,
		logLevelName: logging.LevelErrorName,
		expected:     true,
	},
	{
		name:         "valid-" + logging.LevelFatalName,
		logLevelName: logging.LevelFatalName,
		expected:     true,
	},
	{
		name:         "valid-" + logging.LevelPanicName,
		logLevelName: logging.LevelPanicName,
		expected:     true,
	},
	{
		name:         "valid-bad-XYZZY",
		logLevelName: "XYZZY",
		expected:     false,
	},
}

var testCasesForIsXxxx = []struct { //nolint
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
		name:          "is-" + logging.LevelTraceName,
		logLevelName:  logging.LevelTraceName,
		expectedTrace: true,
		expectedDebug: true,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + logging.LevelDebugName,
		logLevelName:  logging.LevelDebugName,
		expectedTrace: false,
		expectedDebug: true,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + logging.LevelInfoName,
		logLevelName:  logging.LevelInfoName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + logging.LevelWarnName,
		logLevelName:  logging.LevelWarnName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + logging.LevelErrorName,
		logLevelName:  logging.LevelErrorName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + logging.LevelFatalName,
		logLevelName:  logging.LevelFatalName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: false,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "is-" + logging.LevelPanicName,
		logLevelName:  logging.LevelPanicName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: false,
		expectedFatal: false,
		expectedPanic: true,
	},
}

// ----------------------------------------------------------------------------
// Test interface methods
// ----------------------------------------------------------------------------

func TestBasicLogging_GetLogLevel(test *testing.T) {
	test.Parallel()

	loggerOptions := []interface{}{}
	logger, err := logging.NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)

	actual := logger.GetLogLevel()
	assert.Equal(test, "INFO", actual)
}

func TestBasicLogging_IsXxxx(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForIsXxxx {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			testObject, _ := logging.New()
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
	test.Parallel()

	loggerOptions := []interface{}{}
	logger, err := logging.NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)

	actual := logger.JSON(9999, "detail")
	assert.NotEmpty(test, actual)
}

func TestBasicLogging_NewError(test *testing.T) {
	test.Parallel()

	loggerOptions := []interface{}{
		getOptionIDMessages(),
		getOptionCallerSkip(),
		getOptionTimeHidden(),
	}
	logger, err := logging.NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)
	err = logger.NewError(4000, "A bad thing")
	require.Error(test, err)
}

func TestBasicLogging_SetLogLevel_badLogLevelName(test *testing.T) {
	test.Parallel()

	loggerOptions := []interface{}{}
	logger, err := logging.NewSenzingLogger(componentID, idMessagesTest, loggerOptions...)
	require.NoError(test, err)
	err = logger.SetLogLevel(badLogLevelName)
	require.Error(test, err)
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestLogging_IsValidLogLevelName(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForIsValidLogLevelName {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual := logging.IsValidLogLevelName(testCase.logLevelName)
			assert.Equal(test, testCase.expected, actual, testCase.name)
		})
	}
}

func TestLogging_New(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-New", func(test *testing.T) {
			test.Parallel()

			outputString := new(bytes.Buffer)
			options := testCase.options
			options = append(options, optionOutput(outputString))
			testObject, err := logging.New(options...)
			require.NoError(test, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNew, outputString.String(), testCase.name)
		})
	}
}

func TestLogging_New_badComponentIdentifier(test *testing.T) {
	test.Parallel()

	options := []interface{}{
		logging.OptionComponentID{
			Value: badComponentIdentifier,
		},
	}
	_, err := logging.New(options...)
	require.Error(test, err)
}

func TestLogging_New_badIDMessages(test *testing.T) {
	test.Parallel()

	options := []interface{}{
		logging.OptionIDMessages{
			Value: badIDMessages,
		},
	}
	_, err := logging.New(options...)
	require.Error(test, err)
}

func TestLogging_New_badIDStatuses(test *testing.T) {
	test.Parallel()

	options := []interface{}{
		logging.OptionIDStatuses{
			Value: badIDStatuses,
		},
	}
	_, err := logging.New(options...)
	require.Error(test, err)
}

func TestLogging_New_badLogLevelName(test *testing.T) {
	test.Parallel()

	options := []interface{}{
		logging.OptionLogLevel{
			Value: badLogLevelName,
		},
	}
	_, err := logging.New(options...)
	require.Error(test, err)
}

func TestLogging_NewSenzingLogger(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-NewSenzingLogger", func(test *testing.T) {
			test.Parallel()

			outputString := new(bytes.Buffer)
			options := testCase.options
			options = append(options, optionOutput(outputString))
			testObject, err := logging.NewSenzingLogger(componentID, idMessagesTest, options...)
			require.NoError(test, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNewSenzingLogger, outputString.String(), testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test private method functions
// ----------------------------------------------------------------------------

// func TestBasicLogging_initialize_badLeveler(test *testing.T) {
// 	messenger, err := messenger.New()
// 	require.NoError(test, err)
// 	var output io.Writer = os.Stderr
// 	var slogLeveler = new(slog.LevelVar)
// 	slogLeveler.Set(slog.LevelInfo)
// 	sLogger := slog.New(slog.NewJSONHandler(output, logging.SlogHandlerOptions(slogLeveler)))
// 	logger := &logging.BasicLogging{
// 		Ctx:       context.TODO(),
// 		logger:    sLogger,
// 		messenger: messenger,
// 	}
// 	assert.Panics(test, func() { _ = logger.initialize() })
// }

// func TestBasicLogging_initialize_badLogger(test *testing.T) {
// 	messenger, err := messenger.New()
// 	require.NoError(test, err)
// 	logger := &logging.BasicLogging{
// 		Ctx:       context.TODO(),
// 		messenger: messenger,
// 	}
// 	assert.Panics(test, func() { _ = logger.initialize() })
// }

// func TestBasicLogging_initialize_badMessenger(test *testing.T) {
// 	logger := &logging.BasicLogging{
// 		Ctx: context.TODO(),
// 	}
// 	assert.Panics(test, func() { _ = logger.initialize() })
// }

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getOptionCallerSkip() logging.OptionCallerSkip {
	return logging.OptionCallerSkip{
		Value: 0,
	}
}

func getOptionIDMessages() logging.OptionIDMessages {
	return logging.OptionIDMessages{
		Value: idMessagesTest,
	}
}

func getOptionIDStatuses() logging.OptionIDStatuses {
	return logging.OptionIDStatuses{
		Value: idStatusesTest,
	}
}

func getOptionIDTemplate() logging.OptionMessageIDTemplate {
	return logging.OptionMessageIDTemplate{
		Value: messageIDTemplate,
	}
}

func getOptionLogLevel(logLevelName string) logging.OptionLogLevel {
	return logging.OptionLogLevel{
		Value: logLevelName,
	}
}

func getMessageFields() logging.OptionMessageFields {
	return logging.OptionMessageFields{
		Value: logging.AllMessageFields,
	}
}

func optionOutput(outputString *bytes.Buffer) logging.OptionOutput {
	return logging.OptionOutput{
		Value: outputString,
	}
}

func getOptionTimeHidden() logging.OptionTimeHidden {
	return logging.OptionTimeHidden{
		Value: true,
	}
}
