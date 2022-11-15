package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name          string
	logLevel      Level
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
		logLevel:      LevelTrace,
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
		name:          "Test case: #2 - Debug",
		logLevel:      LevelDebug,
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
		name:          "Test case: #3 - Info",
		logLevel:      LevelInfo,
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
		name:          "Test case: #4 - Warn",
		logLevel:      LevelWarn,
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
		name:          "Test case: #5 - Error",
		logLevel:      LevelError,
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
		name:          "Test case: #6 - Fatal",
		logLevel:      LevelFatal,
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
		name:          "Test case: #7 - Panic",
		logLevel:      LevelPanic,
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

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestLevels(test *testing.T) {
	assert.True(test, LevelTrace < LevelDebug, "Trace")
	assert.True(test, LevelDebug < LevelInfo, "Debug")
	assert.True(test, LevelInfo < LevelWarn, "Info")
	assert.True(test, LevelWarn < LevelError, "Warn")
	assert.True(test, LevelError < LevelFatal, "Error")
	assert.True(test, LevelFatal < LevelPanic, "Fatal")
}

// -- SetLogLevel -------------------------------------------------------------

func TestSetLevel(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			SetLogLevel(testCase.logLevel)
		})
	}
}

// -- SetLogLevelFromString ---------------------------------------------------

func TestSetLogLevelFromString(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			SetLogLevelFromString(testCase.logLevelName)
			assert.Equal(test, testCase.logLevel, GetLogLevel(), testCase.name)
		})
	}
}

func TestSetLogLevelFromStringBadString(test *testing.T) {
	levelString := "Bad-Level-String"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelPanic == GetLogLevel(), "Unknown string returns Panic")
}

// -- GetLogLevel -------------------------------------------------------------

func TestGetLogLevel(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			SetLogLevel(testCase.logLevel)
			assert.Equal(test, testCase.logLevel, GetLogLevel(), testCase.name)
		})
	}
}

// -- GetLogLevelAsString -----------------------------------------------------

func TestGetLogLevelAsString(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			SetLogLevel(testCase.logLevel)
			assert.Equal(test, testCase.logLevelName, GetLogLevelAsString(), testCase.name)
		})
	}
}

// -- IsXxxxx -----------------------------------------------------------------

func TestIsXxxx(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			SetLogLevel(testCase.logLevel)
			assert.Equal(test, testCase.logLevelName, GetLogLevelAsString(), testCase.name)
			assert.Equal(test, testCase.expectedTrace, IsTrace(), "Trace")
			assert.Equal(test, testCase.expectedDebug, IsDebug(), "Debug")
			assert.Equal(test, testCase.expectedInfo, IsInfo(), "Info")
			assert.Equal(test, testCase.expectedWarn, IsWarn(), "Warn")
			assert.Equal(test, testCase.expectedError, IsError(), "Error")
			assert.Equal(test, testCase.expectedFatal, IsFatal(), "Fatal")
			assert.Equal(test, testCase.expectedPanic, IsPanic(), "Panic")
		})
	}
}

// -- Error -------------------------------------------------------------------

func TestError(test *testing.T) {
	SetLogLevel(LevelError)
	assert.NotZero(test, Error("test"), "string")
	assert.NotZero(test, Errorf("test %s", "something"), "format")
}

// -- Warn --------------------------------------------------------------------

func TestWarn(test *testing.T) {
	SetLogLevel(LevelWarn)
	assert.NotZero(test, Warn("test"), "string")
	assert.NotZero(test, Warnf("test %s", "something"), "format")
}

// -- Info --------------------------------------------------------------------

func TestInfo(test *testing.T) {
	SetLogLevel(LevelInfo)
	assert.NotZero(test, Info("test"), "string")
	assert.NotZero(test, Infof("test %s", "something"), "format")
}

// -- Debug -------------------------------------------------------------------

func TestDebug(test *testing.T) {
	SetLogLevel(LevelDebug)
	assert.NotZero(test, Debug("test"), "string")
	assert.NotZero(test, Debugf("test %s", "something"), "format")
}

// -- Trace -------------------------------------------------------------------

func TestTrace(test *testing.T) {
	SetLogLevel(LevelTrace)
	assert.NotZero(test, Trace("test"), "string")
	assert.NotZero(test, Tracef("test %s", "something"), "format")
}

// -- Miscellaneous -----------------------------------------------------------

func TestFluentInterface(test *testing.T) {
	SetLogLevel(LevelTrace)
	Trace("trace").Debug("debug").Info("info").Warn("warn").Error("error")
}

func TestVaradic(test *testing.T) {
	SetLogLevel(LevelDebug)
	_, err := time.LoadLocation("bob")
	Info("Should be error: ", err)
}
