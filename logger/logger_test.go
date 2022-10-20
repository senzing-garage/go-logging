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

// -- SetLevel ----------------------------------------------------------------

func TestSetLevel(test *testing.T) {
	SetLevel(LevelPanic)
	SetLevel(LevelFatal)
	SetLevel(LevelError)
	SetLevel(LevelWarn)
	SetLevel(LevelInfo)
	SetLevel(LevelDebug)
	SetLevel(LevelTrace)
}

// -- SetLevelFromString ------------------------------------------------------

func TestSetLevelFromString(test *testing.T) {

	var levelString string

	levelString = "PANIC"
	SetLevelFromString(levelString)
	assert.True(test, LevelPanic == GetLevel(), "Panic")

	levelString = "fatal"
	SetLevelFromString(levelString)
	assert.True(test, LevelFatal == GetLevel(), "Fatal")

	levelString = "ErRoR"
	SetLevelFromString(levelString)
	assert.True(test, LevelError == GetLevel(), "Error")

	levelString = "waRN"
	SetLevelFromString(levelString)
	assert.True(test, LevelWarn == GetLevel(), "Warn")

	levelString = "INFO"
	SetLevelFromString(levelString)
	assert.True(test, LevelInfo == GetLevel(), "Info")

	levelString = "DEBUG"
	SetLevelFromString(levelString)
	assert.True(test, LevelDebug == GetLevel(), "Debug")

	levelString = "TRACE"
	SetLevelFromString(levelString)
	assert.True(test, LevelTrace == GetLevel(), "Trace")

	levelString = "Bad-Level-String"
	SetLevelFromString(levelString)
	assert.True(test, LevelPanic == GetLevel(), "Unknown string returns Panic")

}

// -- GetLevel ----------------------------------------------------------------

func TestGetLevel(test *testing.T) {

	var level Level

	level = LevelPanic
	SetLevel(level)
	assert.True(test, level == GetLevel(), "Panic")

	level = LevelFatal
	SetLevel(level)
	assert.True(test, level == GetLevel(), "Fatal")

	level = LevelError
	SetLevel(level)
	assert.True(test, level == GetLevel(), "Error")

	level = LevelWarn
	SetLevel(level)
	assert.True(test, level == GetLevel(), "Warn")

	level = LevelInfo
	SetLevel(level)
	assert.True(test, level == GetLevel(), "Info")

	level = LevelDebug
	SetLevel(level)
	assert.True(test, level == GetLevel(), "Debug")

	level = LevelTrace
	SetLevel(level)
	assert.True(test, level == GetLevel(), "Trace")

}

// -- GetLevelAsString --------------------------------------------------------

func TestGetLevelAsString(test *testing.T) {

	var level Level

	level = LevelPanic
	SetLevel(level)
	assert.True(test, GetLevelAsString() == "PANIC", "Panic")

	level = LevelFatal
	SetLevel(level)
	assert.True(test, GetLevelAsString() == "FATAL", "Fatal")

	level = LevelError
	SetLevel(level)
	assert.True(test, GetLevelAsString() == "ERROR", "Error")

	level = LevelWarn
	SetLevel(level)
	assert.True(test, GetLevelAsString() == "WARN", "Warn")

	level = LevelInfo
	SetLevel(level)
	assert.True(test, GetLevelAsString() == "INFO", "Info")

	level = LevelDebug
	SetLevel(level)
	assert.True(test, GetLevelAsString() == "DEBUG", "Debug")

	level = LevelTrace
	SetLevel(level)
	assert.True(test, GetLevelAsString() == "TRACE", "Trace")

}

// -- IsPanic -----------------------------------------------------------------

func TestIsPanic(test *testing.T) {
	SetLevel(LevelPanic)
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
	SetLevel(LevelFatal)
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
	SetLevel(LevelError)
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
	SetLevel(LevelWarn)
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
	SetLevel(LevelInfo)
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
	SetLevel(LevelDebug)
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
	SetLevel(LevelTrace)
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
	SetLevel(LevelError)
	assert.NotZero(test, Error("test"), "string")
	assert.NotZero(test, Errorf("test %s", "something"), "format")
}

// -- Warn --------------------------------------------------------------------

func TestWarn(test *testing.T) {
	SetLevel(LevelWarn)
	assert.NotZero(test, Warn("test"), "string")
	assert.NotZero(test, Warnf("test %s", "something"), "format")
}

// -- Info --------------------------------------------------------------------

func TestInfo(test *testing.T) {
	SetLevel(LevelInfo)
	assert.NotZero(test, Info("test"), "string")
	assert.NotZero(test, Infof("test %s", "something"), "format")
}

// -- Debug -------------------------------------------------------------------

func TestDebug(test *testing.T) {
	SetLevel(LevelDebug)
	assert.NotZero(test, Debug("test"), "string")
	assert.NotZero(test, Debugf("test %s", "something"), "format")
}

// -- Trace -------------------------------------------------------------------

func TestTrace(test *testing.T) {
	SetLevel(LevelTrace)
	assert.NotZero(test, Trace("test"), "string")
	assert.NotZero(test, Tracef("test %s", "something"), "format")
}

// -- Miscellaneous -----------------------------------------------------------

func TestFluentInterface(test *testing.T) {
	SetLevel(LevelTrace)
	Trace("trace").Debug("debug").Info("info").Warn("warn").Error("error")
}

func TestVaradic(test *testing.T) {
	SetLevel(LevelDebug)
	_, err := time.LoadLocation("bob")
	Info("Should be error: ", err)
}
