// The logger package is a set of method to help with common tasks.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messageloglevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageLogLevelInterface interface {
	MessageLogLevel(errorNumber int, details ...interface{}) (logger.Level, error)
}
