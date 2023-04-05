package slogconfig

import "golang.org/x/exp/slog"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

/*
LevelXxxx values are an enumeration of typed integers representing logging levels.
Order is important for the LevelXxxx variables.
*/
const (
	LevelDebugSlog = slog.LevelDebug
	LevelErrorSlog = slog.LevelError
	LevelFatalSlog = slog.Level(LevelFatalInt)
	LevelInfoSlog  = slog.LevelInfo
	LevelPanicSlog = slog.Level(LevelPanicInt)
	LevelTraceSlog = slog.Level(LevelTraceInt)
	LevelWarnSlog  = slog.LevelWarn
)

// Strings representing the supported logging levels.
const (
	LevelDebugName = "DEBUG"
	LevelErrorName = "ERROR"
	LevelFatalName = "FATAL"
	LevelInfoName  = "INFO"
	LevelPanicName = "PANIC"
	LevelTraceName = "TRACE"
	LevelWarnName  = "WARN"
)

const (
	LevelTraceInt int = -8
	LevelDebugInt int = -4
	LevelInfoInt  int = 0
	LevelWarnInt  int = 4
	LevelErrorInt int = 8
	LevelFatalInt int = 12
	LevelPanicInt int = 16
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Map from Log level as typed integer to string representation.
var LevelToTextMap = map[slog.Level]string{
	LevelDebugSlog: LevelDebugName,
	LevelErrorSlog: LevelErrorName,
	LevelFatalSlog: LevelFatalName,
	LevelInfoSlog:  LevelInfoName,
	LevelPanicSlog: LevelPanicName,
	LevelTraceSlog: LevelTraceName,
	LevelWarnSlog:  LevelWarnName,
}

// Map from string representation to Log level as typed integer.
var TextToLevelMap = map[string]slog.Level{
	LevelDebugName: LevelDebugSlog,
	LevelErrorName: LevelErrorSlog,
	LevelFatalName: LevelFatalSlog,
	LevelInfoName:  LevelInfoSlog,
	LevelPanicName: LevelPanicSlog,
	LevelTraceName: LevelTraceSlog,
	LevelWarnName:  LevelWarnSlog,
}

// Map from string representation to integer.
var TextToIntMap = map[string]int{
	LevelDebugName: LevelDebugInt,
	LevelErrorName: LevelErrorInt,
	LevelFatalName: LevelFatalInt,
	LevelInfoName:  LevelInfoInt,
	LevelPanicName: LevelPanicInt,
	LevelTraceName: LevelTraceInt,
	LevelWarnName:  LevelWarnInt,
}

// Map from integer to string representation.
var InToTextMap = map[int]string{
	LevelDebugInt: LevelDebugName,
	LevelErrorInt: LevelErrorName,
	LevelFatalInt: LevelFatalName,
	LevelInfoInt:  LevelInfoName,
	LevelPanicInt: LevelPanicName,
	LevelTraceInt: LevelTraceName,
	LevelWarnInt:  LevelWarnName,
}
