/*
The MessageLogLevelSenzingApi implementation returns the logger.Level based on the "status" value.
*/
package messageloglevel

import (
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagestatus"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLogLevelSenzingApi type is for calculating the log level based on the status value.
type MessageLogLevelSenzingApi struct{}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Pseudo-constant.
var nameToLevelMap = map[string]logger.Level{
	logger.LevelInfoName:  logger.LevelInfo,
	logger.LevelWarnName:  logger.LevelWarn,
	logger.LevelErrorName: logger.LevelError,
	logger.LevelDebugName: logger.LevelDebug,
	logger.LevelTraceName: logger.LevelTrace,
	logger.LevelFatalName: logger.LevelFatal,
	logger.LevelPanicName: logger.LevelPanic,
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLogLevel method returns a log level based on the status value.
func (messageLogLevel *MessageLogLevelSenzingApi) MessageLogLevel(messageNumber int, details ...interface{}) (logger.Level, error) {
	var result = logger.LevelError

	// Get Status string.

	messageStatus := &messagestatus.MessageStatusSenzingApi{
		IdRanges: map[int]string{
			0000: logger.LevelInfoName,
			1000: logger.LevelWarnName,
			2000: logger.LevelErrorName,
			3000: logger.LevelDebugName,
			4000: logger.LevelTraceName,
			5000: logger.LevelFatalName,
			6000: logger.LevelPanicName,
		},
	}
	status, err := messageStatus.MessageStatus(messageNumber, details...)
	if err != nil {
		return result, err
	}

	// If status is a known logging level, return it.

	result, ok := nameToLevelMap[status]
	if ok {
		return result, err
	}

	// Anything else is an "ERROR"

	return logger.LevelError, err
}
