/*
The MessageLogLevelDefault implementation returns the logger.Level based on a any logger.Level in details parameter.
*/
package messageloglevel

import (
	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLogLevelDefault struct{}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var logLevelPrecedence = []logger.Level{
	logger.LevelPanic,
	logger.LevelFatal,
	logger.LevelError,
	logger.LevelWarn,
	logger.LevelInfo,
	logger.LevelDebug,
	logger.LevelTrace,
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Return a logger.level based on one or more logger.level types in the details parameter.
func (messagelevel *MessageLogLevelDefault) MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	result := logger.LevelInfo

	var senzingErrors []logger.Level
	if len(details) > 0 {
		for index := len(details) - 1; index >= 0; index-- {
			detail := details[index]
			switch typedDetail := detail.(type) {
			case logger.Level:
				senzingErrors = append(senzingErrors, typedDetail)
			}
		}
	}

	if len(senzingErrors) > 0 {
		for _, logLevelPrecedenceLevel := range logLevelPrecedence {
			for _, senzingError := range senzingErrors {
				if senzingError == logLevelPrecedenceLevel {
					return senzingError, err
				}
			}
		}
	}

	return result, err

}
