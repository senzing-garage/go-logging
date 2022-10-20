// The logger package is a set of method to help with common tasks.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagelevel

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageLevelInterface interface {
	CalculateMessageLevel(errorNumber int, message string) (logger.Level, error)
}
