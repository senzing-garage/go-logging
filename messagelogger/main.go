/*
The messagelogger package generates messages, logs messages, or creates errors from messages.
*/
package messagelogger

import (
	"errors"
	"fmt"
	"sync"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagedate"
	"github.com/senzing/go-logging/messagedetails"
	"github.com/senzing/go-logging/messageduration"
	"github.com/senzing/go-logging/messageerrors"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelevel"
	"github.com/senzing/go-logging/messagelocation"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
	"github.com/senzing/go-logging/messagetime"
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
// Variables
// ----------------------------------------------------------------------------

var (
	lock                   = &sync.Mutex{}
	isSystemLogLevelSet    = false
	systemLogLevel         = LevelInfo
	messageLoggerObservers = []MessageLoggerInterface{}
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
  - messagelevel.MessageLevelInterface
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
		MessageDate:     &messagedate.MessageDateNull{},
		MessageDetails:  &messagedetails.MessageDetailsNull{},
		MessageDuration: &messageduration.MessageDurationNull{},
		MessageErrors:   &messageerrors.MessageErrorsNull{},
		MessageFormat:   &messageformat.MessageFormatDefault{},
		MessageId:       &messageid.MessageIdDefault{},
		MessageLocation: &messagelocation.MessageLocationNull{},
		MessageLevel: &messagelevel.MessageLevelDefault{
			DefaultLogLevel: logger.LevelInfo,
		},
		MessageStatus: &messagestatus.MessageStatusNull{},
		MessageText:   &messagetext.MessageTextNull{},
		MessageTime:   &messagetime.MessageTimeNull{},
	}

	// Incorporate parameters.

	var errorsList []interface{}
	if len(interfaces) > 0 {
		for _, value := range interfaces {
			switch typedValue := value.(type) {
			case logger.LoggerInterface:
				result.Logger = typedValue
			case messagedate.MessageDateInterface:
				result.MessageDate = typedValue
			case messagedetails.MessageDetailsInterface:
				result.MessageDetails = typedValue
			case messageduration.MessageDurationInterface:
				result.MessageDuration = typedValue
			case messageerrors.MessageErrorsInterface:
				result.MessageErrors = typedValue
			case messageformat.MessageFormatInterface:
				result.MessageFormat = typedValue
			case messageid.MessageIdInterface:
				result.MessageId = typedValue
			case messagelevel.MessageLevelInterface:
				result.MessageLevel = typedValue
			case messagelocation.MessageLocationInterface:
				result.MessageLocation = typedValue
			case messagestatus.MessageStatusInterface:
				result.MessageStatus = typedValue
			case messagetext.MessageTextInterface:
				result.MessageText = typedValue
			case messagetime.MessageTimeInterface:
				result.MessageTime = typedValue
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

	// If system logging level set, set this logger to that level and
	// add this messageLogger to the Observers list.
	// Do this in a thread-safe way.

	lock.Lock()
	defer lock.Unlock()

	if isSystemLogLevelSet {
		result.SetLogLevel(systemLogLevel)
	}

	if messageLoggerObservers == nil {
		messageLoggerObservers = make([]MessageLoggerInterface, 0)
	}

	messageLoggerObservers = append(messageLoggerObservers, result)

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

	// Defaults

	callerSkip := 1

	// Look for specific flags in details.

	if len(interfaces) > 0 {
		for _, value := range interfaces {
			switch typedValue := value.(type) {
			case messagelocation.CallerSkip:
				callerSkip = int(typedValue)
			}
		}
	}

	// Construct the components of the messagelogger.

	messageDate := &messagedate.MessageDateSenzing{}
	messageTime := &messagetime.MessageTimeSenzing{}
	messageDuration := &messageduration.MessageDurationSenzing{}
	messageFormat := &messageformat.MessageFormatSenzing{}
	messageErrors := &messageerrors.MessageErrorsSenzing{}
	messageDetails := &messagedetails.MessageDetailsSenzing{}

	messageLocation := &messagelocation.MessageLocationSenzing{
		CallerSkip: callerSkip,
	}

	messageId := &messageid.MessageIdSenzing{
		MessageIdTemplate: fmt.Sprintf("senzing-%04d", productIdentifier) + "%04d",
	}

	messageLogLevel := &messagelevel.MessageLevelSenzing{
		DefaultLogLevel: logger.LevelInfo,
		IdRanges: map[int]logger.Level{
			0000: logger.LevelTrace,
			1000: logger.LevelDebug,
			2000: logger.LevelInfo,
			3000: logger.LevelWarn,
			4000: logger.LevelError,
			5000: logger.LevelFatal,
			6000: logger.LevelPanic,
		},
	}

	messageStatus := &messagestatus.MessageStatusSenzing{
		IdRanges: map[int]string{
			0000: logger.LevelTraceName,
			1000: logger.LevelDebugName,
			2000: logger.LevelInfoName,
			3000: logger.LevelWarnName,
			4000: logger.LevelErrorName,
			5000: logger.LevelFatalName,
			6000: logger.LevelPanicName,
		},
	}

	messageText := &messagetext.MessageTextSenzing{
		IdMessages: idMessages,
	}

	var newInterfaces = []interface{}{
		messageDate,
		messageTime,
		messageLocation,
		messageId,
		messageFormat,
		messageLogLevel,
		messageStatus,
		messageText,
		messageDuration,
		messageErrors,
		messageDetails,
	}

	// Add other user-supplied interfaces to newInterfaces.

	newInterfaces = append(newInterfaces, interfaces...)

	// Using a Factory Pattern, build the messagelogger.

	return New(newInterfaces...)
}

// ----------------------------------------------------------------------------
// Functions
// ----------------------------------------------------------------------------

func GetLogLevel() (Level, error) {
	var err error = nil
	if !isSystemLogLevelSet {
		err = fmt.Errorf("system log level not set")
	}
	return systemLogLevel, err
}

func GetLogLevelAsString() (string, error) {
	var err error = nil

	logLevel, err := GetLogLevel()
	if err != nil {
		return "", err
	}

	levelName, ok := logger.LevelToTextMap[logger.Level(logLevel)]
	if !ok {
		err = fmt.Errorf("logLevel not found")
	}

	return levelName, err
}

func SetLogLevel(level Level) error {
	var err error = nil

	// Make this thread-safe.

	lock.Lock()
	defer lock.Unlock()

	// Propagate level across all messageLoggers.

	isSystemLogLevelSet = true
	systemLogLevel = level
	for _, messageLogger := range messageLoggerObservers {
		messageLogger.SetLogLevel(systemLogLevel)
	}
	return err
}

func SetLogLevelFromString(levelName string) error {

	level, ok := logger.TextToLevelMap[levelName]
	if ok {
		return SetLogLevel(Level(level))
	}

	return fmt.Errorf("unknown log level name: %s", levelName)
}
