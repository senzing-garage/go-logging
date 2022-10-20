package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
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
	SetLogLevel(LevelPanic)
	SetLogLevel(LevelFatal)
	SetLogLevel(LevelError)
	SetLogLevel(LevelWarn)
	SetLogLevel(LevelInfo)
	SetLogLevel(LevelDebug)
	SetLogLevel(LevelTrace)
}

// -- SetLogLevelFromString ---------------------------------------------------

func TestSetLogLevelFromString(test *testing.T) {

	var levelString string

	levelString = "PANIC"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelPanic == GetLogLevel(), "Panic")

	levelString = "fatal"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelFatal == GetLogLevel(), "Fatal")

	levelString = "ErRoR"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelError == GetLogLevel(), "Error")

	levelString = "waRN"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelWarn == GetLogLevel(), "Warn")

	levelString = "INFO"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelInfo == GetLogLevel(), "Info")

	levelString = "DEBUG"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelDebug == GetLogLevel(), "Debug")

	levelString = "TRACE"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelTrace == GetLogLevel(), "Trace")

	levelString = "Bad-Level-String"
	SetLogLevelFromString(levelString)
	assert.True(test, LevelPanic == GetLogLevel(), "Unknown string returns Panic")

}

// -- GetLogLevel -------------------------------------------------------------

func TestGetLogLevel(test *testing.T) {

	var level Level

	level = LevelPanic
	SetLogLevel(level)
	assert.True(test, level == GetLogLevel(), "Panic")

	level = LevelFatal
	SetLogLevel(level)
	assert.True(test, level == GetLogLevel(), "Fatal")

	level = LevelError
	SetLogLevel(level)
	assert.True(test, level == GetLogLevel(), "Error")

	level = LevelWarn
	SetLogLevel(level)
	assert.True(test, level == GetLogLevel(), "Warn")

	level = LevelInfo
	SetLogLevel(level)
	assert.True(test, level == GetLogLevel(), "Info")

	level = LevelDebug
	SetLogLevel(level)
	assert.True(test, level == GetLogLevel(), "Debug")

	level = LevelTrace
	SetLogLevel(level)
	assert.True(test, level == GetLogLevel(), "Trace")

}

// -- GetLogLevelAsString -----------------------------------------------------

func TestGetLogLevelAsString(test *testing.T) {

	var level Level

	level = LevelPanic
	SetLogLevel(level)
	assert.True(test, GetLogLevelAsString() == "PANIC", "Panic")

	level = LevelFatal
	SetLogLevel(level)
	assert.True(test, GetLogLevelAsString() == "FATAL", "Fatal")

	level = LevelError
	SetLogLevel(level)
	assert.True(test, GetLogLevelAsString() == "ERROR", "Error")

	level = LevelWarn
	SetLogLevel(level)
	assert.True(test, GetLogLevelAsString() == "WARN", "Warn")

	level = LevelInfo
	SetLogLevel(level)
	assert.True(test, GetLogLevelAsString() == "INFO", "Info")

	level = LevelDebug
	SetLogLevel(level)
	assert.True(test, GetLogLevelAsString() == "DEBUG", "Debug")

	level = LevelTrace
	SetLogLevel(level)
	assert.True(test, GetLogLevelAsString() == "TRACE", "Trace")

}

// -- IsPanic -----------------------------------------------------------------

func TestIsPanic(test *testing.T) {
	SetLogLevel(LevelPanic)
	assert.False(test, IsTrace(), "Trace")
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.False(test, IsWarn(), "Warn")
	assert.False(test, IsError(), "Error")
	assert.False(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

// -- IsFatal -----------------------------------------------------------------

func TestIsFatal(test *testing.T) {
	SetLogLevel(LevelFatal)
	assert.False(test, IsTrace(), "Trace")
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.False(test, IsWarn(), "Warn")
	assert.False(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

// -- IsError -----------------------------------------------------------------

func TestIsError(test *testing.T) {
	SetLogLevel(LevelError)
	assert.False(test, IsTrace(), "Trace")
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.False(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

// -- IsWarn ------------------------------------------------------------------

func TestIsWarn(test *testing.T) {
	SetLogLevel(LevelWarn)
	assert.False(test, IsTrace(), "Trace")
	assert.False(test, IsDebug(), "Debug")
	assert.False(test, IsInfo(), "Info")
	assert.True(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

// -- IsInfo ------------------------------------------------------------------

func TestIsInfo(test *testing.T) {
	SetLogLevel(LevelInfo)
	assert.False(test, IsTrace(), "Trace")
	assert.False(test, IsDebug(), "Debug")
	assert.True(test, IsInfo(), "Info")
	assert.True(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

// -- IsDebug -----------------------------------------------------------------

func TestIsDebug(test *testing.T) {
	SetLogLevel(LevelDebug)
	assert.False(test, IsTrace(), "Trace")
	assert.True(test, IsDebug(), "Debug")
	assert.True(test, IsInfo(), "Info")
	assert.True(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

// -- IsTrace -----------------------------------------------------------------

func TestIsTrace(test *testing.T) {
	SetLogLevel(LevelTrace)
	assert.True(test, IsTrace(), "Trace")
	assert.True(test, IsDebug(), "Debug")
	assert.True(test, IsInfo(), "Info")
	assert.True(test, IsWarn(), "Warn")
	assert.True(test, IsError(), "Error")
	assert.True(test, IsFatal(), "Fatal")
	assert.True(test, IsPanic(), "Panic")
}

// -- Panic -------------------------------------------------------------------

// No tests.

// -- Fatal -------------------------------------------------------------------

// No tests.

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
