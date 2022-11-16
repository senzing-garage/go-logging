/*
The messagelevel package produces a log level.
*/
package messagelevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelInterface type defines methods for determining log level.
type MessageLevelInterface interface {
	MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error)
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdLevelRanges = map[int]logger.Level{
	0000: logger.LevelTrace,
	1000: logger.LevelDebug,
	2000: logger.LevelInfo,
	3000: logger.LevelWarn,
	4000: logger.LevelError,
	5000: logger.LevelFatal,
	6000: logger.LevelPanic,
}

var IdLevelRangesAsString = map[int]string{
	0000: logger.LevelTraceName,
	1000: logger.LevelDebugName,
	2000: logger.LevelInfoName,
	3000: logger.LevelWarnName,
	4000: logger.LevelErrorName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}
