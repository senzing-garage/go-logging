/*
Package messageloglevel produces a log level.
*/
package messageloglevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLogLevelInterface interface {
	MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error)
}
