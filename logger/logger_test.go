package logger_test

import (
	"testing"
	"time"

	"github.com/senzing-garage/go-logging/logger"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCases = []struct { //nolint
	name          string
	logLevel      logger.Level
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
		name:          "Test case: #1 - Trace",
		logLevel:      logger.LevelTrace,
		logLevelName:  logger.LevelTraceName,
		expectedTrace: true,
		expectedDebug: true,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "Test case: #2 - Debug",
		logLevel:      logger.LevelDebug,
		logLevelName:  logger.LevelDebugName,
		expectedTrace: false,
		expectedDebug: true,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "Test case: #3 - Info",
		logLevel:      logger.LevelInfo,
		logLevelName:  logger.LevelInfoName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  true,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "Test case: #4 - Warn",
		logLevel:      logger.LevelWarn,
		logLevelName:  logger.LevelWarnName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  true,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "Test case: #5 - Error",
		logLevel:      logger.LevelError,
		logLevelName:  logger.LevelErrorName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: true,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "Test case: #6 - Fatal",
		logLevel:      logger.LevelFatal,
		logLevelName:  logger.LevelFatalName,
		expectedTrace: false,
		expectedDebug: false,
		expectedInfo:  false,
		expectedWarn:  false,
		expectedError: false,
		expectedFatal: true,
		expectedPanic: true,
	},
	{
		name:          "Test case: #7 - Panic",
		logLevel:      logger.LevelPanic,
		logLevelName:  logger.LevelPanicName,
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
// Test interface functions
// ----------------------------------------------------------------------------

func TestLevels(test *testing.T) {
	test.Parallel()
	assert.Less(test, logger.LevelTrace, logger.LevelDebug, "Trace")
	assert.Less(test, logger.LevelDebug, logger.LevelInfo, "Debug")
	assert.Less(test, logger.LevelInfo, logger.LevelWarn, "Info")
	assert.Less(test, logger.LevelWarn, logger.LevelError, "Warn")
	assert.Less(test, logger.LevelError, logger.LevelFatal, "Error")
	assert.Less(test, logger.LevelFatal, logger.LevelPanic, "Fatal")
}

// -- SetLogLevel -------------------------------------------------------------

func TestSetLevel(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()
			logger.SetLogLevel(testCase.logLevel)
		})
	}
}

// -- SetLogLevelFromString ---------------------------------------------------

func TestSetLogLevelFromString(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()
			logger.SetLogLevelFromString(testCase.logLevelName)
			assert.Equal(test, testCase.logLevel, logger.GetLogLevel(), testCase.name)
		})
	}
}

func TestSetLogLevelFromStringBadString(test *testing.T) {
	test.Parallel()

	levelString := "Bad-Level-String"
	logger.SetLogLevelFromString(levelString)
	assert.Equal(test, logger.LevelPanic, logger.GetLogLevel(), "Unknown string returns Panic")
}

// -- GetLogLevel -------------------------------------------------------------

func TestGetLogLevel(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()
			logger.SetLogLevel(testCase.logLevel)
			assert.Equal(test, testCase.logLevel, logger.GetLogLevel(), testCase.name)
		})
	}
}

// -- GetLogLevelAsString -----------------------------------------------------

func TestGetLogLevelAsString(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()
			logger.SetLogLevel(testCase.logLevel)
			assert.Equal(test, testCase.logLevelName, logger.GetLogLevelAsString(), testCase.name)
		})
	}
}

// -- IsXxxxx -----------------------------------------------------------------

func TestIsXxxx(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			logger, err := logging.New()
			require.NoError(test, err)
			err = logger.SetLogLevel(testCase.logLevelName)
			assert.Equal(test, testCase.logLevelName, logger.GetLogLevel(), testCase.name)
			assert.Equal(test, testCase.expectedTrace, logger.IsTrace(), "Trace")
			assert.Equal(test, testCase.expectedDebug, logger.IsDebug(), "Debug")
			assert.Equal(test, testCase.expectedInfo, logger.IsInfo(), "Info")
			assert.Equal(test, testCase.expectedWarn, logger.IsWarn(), "Warn")
			assert.Equal(test, testCase.expectedError, logger.IsError(), "Error")
			assert.Equal(test, testCase.expectedFatal, logger.IsFatal(), "Fatal")
			assert.Equal(test, testCase.expectedPanic, logger.IsPanic(), "Panic")
			require.NoError(test, err)
		})
	}
}

// -- Trace -------------------------------------------------------------------

func TestTrace(test *testing.T) {
	test.Parallel()
	testObject := logger.New()
	testObject.SetLogLevel(logger.LevelTrace)
	assert.NotZero(test, testObject.Trace("test"), "string")
	assert.NotZero(test, testObject.Tracef("test %s", "something"), "format")
	require.True(test, testObject.IsTrace())
}

func TestTrace_Global(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelTrace)
	require.True(test, logger.IsTrace())
}

// -- Debug -------------------------------------------------------------------

func TestDebug(test *testing.T) {
	test.Parallel()
	testObject := logger.New()
	testObject.SetLogLevel(logger.LevelDebug)
	assert.NotZero(test, testObject.Debug("test"), "string")
	assert.NotZero(test, testObject.Debugf("test %s", "something"), "format")
	require.True(test, testObject.IsDebug())
}

func TestDebug_Global(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelDebug)
	require.True(test, logger.IsDebug())
	assert.NotZero(test, logger.Debug("test"), "string")
	assert.NotZero(test, logger.Debugf("test %s", "something"), "format")
}

// -- Info --------------------------------------------------------------------

func TestInfo(test *testing.T) {
	test.Parallel()
	testObject := logger.New()
	testObject.SetLogLevel(logger.LevelInfo)
	assert.NotZero(test, testObject.Info("test"), "string")
	assert.NotZero(test, testObject.Infof("test %s", "something"), "format")
	require.True(test, testObject.IsInfo())
}

func TestInfo_Global(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelInfo)
	require.True(test, logger.IsInfo())
	assert.NotZero(test, logger.Info("test"), "string")
	assert.NotZero(test, logger.Infof("test %s", "something"), "format")
}

// -- Warn --------------------------------------------------------------------

func TestWarn(test *testing.T) {
	test.Parallel()
	testObject := logger.New()
	testObject.SetLogLevel(logger.LevelWarn)
	assert.NotZero(test, testObject.Warn("test"), "string")
	assert.NotZero(test, testObject.Warnf("test %s", "something"), "format")
	require.True(test, testObject.IsWarn())
}

func TestWarn_Global(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelWarn)
	require.True(test, logger.IsWarn())
	assert.NotZero(test, logger.Warn("test"), "string")
	assert.NotZero(test, logger.Warnf("test %s", "something"), "format")
}

// -- Error -------------------------------------------------------------------

func TestError(test *testing.T) {
	test.Parallel()
	testObject := logger.New()
	testObject.SetLogLevel(logger.LevelError)
	assert.NotZero(test, testObject.Error("test"), "string")
	assert.NotZero(test, testObject.Errorf("test %s", "something"), "format")
}

func TestError_Global(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelError)
	require.True(test, logger.IsError())
	assert.NotZero(test, logger.Error("test"), "string")
	assert.NotZero(test, logger.Errorf("test %s", "something"), "format")
}

// -- Fatal -------------------------------------------------------------------

func TestFatal(test *testing.T) {
	// IMPROVE: Figure out how to test Fatal and Fatalf
	test.Parallel()
	testObject := logger.New()
	testObject.SetLogLevel(logger.LevelFatal)
	require.True(test, testObject.IsFatal())
}

func TestFatal_Global(test *testing.T) {
	// IMPROVE: Figure out how to test Fatal and Fatalf
	test.Parallel()
	logger.SetLogLevel(logger.LevelFatal)
	require.True(test, logger.IsFatal())
}

// -- Panic -------------------------------------------------------------------

func TestPanic(test *testing.T) {
	test.Parallel()
	testObject := logger.New()
	testObject.SetLogLevel(logger.LevelPanic)
	assert.Panics(test, func() { testObject.Panic("test") })
	assert.Panics(test, func() { testObject.Panicf("test %s", "something") })
	require.True(test, testObject.IsPanic())
}

func TestPanic_Global(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelPanic)
	require.True(test, logger.IsPanic())
	assert.Panics(test, func() { logger.Panic("test") })
	assert.Panics(test, func() { logger.Panicf("test %s", "something") })
}

// -- Miscellaneous -----------------------------------------------------------

func TestFluentInterface(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelTrace)
	logger.Trace("trace").Debug("debug").Info("info").Warn("warn").Error("error")
}

func TestVaradic(test *testing.T) {
	test.Parallel()
	logger.SetLogLevel(logger.LevelDebug)

	_, err := time.LoadLocation("bob")
	logger.Info("Should be error: ", err)
}
