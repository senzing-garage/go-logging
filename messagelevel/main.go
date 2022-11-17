/*
The messagelevel package produces a value for the "level" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagelevel/messagelevel_test.go
*/
package messagelevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLevelInterface type defines methods for determining log level.
type MessageLevelInterface interface {
	MessageLevel(messageNumber int, details ...interface{}) (logger.Level, error) // Get the "level" value from the messageNumber and details.
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// IdLevelRanges is a map from message IDs to log level "lower-bound" used by Senzing applications.
var IdLevelRanges = map[int]logger.Level{
	0000: logger.LevelTrace,
	1000: logger.LevelDebug,
	2000: logger.LevelInfo,
	3000: logger.LevelWarn,
	4000: logger.LevelError,
	5000: logger.LevelFatal,
	6000: logger.LevelPanic,
}

// IdLevelRangesAsString is a map from message IDs to log level (as string) "lower-bound" used by Senzing applications.
var IdLevelRangesAsString = map[int]string{
	0000: logger.LevelTraceName,
	1000: logger.LevelDebugName,
	2000: logger.LevelInfoName,
	3000: logger.LevelWarnName,
	4000: logger.LevelErrorName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}
