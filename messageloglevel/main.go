/*
Package logger provides...
*/
package messageloglevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageLogLevelInterface interface {
	MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error)
}
