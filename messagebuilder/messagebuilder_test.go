package messagebuilder

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestBuildError(test *testing.T) {
	err := BuildError("unique-id-%04d", 5, "Error message")
	test.Log("Actual:", err)
}

func TestLogMessage(test *testing.T) {
	LogMessage(MessageIdFormat, 2000, "Test message", "Variable1", "Variable2")
}

func TestLogMessageFromError(test *testing.T) {
	anError := errors.New("This is a new error")
	LogMessageFromError(MessageIdFormat, 2002, "Test message", anError, "Variable1", "Variable2")
}

func TestLevels(test *testing.T) {
	assert.True(test, LevelTrace < LevelDebug, "Trace")
	assert.True(test, LevelDebug < LevelInfo, "Debug")
	assert.True(test, LevelInfo < LevelWarn, "Info")
	assert.True(test, LevelWarn < LevelError, "Warn")
	assert.True(test, LevelError < LevelFatal, "Error")
	assert.True(test, LevelFatal < LevelPanic, "Fatal")
}

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

func TestTrace(test *testing.T) {
	SetLevel(LevelTrace)
	assert.NotZero(test, Trace("test"), "string")
	assert.NotZero(test, Tracef("test %s", "something"), "format")
}

func TestDebug(test *testing.T) {
	SetLevel(LevelDebug)
	assert.NotZero(test, Debug("test"), "string")
	assert.NotZero(test, Debugf("test %s", "something"), "format")
}

func TestInfo(test *testing.T) {
	SetLevel(LevelInfo)
	assert.NotZero(test, Info("test"), "string")
	assert.NotZero(test, Infof("test %s", "something"), "format")
}

func TestWarn(test *testing.T) {
	SetLevel(LevelWarn)
	assert.NotZero(test, Warn("test"), "string")
	assert.NotZero(test, Warnf("test %s", "something"), "format")
}

func TestError(test *testing.T) {
	SetLevel(LevelError)
	assert.NotZero(test, Error("test"), "string")
	assert.NotZero(test, Errorf("test %s", "something"), "format")
}

func TestFluentInterface(test *testing.T) {
	SetLevel(LevelTrace)
	Trace("trace").Debug("debug").Info("info").Warn("warn").Error("error")
}

func TestVaradic(test *testing.T) {
	SetLevel(LevelDebug)
	_, err := time.LoadLocation("bob")
	Info("Should be error: ", err)
}
