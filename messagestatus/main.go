/*
The messagestatus package produces a value for the "status" field.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagestatus/messagestatus_test.go
*/
package messagestatus

import "github.com/senzing/go-logging/logger"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageStatusInterface type defines methods for determining status.
type MessageStatusInterface interface {
	MessageStatus(messageNumber int, details ...interface{}) (string, error) // Get the "status" value from the messageNumber and details.
}

// The Status type is used to identify strings as being status strings in details parameter.
type Status string

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

/*
Types of Senzing errors.
These are the strings that may be returned from MessageStatus()
*/
const (
	Debug              = logger.LevelDebugName
	Error              = logger.LevelErrorName
	ErrorBadUserInput  = logger.LevelErrorName + "_bad_user_input"
	ErrorRetryable     = logger.LevelErrorName + "_retryable"
	ErrorUnrecoverable = logger.LevelErrorName + "_unrecoverable"
	Fatal              = logger.LevelFatalName
	Info               = logger.LevelInfoName
	Panic              = logger.LevelPanicName
	Trace              = logger.LevelTraceName
	Warn               = logger.LevelWarnName
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A map of Senzing errors to the corresponding error level.
var SenzingApiErrorsMap = map[string]string{
	"0002E":  Info,
	"0019E":  ErrorUnrecoverable,
	"0037E":  ErrorBadUserInput,  // Unknown resolved entity value
	"0052E":  ErrorBadUserInput,  // Unknown relationship ID value
	"0063E":  ErrorUnrecoverable, // G2ConfigMgr is not initialized
	"7221E":  ErrorUnrecoverable, // No engine configuration registered
	"9990E":  Trace,              // Mock error
	"9991E":  Debug,              // Mock error
	"9992E":  Info,               // Mock error
	"9993E":  Warn,               // Mock error
	"9994E":  Error,              // Mock error
	"9995E":  ErrorRetryable,     // Mock error
	"9996E":  ErrorBadUserInput,  // Mock error
	"9997E":  ErrorUnrecoverable, // Mock error
	"9998E":  Fatal,              // Mock error
	"9999E":  Panic,              // Mock error
	"30121E": ErrorBadUserInput,  // JSON parsing Failure
}

// The order of severity/verbosity from most severe to most verbose.
var MessagePrecedence = []string{
	Panic,
	Fatal,
	ErrorUnrecoverable,
	ErrorBadUserInput,
	ErrorRetryable,
	Error,
	Warn,
	Info,
	Debug,
	Trace,
}

var IdLevelRangesAsString = map[int]string{
	0000: logger.LevelTraceName,
	1000: logger.LevelDebugName,
	2000: logger.LevelInfoName,
	3000: logger.LevelWarnName,
	4000: logger.LevelErrorName,
	5000: logger.LevelFatalName,
	6000: logger.LevelPanicName,
}
