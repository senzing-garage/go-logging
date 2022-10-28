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

// The MessageLogLevelDefault type returns the logger.Level based on a any logger.Level in details parameter.
type MessageLogLevelDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLogLevel method returns a logger.level based on one or more logger.level types in the details parameter.
func (messageLogLevel *MessageLogLevelDefault) MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	result := logger.LevelInfo

	for _, value := range details {
		switch typedValue := value.(type) {
		case logger.Level:
			result = typedValue
		}
	}

	return result, err
}
