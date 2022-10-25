/*
Package messageloglevel returns the "final" log level.
*/
package messageloglevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLogLevelInterface interface {
	MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error)
}
