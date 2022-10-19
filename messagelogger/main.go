// The messagelogger package is a set of methods logging messages.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagelogger

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Level int

type MessageLogger struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Must match what's in logger/main.go

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messagelogger *MessageLogger
