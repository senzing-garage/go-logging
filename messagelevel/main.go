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
