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
type MessageLogLevelSenzingApi struct {
	IdRanges   map[int]string
	IdStatuses map[int]string
}

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
		IdRanges:   messageLogLevel.IdRanges,
		IdStatuses: messageLogLevel.IdStatuses,
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
