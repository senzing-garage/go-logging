/*
Package helper ...
*/
package messageloglevel

import (
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagestatus"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageLogLevelSenzingApi struct{}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var NameToLevelMap = map[string]logger.Level{
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

// TODO:
func (messagelevel *MessageLogLevelSenzingApi) MessageLogLevel(errorNumber int, details ...interface{}) (logger.Level, error) {
	var err error = nil
	var result = logger.LevelError

	messageStatus := &messagestatus.MessageStatusSenzingApi{}
	status, err := messageStatus.MessageStatus(errorNumber, details...)

	if err != nil {
		return result, err
	}

	result, ok := NameToLevelMap[status]
	if ok {
		return result, err
	}

	return logger.LevelError, err
}
