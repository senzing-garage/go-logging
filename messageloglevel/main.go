/*
The messageloglevel package produces a log level.
*/
package messageloglevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLogLevelInterface type defines methods for determining log level.
type MessageLogLevelInterface interface {
	MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error)
}
