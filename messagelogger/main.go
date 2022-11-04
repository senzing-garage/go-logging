/*
The messagelogger package generates messages, logs messages, or creates errors from messages.
*/
package messagelogger

import (
	"errors"
	"fmt"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messageloglevel"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The Level type is used for logging levels. (e.g. LevelInfo, LevelWarn, etc.)
type Level int

/*
The MessageLoggerInterface type defines methods for creating messages, logging messages, and creating errors from messages.
It also has convenience methods for setting and getting the current log level.
*/
type MessageLoggerInterface interface {
	Error(messageNumber int, details ...interface{}) error
	GetLogLevel() Level
	GetLogLevelAsString() string
	IsDebug() bool
	IsError() bool
	IsFatal() bool
	IsInfo() bool
	IsPanic() bool
	IsTrace() bool
	IsWarn() bool
	Log(messageNumber int, details ...interface{}) error
	Message(messageNumber int, details ...interface{}) (string, error)
	SetLogLevel(level Level) MessageLoggerInterface
	SetLogLevelFromString(levelString string) MessageLoggerInterface
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// LevelXxxx values are an enumeration of typed integers representing logging levels.
// Must match what's in logger/main.go
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

/*
The New function creates a new instance of MessageLoggerDefault.
The default message logger uses default and null subcomponents.
To use non-default subcomponents,
adding parameters to New() can specify the subcomponent desired.
The parameters can be of the following type and in any order:

  - logger.LoggerInterface
  - messageformat.MessageFormatInterface
  - messageid.MessageIdInterface
  - messageloglevel.MessageLogLevelInterface
  - messagestatus.MessageStatusInterface
  - messagetext.MessageTextInterface
  - logger.Level

If a type is specified multiple times,
the last instance instance of the type specified wins.
*/
func New(interfaces ...interface{}) (MessageLoggerInterface, error) {
	var err error = nil

	// Start with default values.

	logLevel := LevelInfo
	result := &MessageLoggerDefault{
		Logger:          &logger.LoggerDefault{},
		MessageFormat:   &messageformat.MessageFormatDefault{},
		MessageId:       &messageid.MessageIdDefault{},
		MessageLogLevel: &messageloglevel.MessageLogLevelDefault{},
		MessageStatus:   &messagestatus.MessageStatusNull{},
		MessageText:     &messagetext.MessageTextNull{},
	}

	// Incorporate parameters.

	var errorsList []interface{}
	if len(interfaces) > 0 {
		for _, value := range interfaces {
			switch typedValue := value.(type) {
			case logger.LoggerInterface:
				result.Logger = typedValue
			case messageformat.MessageFormatInterface:
				result.MessageFormat = typedValue
			case messageid.MessageIdInterface:
				result.MessageId = typedValue
			case messageloglevel.MessageLogLevelInterface:
				result.MessageLogLevel = typedValue
			case messagestatus.MessageStatusInterface:
				result.MessageStatus = typedValue
			case messagetext.MessageTextInterface:
				result.MessageText = typedValue
			case logger.Level:
				logLevelCandidate, ok := value.(logger.Level)
				if ok {
					logLevel = Level(logLevelCandidate)
				}
			default:
				errorsList = append(errorsList, typedValue)
			}
		}
	}
	result.SetLogLevel(logLevel)

	// Report any unknown parameters.

	if len(errorsList) > 0 {
		err = fmt.Errorf("unsupported interfaces: %#v", errorsList)
	}

	return result, err
}

func NewSenzingLogger(productIdentifier int, idMessages map[int]string, interfaces ...interface{}) (MessageLoggerInterface, error) {
	var err error = nil

	// Detect incorrect parameter values.

	if productIdentifier <= 0 || productIdentifier >= 10000 {
		err = errors.New("productIdentifier must be in range 1..9999. See https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-product-ids.md")
		return nil, err
	}

	if idMessages == nil {
		err = errors.New("messages must be a map[int]string")
		return nil, err
	}

	// Construct the components of the messagelogger.

	messageFormat := &messageformat.MessageFormatSenzing{}

	messageId := &messageid.MessageIdSenzing{
		MessageIdTemplate: fmt.Sprintf("senzing-%04d", productIdentifier) + "%04d",
	}

	messageLogLevel := &messageloglevel.MessageLogLevelSenzing{
		DefaultLogLevel: logger.LevelInfo,
		IdRanges: map[int]logger.Level{
			0000: logger.LevelInfo,
			1000: logger.LevelWarn,
			2000: logger.LevelError,
			3000: logger.LevelDebug,
			4000: logger.LevelTrace,
			5000: logger.LevelFatal,
			6000: logger.LevelPanic,
		},
	}

	messageStatus := &messagestatus.MessageStatusSenzing{
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

	messageText := &messagetext.MessageTextSenzing{
		IdMessages: idMessages,
	}

	var newInterfaces = []interface{}{
		messageId,
		messageFormat,
		messageLogLevel,
		messageStatus,
		messageText,
	}

	// Add other user-supplied interfaces to newInterfaces.

	newInterfaces = append(newInterfaces, interfaces...)

	// Using a Factory Pattern, build the messagelogger.

	return New(newInterfaces...)
}
